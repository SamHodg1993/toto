package project

import "fmt"

// HandleNoExistingProject prompts the user for actions when no project exists
func (s *Service) HandleNoExistingProject() (int, error) {
	var cancel string

	fmt.Println(`There is currently no project for this filepath.
			Would you like to
			0 - Cancel
			1 - Add to the global todo list?
			OR
			2 - Create a new project for this filepath?`)
	fmt.Scanf("%s", &cancel)
	if cancel == "1" {
		return 1, nil
	} else if cancel == "2" {
		return 2, nil
	} else {
		fmt.Println("Aborting.")
		return 0, fmt.Errorf("operation cancelled by user")
	}
}
