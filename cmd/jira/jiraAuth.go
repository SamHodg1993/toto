package jira

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/google/uuid"
	"github.com/samhodg1993/toto-todo-cli/internal/service"
	"github.com/samhodg1993/toto-todo-cli/internal/utilities"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var JiraAuth = &cobra.Command{
	Use:   "jira-auth",
	Short: "Authenticate with Jira",
	Long:  "Authenticate with Jira and store user credentials in operating system keyring",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		uuid := uuid.New().String()

		redirectUrl := "http://localhost:8989/callback"

		authUrl := "https://auth.atlassian.com/authorize?audience=api.atlassian.com&client_id=aJjbbYpk9l3k3xzZ2YR7s0giptkT9ppg&scope=read%3Ajira-work%20write%3Ajira-work%20offline_access&redirect_uri=" + url.QueryEscape(redirectUrl) + "&state=" + uuid + "&response_type=code&prompt=consent"

		// Load credentials from env
		clientId := os.Getenv("JIRA_CLIENT_ID")
		clientSecret := os.Getenv("JIRA_CLIENT_SECRET")
		if clientId == "" || clientSecret == "" {
			fmt.Printf("Either the client ID or client secret variables are missing from the env file.\n")
			return
		}

		utilities.OpenBrowser(authUrl)

		code, err := service.StartCallbackServer(uuid)
		if err != nil {
			fmt.Printf("Authentication failed: %v\n", err)
			return
		}

		exchangeUrl := "https://auth.atlassian.com/oauth/token"
		data := url.Values{
			"grant_type":    {"authorization_code"},
			"client_id":     {clientId},
			"client_secret": {clientSecret},
			"code":          {code},
			"redirect_uri":  {redirectUrl},
		}

		resp, err := http.PostForm(exchangeUrl, data)
		if err != nil {
			fmt.Printf("Failed to exchange code for token %v\n", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			fmt.Printf("Token exchange failed with status code %d", resp.StatusCode)
			return
		}

		utilities.StoreJiraCredentialsInKeyring(resp)
	},
}

var JiraTest = &cobra.Command{
	Use:   "jira-test-print",
	Short: "testing real quick",
	Long:  "",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		accessToken, err := keyring.Get("toto-cli", "jira-access-token")
		if err != nil {
			fmt.Printf("There was an error %v\n", err)
		}

		refreshToken, err := keyring.Get("toto-cli", "jira-refresh-token")
		if err != nil {
			fmt.Printf("There was an error %v\n", err)
		}

		fmt.Printf("acc: %s\n\n", accessToken)
		fmt.Printf("ref: %s\n\n", refreshToken)
	},
}
