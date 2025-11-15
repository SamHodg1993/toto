package jira

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/odgy8/toto/internal/utilities"

	"github.com/spf13/cobra"
)

var cloudId string = ""

var JiraAuth = &cobra.Command{
	Use:   "jira-auth",
	Short: "Authenticate with Jira",
	Long:  "Authenticate with Jira and store user credentials in operating system keyring",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			jiraURL,
			email,
			apiKey string
		)

		fmt.Println("\nJira Authentication needed")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("To use Jira integration, provide your organisations Jira URL")
		fmt.Println("Here is an example Jira URL: https://mycompany.atlassian.net.")
		fmt.Println("You will also need to provide your email address and api key.")
		fmt.Println("1. Go to: https://id.atlassian.com/manage-profile/security/api-tokens")
		fmt.Println("2. Name it (e.g., 'Toto CLI')")
		fmt.Println("3. Click 'Create'")
		fmt.Println("4. Copy the token immediately - it won't be shown again!")
		fmt.Println("5. Store it securely then provide it when prompted.\n")

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter the Jira URL for your organisation: ")
		jiraURL, _ = reader.ReadString('\n')
		jiraURL = strings.TrimSpace(jiraURL)

		fmt.Print("Enter the email address with jira access: ")
		email, _ = reader.ReadString('\n')
		email = strings.TrimSpace(email)

		fmt.Print("Enter your jira API key: ")
		apiKey, _ = reader.ReadString('\n')
		apiKey = strings.TrimSpace(apiKey)

		utilities.StoreJiraCredentialsInKeyring(jiraURL, email, apiKey)
	},
}
