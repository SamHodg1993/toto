package jira

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/samhodg1993/toto/internal/models"
	"github.com/samhodg1993/toto/internal/utilities"
	"github.com/zalando/go-keyring"
)

// GetSingleJiraTicket fetches a single Jira ticket by issue key from the API
func (s *Service) GetSingleJiraTicket(issueKey string) (*models.JiraBasedTicket, error) {
	accessToken, err := utilities.HandleJiraSessionBeforeCall()
	if err != nil {
		return nil, fmt.Errorf("%v\n", err)
	}

	cloudId, err := keyring.Get("toto-cli", "jira-cloud-id")
	if err != nil {
		newId, err := utilities.GetUsersJiraCloudId(accessToken)
		if err != nil {
			return nil, fmt.Errorf("Could not retrieve cloud ID automatically: %v. Try running 'toto jira-set-cloud-id -i <your-cloud-id>'", err)
		}

		if newId == "" {
			return nil, fmt.Errorf("There was an error getting the cloud id for jira. An attempt to fix this automatically was attempted but retuned an empty id string")
		}
		cloudId = newId
	}

	url := fmt.Sprintf("https://api.atlassian.com/ex/jira/%s/rest/api/3/issue/%s", cloudId, issueKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request client with error: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
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
