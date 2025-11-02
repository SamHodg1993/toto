package utilityCommands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var (
	jiraDefaultUrlString string = ""
)

var SetDefaultJiraUrl = &cobra.Command{
	Use:   "jira-set-default-url",
	Short: "Updates the stored default jira url.",
	Long:  "Updates the jira url which is used as a default when using jira on a project. Stored securely in the system keyring.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		jiraDefaultUrlString = strings.TrimSpace(jiraDefaultUrlString)
		keyring.Set("toto-cli", "jiraURL", jiraDefaultUrlString)
		fmt.Println("Updated default Jira URL")
	},
}

func init() {
	SetDefaultJiraUrl.Flags().StringVarP(&jiraDefaultUrlString, "jira-url", "u", "", "Jira URL to be stored as the default Jira URL")
}
