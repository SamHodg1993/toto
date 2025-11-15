package jira

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/odgy8/toto/internal/models"
	"github.com/odgy8/toto/internal/utilities"
)

// GetSingleJiraTicket fetches a single Jira ticket by issue key from the API
func (s *Service) GetSingleJiraTicket(issueKey string) (*models.JiraBasedTicket, error) {
	jiraURL, email, apiKey, err := utilities.HandleJiraSessionBeforeCall(s.projectService)
	if err != nil {
		return nil, fmt.Errorf("%v\n", err)
	}

	url := fmt.Sprintf("%s/rest/api/3/issue/%s", jiraURL,
		issueKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request client with error: %v", err)
	}
	req.SetBasicAuth(email, apiKey)
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
