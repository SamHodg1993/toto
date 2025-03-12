package todo

import (
	"fmt"

	"github.com/samhodg1993/toto-todo-cli/cmd/projects"

	"github.com/spf13/cobra"
)

var RemoveCompleteForProject = &cobra.Command{
	Use:   "remove-complete",
	Short: "Remove all the completed todos for the current project.",
	Long:  "Remove all of the completed todos for the project assigned to the current directory",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := projects.ProjectService.GetProjectIdByFilepath()
		if err != nil {
			fmt.Printf("%v.\n", err)
		}

		err = TodoService.RemoveCompletedTodosForProject(id)
		if err != nil {
			fmt.Printf("%v.\n", err)
		}
	},
}

var RemoveCompForProj = &cobra.Command{
	Use:   "cls-comp",
	Short: "Remove all the completed todos for the current project.",
	Long:  "Remove all of the completed todos for the project assigned to the current directory",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := projects.ProjectService.GetProjectIdByFilepath()
		if err != nil {
			fmt.Printf("%v.\n", err)
		}

		err = TodoService.RemoveCompletedTodosForProject(id)
		if err != nil {
			fmt.Printf("%v.\n", err)
		}
	},
}
