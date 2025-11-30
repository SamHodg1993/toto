package projects

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	jiraUrlString string = ""
	projectID     int    = 0
)

var SetProjectsJiraUrl = &cobra.Command{
	Use:   "project-set-jira-url",
	Short: "Update a single projects jira url. Defaults to current project",
	Long:  "Update a single projects jira url. e.g. `https://mycompany.atlassian.net`. Defaults to current project.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		jiraUrlString = strings.TrimSpace(jiraUrlString)

		if projectID < 1 {
			projectId, err := ProjectService.GetProjectIdByFilepath()
			if err != nil {
				fmt.Println(err)
				return
			}

			projectID = projectId
		}

		if jiraUrlString == "" {
			fmt.Println("Please enter a jira URL. e.g. `https://mycompany.atlassian.net`")
			return
		}

		err := ProjectService.UpdateProjectsJiraUrl(projectID,
			jiraUrlString)
		if err != nil {
			fmt.Printf("Error setting project jira_url. err: %v\n",
				err)
			return
		}

		fmt.Println("Updated this projects Jira URL")
	},
}

func init() {
	SetProjectsJiraUrl.Flags().StringVarP(&jiraUrlString, "jira-url", "u", "", "Jira URL for this project.")
	SetProjectsJiraUrl.Flags().IntVarP(&projectID, "project", "p", 0, "Project ID which is to be updated.")
}
