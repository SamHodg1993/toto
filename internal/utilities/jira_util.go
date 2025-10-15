package utilities

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/samhodg1993/toto/internal/models"

	"github.com/zalando/go-keyring"
)

func HandleJiraSessionBeforeCall() (string, error) {
	accessTokenExpiry, err := keyring.Get("toto-cli", "access-token-expiry")
	if err != nil {
		_, err := keyring.Get("toto-cli", "jira-access-token")
		if err != nil {
			return "", fmt.Errorf("No valid access token found. Please run 'toto jira-auth' to authenticate")
		}

		refreshToken, err := keyring.Get("toto-cli", "jira-refresh-token")
		if err != nil {
			return "", fmt.Errorf("No refresh token found. Please run 'toto jira-auth' to re-authenticate")
		}

		return refreshTokens(refreshToken)
	}

	unixTime, err := strconv.ParseInt(accessTokenExpiry, 10, 64)
	if err != nil {
		return "", err
	}

	parsedTime := time.Unix(unixTime, 0)
	if time.Now().Before(parsedTime) {
		accessToken, err := keyring.Get("toto-cli", "jira-access-token")
		if err != nil {
			return "", fmt.Errorf("There was an error getting the access token for jira %v\n", err)
		}
		return accessToken, nil
	}

	refreshToken, err := keyring.Get("toto-cli", "jira-refresh-token")
	if err != nil {
		return "", fmt.Errorf("There was an error getting the refresh token for jira %v\n", err)
	}

	return refreshTokens(refreshToken)
}

func refreshTokens(refreshToken string) (string, error) {
	clientId := os.Getenv("JIRA_CLIENT_ID")
	clientSecret := os.Getenv("JIRA_CLIENT_SECRET")
	if clientId == "" || clientSecret == "" {
		return "", fmt.Errorf("Failed to collect the client secret of client id from environment variables during routine HandleJiraSessionBeforeCall")
	}

	data := url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {clientId},
		"client_secret": {clientSecret},
		"refresh_token": {refreshToken},
	}

	resp, err := http.PostForm("https://auth.atlassian.com/oauth/token",
		data)
	if err != nil {
		return "", fmt.Errorf("Failed to refresh token: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Failed to refresh token, status code: %d received", resp.StatusCode)
	}

	token, err := StoreJiraCredentialsInKeyring(resp)
	if err != nil {
		return "", err
	}

	return token, nil
}

func StoreJiraCredentialsInKeyring(resp *http.Response) (string, error) {
	var tokenResp models.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("Failed to decode response: %v\n", err)
	}

	expiryTime := time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)
	if err := keyring.Set("toto-cli", "jira-access-token", tokenResp.AccessToken); err != nil {
		return "", fmt.Errorf("Failed to store access token: %v\n", err)
	}
	if err := keyring.Set("toto-cli", "jira-refresh-token", tokenResp.RefreshToken); err != nil {
		return "", fmt.Errorf("Failed to store refresh token: %v\n", err)
	}
	if err := keyring.Set("toto-cli", "access-token-expiry", strconv.FormatInt(expiryTime.Unix(), 10)); err != nil {
		return "", fmt.Errorf("Failed to store access token expiry: %v\n", err)
	}
	return tokenResp.AccessToken, nil
}

func GetUsersJiraCloudId(accessToken string) (string, error) {
	fmt.Println("Getting users cloud ID")

	req, err := http.NewRequest("GET", "https://api.atlassian.com/oauth/token/accessible-resources", nil)
	if err != nil {
		return "", fmt.Errorf("Failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Unable to get users cloud id: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Unable to get users cloud id. Request returned status code: %d", resp.StatusCode)
	}

	var resources []struct {
		ID   string `json:"id"`
		URL  string `json:"url"`
		Name string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&resources); err != nil {
		return "", fmt.Errorf("Failed to decode response: %v", err)
	}

	if len(resources) == 0 {
		return "", fmt.Errorf("User has no available cloud id's. User must manually set their own using command jira-set-cloud-id -i <client id string>")
	}

	if err := keyring.Set("toto-cli", "jira-cloud-id", resources[0].ID); err != nil {
		return "", fmt.Errorf("Failed to set jira cloud ID in keyring with error: %v", err)
	}

	fmt.Println("Cloud ID stored successfully")
	return resources[0].ID, nil
}
