package activitygroups

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
	GetAll(param map[string]interface{}) ([]valueobject.ActivityGroups, error)
	GetOne(param map[string]interface{}) (valueobject.ActivityGroups, error)

	Store(payload valueobject.ActivityGroups) (valueobject.ActivityGroups, error)
	Update(payload valueobject.ActivityGroups) error
	Delete(payload valueobject.ActivityGroups) error

	ProcessStore(payload valueobject.ActivityGroupsPayloadInsert) ([]database.QueryConfig, error)
	ProcessUpdate(payload valueobject.ActivityGroupsPayloadUpdate) ([]database.QueryConfig, error)
	ProcessDelete(payload valueobject.ActivityGroupsPayloadDelete) ([]database.QueryConfig, error)
}
