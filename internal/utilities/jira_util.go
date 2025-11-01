package utilities

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zalando/go-keyring"
)

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

	// If status code is not 200 then the email + api key combo is incorrect
	if resp.StatusCode != 200 {
		return fmt.Errorf("Status code %d returned. Incorrect URL, email, api key combination.", resp.StatusCode)
	}

	// Store in keyring for future use
	if err := keyring.Set("toto-cli", "jiraURL", jiraURL); err != nil {
		return fmt.Errorf("Failed to store Jira URL: %w", err)
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

func HandleJiraSessionBeforeCall() (jiraURL string, email string, apiKey string, err error) {
	jiraURL, err = keyring.Get("toto-cli", "jiraURL")
	if err != nil {
		return "", "", "", fmt.Errorf("No Jira URL found. Please run 'toto jira-auth' to authenticate")
	}

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
