package todoitems

import (
	"golang-todo/utils/database"
	"golang-todo/valueobject"
)

/*
*
why there's only one usecase interface while there can more than one repository interface?...
... because, at DDD (Domain Design Driven), there's only one set of usecase and...
... the function name inside the usecase should be unique and represent the business process...
... tl;dr: function name is telling what exactly are they doing.
*/
type Usecase interface {
	GetAll(param map[string]interface{}) ([]valueobject.TodoItems, error)
	GetOne(param map[string]interface{}) (valueobject.TodoItems, error)

	Store(payload valueobject.TodoItems) (valueobject.TodoItems, error)
	Update(payload valueobject.TodoItems) error
	Delete(payload valueobject.TodoItems) error

	ProcessStore(payload valueobject.TodoItemsPayloadInsert) ([]database.QueryConfig, error)
	ProcessUpdate(payload valueobject.TodoItemsPayloadUpdate) ([]database.QueryConfig, error)
	ProcessDelete(payload valueobject.TodoItemsPayloadDelete) ([]database.QueryConfig, error)
}
