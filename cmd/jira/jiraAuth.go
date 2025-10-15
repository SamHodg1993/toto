package jira

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/samhodg1993/toto/internal/utilities"
	"github.com/zalando/go-keyring"

	"github.com/spf13/cobra"
)

var cloudId string = ""

var JiraAuth = &cobra.Command{
	Use:   "jira-auth",
	Short: "Authenticate with Jira",
	Long:  "Authenticate with Jira using API tokens and store credentials in operating system keyring",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Check keyring for existing credentials
		jiraEmail, _ := keyring.Get("toto-cli", "jira-email")
		jiraApiToken, _ := keyring.Get("toto-cli", "jira-api-token")

		if jiraEmail == "" || jiraApiToken == "" {
			fmt.Println("\nJira API Token Setup Required")
			fmt.Println("==============================\n")
			fmt.Println("To use Jira integration, you need to create an API token:\n")
			fmt.Println("1. Go to: https://id.atlassian.com/manage-profile/security/api-tokens")
			fmt.Println("2. Click 'Create API token'")
			fmt.Println("3. Give it a label (e.g., 'Toto CLI')")
			fmt.Println("4. Copy the generated token\n")

			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Enter your Jira account email: ")
			jiraEmail, _ = reader.ReadString('\n')
			jiraEmail = strings.TrimSpace(jiraEmail)

			fmt.Print("Enter your Jira API token: ")
			jiraApiToken, _ = reader.ReadString('\n')
			jiraApiToken = strings.TrimSpace(jiraApiToken)

			// Store in keyring for future use
			err := keyring.Set("toto-cli", "jira-email", jiraEmail)
			if err != nil {
				fmt.Printf("Failed to store email in keyring: %v\n", err)
				return
			}

			err = keyring.Set("toto-cli", "jira-api-token", jiraApiToken)
			if err != nil {
				fmt.Printf("Failed to store API token in keyring: %v\n", err)
				return
			}

			fmt.Println("\nCredentials saved to keyring!")
		} else {
			fmt.Println("\nJira credentials already configured!")
			fmt.Printf("Email: %s\n", jiraEmail)
			fmt.Println("\nTo reconfigure, delete the credentials from your keyring first.")
			return
		}

		// Attempt to fetch and store cloud ID
		fmt.Println("\nFetching Jira Cloud ID...")
		err := utilities.GetUsersJiraCloudId()
		if err != nil {
			fmt.Printf("Warning: Failed to automatically fetch cloud ID: %v\n", err)
			fmt.Println("You can manually set it later using: toto jira-set-cloud-id -i <cloud-id>")
		} else {
			fmt.Println("Cloud ID stored successfully!")
		}

		fmt.Println("\nJira authentication complete! You can now use jira-pull commands.")
	},
}

var JiraSetCloudId = &cobra.Command{
	Use:   "jira-set-cloud-id",
	Short: "Manually set cloud id",
	Long:  "To be used in the event a user needs to manually set their jira cloud id",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if cloudId == "" {
			fmt.Println("Cloud id not presented. Please use -i <cloud id here>")
			return
		}

		if err := keyring.Set("toto-cli", "jira-cloud-id", cloudId); err != nil {
			return
		}
		fmt.Println("Cloud ID stored successfully")
	},
}

func init() {
	JiraSetCloudId.PersistentFlags().StringVarP(&cloudId, "cloudId", "i", "", "Cloud ID to be set")
}
