package jira

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/google/uuid"
	"github.com/samhodg1993/toto/internal/service/jira"
	"github.com/samhodg1993/toto/internal/utilities"
	"github.com/zalando/go-keyring"

	"github.com/spf13/cobra"
)

var cloudId string = ""

var JiraAuth = &cobra.Command{
	Use:   "jira-auth",
	Short: "Authenticate with Jira",
	Long:  "Authenticate with Jira and store user credentials in operating system keyring",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		uuid := uuid.New().String()

		redirectUrl := "http://localhost:8989/callback"

		// Load credentials from env
		clientId := os.Getenv("JIRA_CLIENT_ID")
		clientSecret := os.Getenv("JIRA_CLIENT_SECRET")
		if clientId == "" || clientSecret == "" {
			fmt.Printf("Either the client ID or client secret variables are missing from the env file.\n")
			return
		}

		authUrl := "https://auth.atlassian.com/authorize?audience=api.atlassian.com&client_id=" + clientId + "&scope=read%3Ajira-work%20write%3Ajira-work%20offline_access&redirect_uri=" + url.QueryEscape(redirectUrl) + "&state=" + uuid + "&response_type=code&prompt=consent"
		utilities.OpenBrowser(authUrl)

		code, err := jira.StartCallbackServer(uuid)
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
