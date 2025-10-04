package utilities

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/samhodg1993/toto-todo-cli/internal/models"

	"github.com/zalando/go-keyring"
)

func HandleJiraSessionBeforeCall() (string, error) {
	accessTokenExpiry, err := keyring.Get("toto-cli", "access-token-expiry")
	if err != nil {
		return "", fmt.Errorf("There was an error getting the access token expiry %v\n", err)
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

	// Get the refresh token and get a new access token
	refreshToken, err := keyring.Get("toto-cli", "jira-refresh-token")
	if err != nil {
		return "", fmt.Errorf("There was an error getting the refresh token for jira %v\n", err)
	}

	clientId := os.Getenv("JIRA_CLIENT_ID")
	clientSecret := os.Getenv("JIRA_CLIENT_SECRET")
	data := url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {clientId},
		"client_secret": {clientSecret},
		"refresh_token": {refreshToken},
	}

	resp, err := http.PostForm("https://auth.atlassian.com/oauth/token", data)
	if err != nil {
		return "", fmt.Errorf("Failed to refresh token: %v", err)
	}
	defer resp.Body.Close()

	token, err := StoreJiraCredentialsInKeyring(resp)
	if err != nil {
		return "", err
	}

	// return the access token
	return token, nil
}

func StoreJiraCredentialsInKeyring(resp *http.Response) (string, error) {
	var tokenResp models.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("Failed to decode response: %v\n", err)
	}

	expiryTime := time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)
	err := keyring.Set("toto-cli", "jira-access-token", tokenResp.AccessToken)
	if err != nil {
		return "", fmt.Errorf("Failed to store access token: %v\n", err)
	}
	err = keyring.Set("toto-cli", "jira-refresh-token", tokenResp.RefreshToken)
	if err != nil {
		return "", fmt.Errorf("Failed to store refresh token: %v\n", err)
	}
	err = keyring.Set("toto-cli", "access-token-expiry", strconv.FormatInt(expiryTime.Unix(), 10))
	if err != nil {
		return "", fmt.Errorf("Failed to store access token expiry: %v\n", err)
	}
	return tokenResp.AccessToken, nil

}
