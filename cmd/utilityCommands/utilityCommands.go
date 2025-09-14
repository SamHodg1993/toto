package utilityCommands

import (
	"database/sql"
	"github.com/samhodg1993/toto-todo-cli/internal/service"
)

var UtilityService *service.UtilityCommandsService

func SetDatabase(database *sql.DB) {
	UtilityService = service.NewUtilityCommandsService(database)
}