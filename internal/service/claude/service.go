package claude

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/zalando/go-keyring"
)

// SubTask represents a task broken down by Claude AI
type SubTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// GetClaudeAPIKey retrieves the Claude API key from environment or keyring
func GetClaudeAPIKey() (string, error) {
	// Try environment variable
	if apiKey := os.Getenv("CLAUDE_API_KEY"); apiKey != "" {
		return apiKey, nil
	}

	// Fall back to keyring
	apiKey, err := keyring.Get("toto-cli", "claude-api-key")
	if err != nil {
		return "", fmt.Errorf("Claude API key not found. Please set ANTHROPIC_API_KEY environment variable or store in keyring")
	}
	return apiKey, nil
}

// BreakdownJiraTicketWithClaude uses Claude AI to break down a Jira ticket into subtasks
func BreakdownJiraTicketWithClaude(ticketKey, title, description string) ([]SubTask, error) {
	apiKey, err := GetClaudeAPIKey()
	if err != nil {
		return nil, err
	}

	client := anthropic.NewClient(option.WithAPIKey(apiKey))

	prompt := fmt.Sprintf(`You are a task breakdown assistant. You will receive a Jira ticket and you must create actionable todos from it.

Jira Ticket: %s
Title: %s
Description:
%s

YOUR TASK:
1. Read the Title and Description carefully
2. If it contains a list of specific tasks (with bullet points, line breaks, "Create todo to...", etc.), extract EACH item EXACTLY as written
3. If it's vague or high-level with no list, break it down into 3-8 actionable subtasks

Examples of what to extract:
- Description: "Create todo to git add .\n\nCreate todo to git commit\n\nCreate todo to git push"
  → Extract: [{"title": "git add .", "description": "Stage all changes"}, {"title": "git commit -am <message>", "description": "Commit staged changes"}, {"title": "git push", "description": "Push to remote"}]

- Description: "Implement authentication system"
  → Break down: [{"title": "Create user model", "description": "..."}, {"title": "Add password hashing", "description": "..."}, ...]

Return ONLY a JSON array with "title" and "description" fields. No other text.`, ticketKey, title, description)

	message, err := client.Messages.New(context.Background(), anthropic.MessageNewParams{
		Model:     "claude-3-5-sonnet-20241022",
		MaxTokens: 2000,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to call Claude API: %w", err)
	}

	// Extract text from response
	var responseText string
	for _, block := range message.Content {
		if textBlock := block.Text; textBlock != "" {
			responseText += textBlock
		}
	}

	// Clean up response
	responseText = strings.TrimSpace(responseText)
	responseText = strings.TrimPrefix(responseText, "```json")
	responseText = strings.TrimPrefix(responseText, "```")
	responseText = strings.TrimSuffix(responseText, "```")
	responseText = strings.TrimSpace(responseText)

	// Parse JSON
	var subtasks []SubTask
	if err := json.Unmarshal([]byte(responseText), &subtasks); err != nil {
		return nil, fmt.Errorf("failed to parse Claude response as JSON: %w\nResponse: %s", err, responseText)
	}

	return subtasks, nil
}
