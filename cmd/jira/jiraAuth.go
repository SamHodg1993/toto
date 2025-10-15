package jira

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

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

		// Fall back to keyring
		clientId, _ = keyring.Get("toto-cli", "jira-client-id")
		clientSecret, _ = keyring.Get("toto-cli", "jira-client-secret")

		if clientId == "" || clientSecret == "" {
			fmt.Println("\nJira OAuth Setup Required")
			fmt.Println("=========================\n")
			fmt.Println("To use Jira integration, you need to create an OAuth 2.0 app:\n")
			fmt.Println("1. Go to: https://developer.atlassian.com/console/myapps/")
			fmt.Println("2. Click 'Create' → 'OAuth 2.0 integration'")
			fmt.Println("3. Name it (e.g., 'Toto CLI')")
			fmt.Println("4. Add permissions: read:jira-work, write:jira-work, offline_access")
			fmt.Println("5. Set callback URL: http://localhost:8989/callback")
			fmt.Println("6. Save and copy your credentials\n")

			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Enter your Jira OAuth Client ID: ")
			clientId, _ = reader.ReadString('\n')
			clientId = strings.TrimSpace(clientId)

			fmt.Print("Enter your Jira OAuth Client Secret: ")
			clientSecret, _ = reader.ReadString('\n')
			clientSecret = strings.TrimSpace(clientSecret)

			// Store in keyring for future use
			keyring.Set("toto-cli", "jira-client-id", clientId)
			keyring.Set("toto-cli", "jira-client-secret", clientSecret)
			fmt.Println("\nCredentials saved to keyring!")
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
