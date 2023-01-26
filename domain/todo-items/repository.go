package todoitems

import (
	"golang-todo/utils/database"
	"golang-todo/valueobject"
)

type MysqlRepository interface {
	Exec(...database.QueryConfig) error
	GenerateID() (uint64, error)
	GenerateUUID() (string, error)

	GetAll(param map[string]interface{}) ([]valueobject.TodoItems, error)
	GetOne(param map[string]interface{}) (valueobject.TodoItems, error)

	Store(column []string, data []interface{}) (database.QueryConfig, error)
	Update(param map[string]interface{}, data map[string]interface{}) (database.QueryConfig, error)
	Delete(param map[string]interface{}) (database.QueryConfig, error)
}

type OracleRepository interface {
	Exec(...database.QueryConfig) error

	GetAll(param map[string]interface{}) (response []valueobject.TodoItems, err error)
	GetOne(param map[string]interface{}) (response valueobject.TodoItems, err error)

	Store(column []string, data []interface{}) (database.QueryConfig, error)
	Update(param map[string]interface{}, data map[string]interface{}) (database.QueryConfig, error)
	Delete(param map[string]interface{}) (database.QueryConfig, error)
}
