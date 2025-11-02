package utilities

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/zalando/go-keyring"
)

type ProjectService interface {
	GetProjectJiraURL() (string, error)
}

func StoreJiraCredentialsInKeyring(jiraURL, email, apiKey string) error {
	endpoint := fmt.Sprintf("%s/rest/api/3/myself", jiraURL)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(email, apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// If status code is not 200 then the email + api + jira url combop is incorrect
	if resp.StatusCode != 200 {
		return fmt.Errorf("Status code %d returned. Incorrect URL, email, api key combination.", resp.StatusCode)
	}

	// We only want to store the keyring at this point if there isn't already one stored.
	existingKeyring, _ := keyring.Get("toto-cli", "jiraURL")
	if existingKeyring == "" {
		if err := keyring.Set("toto-cli", "jiraURL", jiraURL); err != nil {
			return fmt.Errorf("Failed to store Jira URL: %w", err)
		}
	}

	if err := keyring.Set("toto-cli", "jiraEmail", email); err != nil {
		return fmt.Errorf("Failed to store email address: %w", err)
	}

	if err := keyring.Set("toto-cli", "jiraApiKey", apiKey); err != nil {
		return fmt.Errorf("Failed to store api key: %w", err)
	}

	fmt.Println("\nCredentials saved to keyring!")
	return nil
}

func EnsureHTTPS(url string) string {
	url = strings.TrimSpace(url)

	if strings.HasPrefix(url, "http://") {
		return "https://" + strings.TrimPrefix(url, "http://")
	}

	if !strings.HasPrefix(url, "https://") {
		return "https://" + url
	}

	return url
}

func HandleJiraSessionBeforeCall(projectService ProjectService) (jiraURL, email, apiKey string, err error) {
	jiraURL, err = projectService.GetProjectJiraURL()
	jiraURL = EnsureHTTPS(jiraURL)

	email, err = keyring.Get("toto-cli", "jiraEmail")
	if err != nil {
		return "", "", "", fmt.Errorf("No Jira email found. Please run 'toto jira-auth' to authenticate")
	}

	apiKey, err = keyring.Get("toto-cli", "jiraApiKey")
	if err != nil {
		return "", "", "", fmt.Errorf("No Jira API key found. Please run 'toto jira-auth' to authenticate")
	}

	return jiraURL, email, apiKey, nil
}
