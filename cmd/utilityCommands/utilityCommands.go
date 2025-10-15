package utilityCommands

import (
	"database/sql"

	"github.com/samhodg1993/toto/internal/service"
)

var UtilityService *service.UtilityCommandsService

func SetDatabase(database *sql.DB) {
	UtilityService = service.NewUtilityCommandsService(database)
}

