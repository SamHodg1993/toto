package todo

import (
	"fmt"
	"strconv"
	"time"

	"github.com/odgy8/toto/internal/models"
)

func FormatTodoTableRow(todo models.Todo, strikethrough func(a ...interface{}) string) []string {
	title := todo.Title
	if todo.Completed {
		title = strikethrough(title)
	}

	status := "Pending"
	if todo.Completed {
		status = "Done"
	}

	return []string{
		fmt.Sprintf("%d", todo.ID),
		title,
		status,
	}
}

func FormatTodoTableRowLong(todo models.Todo, fullDate bool, strikethrough func(a ...interface{}) string) []string {
	title := todo.Title
	// If todo is completed, apply strikethrough to the title
	if todo.Completed {
		title = strikethrough(title)
	}

	status := "Pending"
	if todo.Completed {
		status = "Done"
	}

	completedAtString := "-"
	if todo.CompletedAt.Valid {
		if fullDate {
			completedAtString = todo.CompletedAt.Time.Format(time.RFC3339)
		} else {
			completedAtString = todo.CompletedAt.Time.Format("02-01-2006")
		}
	}

	updatedAtString := "-"
	if fullDate {
		updatedAtString = todo.UpdatedAt.Format(time.RFC3339)
	} else {
		updatedAtString = todo.UpdatedAt.Format("02-01-2006")
	}

	createdAtString := "-"
	if fullDate {
		createdAtString = todo.CreatedAt.Format(time.RFC3339)
	} else {
		createdAtString = todo.CreatedAt.Format("02-01-2006")
	}

	return []string{
		fmt.Sprintf("%d", todo.ID),
		title,
		todo.Description,
		strconv.Itoa(todo.ProjectId),
		createdAtString,
		updatedAtString,
		status,
		completedAtString,
	}
}
