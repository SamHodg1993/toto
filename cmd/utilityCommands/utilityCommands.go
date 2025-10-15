package utilityCommands

import (
	"database/sql"

	"github.com/samhodg1993/toto/internal/service/utility"
)

var UtilityService *utility.Service

func SetDatabase(database *sql.DB) {
	UtilityService = utility.New(database)
}

