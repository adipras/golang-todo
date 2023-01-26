package repository

import (
	"golang-todo/utils/database"
)

func (db *mysqlTodoItemsRepository) GenerateID() (id uint64, err error) {
	return database.GenerateID(db.sqlDB)
}

func (db *mysqlTodoItemsRepository) GenerateUUID() (uuid string, err error) {
	return database.GenerateUUID(db.sqlDB)
}

func (db *mysqlTodoItemsRepository) Exec(queryConfig ...database.QueryConfig) (err error) {
	err = database.ExecTransaction(db.sqlDB, queryConfig...)
	return
}
