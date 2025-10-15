package jira

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/samhodg1993/toto/internal/models"
	"github.com/samhodg1993/toto/internal/utilities"
	"github.com/zalando/go-keyring"
)

// GetSingleJiraTicket fetches a single Jira ticket by issue key from the API
func (s *Service) GetSingleJiraTicket(issueKey string) (*models.JiraBasedTicket, error) {
	// Get API token credentials from keyring
	jiraEmail, err := keyring.Get("toto-cli", "jira-email")
	if err != nil {
		return nil, fmt.Errorf("Jira credentials not found. Please run 'toto jira-auth' first")
	}

	jiraApiToken, err := keyring.Get("toto-cli", "jira-api-token")
	if err != nil {
		return nil, fmt.Errorf("Jira API token not found. Please run 'toto jira-auth' first")
	}

	cloudId, err := keyring.Get("toto-cli", "jira-cloud-id")
	if err != nil {
		err := utilities.GetUsersJiraCloudId()
		if err != nil {
			return nil, fmt.Errorf("Could not retrieve cloud ID automatically: %v. Try running 'toto jira-set-cloud-id -i <your-cloud-id>'", err)
		}

		cloudId, err = keyring.Get("toto-cli", "jira-cloud-id")
		if err != nil || cloudId == "" {
			return nil, fmt.Errorf("There was an error getting the cloud id for jira. An attempt to fix this automatically was attempted but failed")
		}
	}

	url := fmt.Sprintf("https://api.atlassian.com/ex/jira/%s/rest/api/3/issue/%s", cloudId, issueKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request client with error: %v", err)
	}

	// Use Basic Auth with email:apiToken base64 encoded
	auth := jiraEmail + ":" + jiraApiToken
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", "Basic "+encodedAuth)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Request for issue failed with error: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Request for issue failed with status code: %d. This issue likely doesn't exist. Please check the input id.", resp.StatusCode)
	}

	var ticket models.JiraBasedTicket
	if err := json.NewDecoder(resp.Body).Decode(&ticket); err != nil {
		return nil, fmt.Errorf("Failed to decode response with error: %v", err)
	}

	return &ticket, nil
}
