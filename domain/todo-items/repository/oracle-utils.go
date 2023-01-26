package repository

import (
	"golang-todo/utils/database"
)

func (db *oracleTodoItemsRepository) Exec(queryConfig ...database.QueryConfig) (err error) {
	err = database.ExecTransaction(db.sqlDB, queryConfig...)
	return
}
