package todo

import (
	"fmt"
	"time"

	"github.com/samhodg1993/toto-todo-cli/cmd/projects"

	"github.com/spf13/cobra"
)

var sql_insert_todo string = "INSERT INTO todos (title, description, created_at, updated_at, project_id) VALUES (?,?,?,?,?)"

var (
	todoTitle       string
	todoDescription string
	todoCreatedAt   string
	todoUpdatedAt   string
	todoProjectId   int
)

var AddCmd = &cobra.Command{
	Use:   "add [todo]",
	Short: "Add a new todo",
	Long:  "Add a new todo to the list of stored todos",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// This has moved to the -t flag
		// todo := args[0]

		// Create default timestamps
		createdAt := time.Now()
		updatedAt := time.Now()
		releventProject := 1

		// Does this directory have a project?
		if todoProjectId != 0 {
			releventProject = todoProjectId
		} else {

			row, err := projects.ProjectService.GetProjectIdByFilepath()
			if err != nil {
				if row == 0 {
					choice, err := projects.ProjectService.HandleNoExistingProject()
					if err != nil {
						fmt.Printf("%v.\n", err)
					}

					switch choice {
					case 0:
						return // User has aborted
					case 1:
						releventProject = 1 // This is the global project
					case 2:
						// Create new project and get its ID
						projects.ProjectService.HandleAddNewProject("", "")
						row, err := projects.ProjectService.GetProjectIdByFilepath()
						if err != nil {
							fmt.Printf("Error getting project ID: %v\n", err)
							return
						}
						releventProject = row
					}
				}
			} else {
				releventProject = row
			}
		}

		// If created at flag has value, update the created at value
		if todoCreatedAt != "" {
			parsed, err := time.Parse(time.RFC3339, todoCreatedAt)
			if err == nil {
				createdAt = parsed
			}
		}

		// If updated at flag has value, update the created at value
		if todoUpdatedAt != "" {
			parsed, err := time.Parse(time.RFC3339, todoUpdatedAt)
			if err == nil {
				updatedAt = parsed
			}
		}

		// Update the todo
		err := TodoService.AddTodo(todoTitle, todoDescription, releventProject, createdAt, updatedAt)
		if err != nil {
			fmt.Printf("There was an error adding the todo: %v\n", err)
			return
		}
		fmt.Printf("New todo added: %s\n", todoTitle)
	},
}

func init() {
	AddCmd.PersistentFlags().StringVarP(&todoTitle, "title", "t", "", "Title of the todo")
	AddCmd.PersistentFlags().StringVarP(&todoDescription, "description", "d", "", "Description of the todo")
	AddCmd.PersistentFlags().StringVarP(&todoCreatedAt, "created-at", "c", "", "Todo creation time")
	AddCmd.PersistentFlags().StringVarP(&todoUpdatedAt, "updated-at", "u", "", "Todo last updated time")
	AddCmd.PersistentFlags().IntVarP(&todoProjectId, "project-id", "p", 0, "Relevent Project Id")

	AddCmd.MarkFlagRequired("title")
}
