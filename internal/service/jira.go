package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"net/http"

	"github.com/samhodg1993/toto-todo-cli/internal/models"
	"github.com/samhodg1993/toto-todo-cli/internal/utilities"
	"github.com/zalando/go-keyring"
)

type JiraService struct {
	db *sql.DB
}

func NewJiraService(db *sql.DB) *JiraService {
	return &JiraService{db: db}
}

func (j *JiraService) InsertJiraTicket(ticket *models.JiraTicket) (int64, error) {
	// Check if ticket already exists
	var existingId int64
	err := j.db.QueryRow("SELECT id FROM jira_tickets WHERE jira_key = ?", ticket.JiraKey).Scan(&existingId)

	if err == nil {
		// Ticket exists, update it and return existing ID
		_, updateErr := j.db.Exec(
			`UPDATE jira_tickets
			SET title = ?, status = ?, project_key = ?, issue_type = ?, url = ?, last_synced_at = ?
			WHERE jira_key = ?`,
			ticket.Title,
			ticket.Status,
			ticket.ProjectKey,
			ticket.IssueType,
			ticket.URL,
			time.Now(),
			ticket.JiraKey,
		)
		if updateErr != nil {
			return 0, fmt.Errorf("Failed to update existing jira ticket: %v", updateErr)
		}
		return existingId, nil
	}

	// Ticket doesn't exist, insert new one
	result, err := j.db.Exec(
		`INSERT INTO jira_tickets (
		  jira_key, title, status, project_key, issue_type, url, last_synced_at
		) VALUES (?,?,?,?,?,?,?)`,
		ticket.JiraKey,
		ticket.Title,
		ticket.Status,
		ticket.ProjectKey,
		ticket.IssueType,
		ticket.URL,
		time.Now(),
	)

	if err != nil {
		return 0, fmt.Errorf("Could not insert new jira ticket, err: %v", err)
	}

	returnId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Insert Jira ticket success, but failed to get row ID, err: %v", err)
	}

	return returnId, nil
}

func (j *JiraService) GetSingleJiraTicket(issueKey string) (*models.JiraBasedTicket, error) {
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

func StartCallbackServer(expectedState string) (string, error) {
	codeChan := make(chan string, 1)
	errorChan := make(chan error, 1)

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")

		// Make sure the state is what we expect
		if state != expectedState {
			errorChan <- fmt.Errorf("invalid state parameter")
			return
		}

		fmt.Println("We have the users creds. Store in keyring next!")

		// Success, tell the user to close the browser window
		w.WriteHeader(200)
		w.Write([]byte("Authentication successful! You can close this window."))

		// Send code to channel
		codeChan <- code
	})

	// Start server
	server := &http.Server{Addr: ":8989"} // Port is specific to the app. Configured in atlassian developer console
	go server.ListenAndServe()

	// Wait for callback or timeout
	select {
	case code := <-codeChan:
		server.Shutdown(context.Background())
		return code, nil
	case err := <-errorChan:
		server.Shutdown(context.Background())
		return "", err
	case <-time.After(2 * time.Minute):
		server.Shutdown(context.Background())
		return "", fmt.Errorf("authentication timeout")
	}
}
