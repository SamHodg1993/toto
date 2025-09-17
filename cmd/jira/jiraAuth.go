package jira

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/samhodg1993/toto-todo-cli/internal/service"
	"github.com/samhodg1993/toto-todo-cli/internal/utilities"

	"github.com/spf13/cobra"
)

var JiraAuth = &cobra.Command{
	Use:   "jira-auth",
	Short: "Authenticate with Jira",
	Long:  "Authenticate with Jira and store user credentials in operating system keyring",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		uuid := uuid.New().String()

		url := "https://auth.atlassian.com/authorize?audience=api.atlassian.com&client_id=aJjbbYpk9l3k3xzZ2YR7s0giptkT9ppg&scope=read%3Ajira-work%20write%3Ajira-work&redirect_uri=http%3A%2F%2Flocalhost%3A8989%2Fcallback&state=" + uuid + "&response_type=code&prompt=consent"

		utilities.OpenBrowser(url)

		code, err := service.StartCallbackServer(uuid)
		if err != nil {
			fmt.Printf("Authentication failed: %v\n", err)
			return
		}

		fmt.Printf("Got authorization code: %s\n", code)

		// TODO: Swap code for access token
		// https://developer.atlassian.com/cloud/jira/platform/oauth-2-3lo-apps/
		// TODO: Store access token and refresh token in keyring probably

		fmt.Println("Making progress. Need to get the access token and then make the jira functions")
	},
}
