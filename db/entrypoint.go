package db

import (
	"github.com/breathbath/go_utils/utils/env"
	baseDb "github.com/breathbath/go_utils/utils/sqlDb"
)

func NewDb() (dg *baseDb.DbGateway, err error) {
	dbConnStr, err := env.ReadEnvOrError("DB_CONN_STRING")
	if err != nil {
		return
	}

	db, err := baseDb.NewDb(dbConnStr, "mysql")
	if err != nil {
		return
	}

	dg = baseDb.NewDbGateway(db)
	return
}
