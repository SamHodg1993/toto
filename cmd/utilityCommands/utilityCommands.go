package utilityCommands

import (
	"database/sql"

	"github.com/odgy8/toto/internal/service/utility"
)

var UtilityService *utility.Service

func SetDatabase(database *sql.DB) {
	UtilityService = utility.New(database)
}

