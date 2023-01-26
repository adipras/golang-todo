package repository

import (
	"golang-todo/utils/database"
)

func (db *mysqlBoilerplateRepository) GenerateID() (id uint64, err error) {
	return database.GenerateID(db.sqlDB)
}

func (db *mysqlBoilerplateRepository) GenerateUUID() (uuid string, err error) {
	return database.GenerateUUID(db.sqlDB)
}

func (db *mysqlBoilerplateRepository) Exec(queryConfig ...database.QueryConfig) (err error) {
	err = database.ExecTransaction(db.sqlDB, queryConfig...)
	return
}
