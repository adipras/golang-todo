package repository

import (
	"golang-todo/utils/database"
)

func (db *oracleBoilerplateRepository) Exec(queryConfig ...database.QueryConfig) (err error) {
	err = database.ExecTransaction(db.sqlDB, queryConfig...)
	return
}
