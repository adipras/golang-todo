package usecase

import (
	"golang-todo/utils/database"
	"golang-todo/valueobject"
)

func (todoItems todoItemsUsecase) ProcessStore(payload valueobject.TodoItemsPayloadInsert) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Data {
		/**
		add data you wanted to insert on this interface{}...
		*/
		data := []interface{}{
			[]interface{}{
				x.ID,
				x.ActivityGroupID,
				x.Title,
				x.IsActive,
				x.Priority,
				x.CreatedAt,
				x.UpdatedAt,
			},
		}

		/**
		column on data and this line should have same order
		*/
		column := []string{
			"todo_id", "activity_group_id", "title", "is_active", "priority", "created_at", "updated_at",
		}

		queryInsert, err := todoItems.mysqlRepository.Store(column, data)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryInsert)
	}
	return
}

func (todoItems todoItemsUsecase) ProcessUpdate(payload valueobject.TodoItemsPayloadUpdate) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Data {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"todo_id": x.Param.ID, // add the parameter to update the row
			},
		}
		var data = map[string]interface{}{
			"updated_at": x.Body.UpdatedAt, // add the data to update the row
			"title":      x.Body.Title,
			"is_active":  x.Body.IsActive,
			"priority":   x.Body.Priority,
		}

		queryUpdate, err := todoItems.mysqlRepository.Update(param, data)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryUpdate)
	}
	return
}

func (todoItems todoItemsUsecase) ProcessDelete(payload valueobject.TodoItemsPayloadDelete) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Param {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"todo_id": x.ID, // add the parameter to delete the row
			},
		}

		queryDelete, err := todoItems.mysqlRepository.Delete(param)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryDelete)
	}
	return
}
