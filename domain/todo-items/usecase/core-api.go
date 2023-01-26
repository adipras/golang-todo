package usecase

import (
	"golang-todo/entity"
	"golang-todo/valueobject"
	"time"
)

func (todoItems todoItemsUsecase) GetAll(param map[string]interface{}) (response []valueobject.TodoItems, err error) {
	response, err = todoItems.mysqlRepository.GetAll(param)
	return
}

func (todoItems todoItemsUsecase) GetOne(param map[string]interface{}) (response valueobject.TodoItems, err error) {
	response, err = todoItems.mysqlRepository.GetOne(param)
	return
}

func (todoItems todoItemsUsecase) Store(payload valueobject.TodoItems) (valueobject.TodoItems, error) {
	now := time.Now()
	var storeData valueobject.TodoItemsPayloadInsert

	payload.ID, _ = todoItems.mysqlRepository.GenerateID()
	payload.CreatedAt = &now
	payload.UpdatedAt = &now
	if payload.Priority == "" {
		payload.Priority = "very-high"
	}
	storeData.Data = append(storeData.Data, payload)

	queryConfig, err := todoItems.ProcessStore(storeData)

	if err != nil {
		return payload, err
	}

	return payload, todoItems.mysqlRepository.Exec(queryConfig...)
}

func (todoItems todoItemsUsecase) Update(payload valueobject.TodoItems) (err error) {
	now := time.Now()
	var updateData valueobject.TodoItemsPayloadUpdate
	updateData.Data = append(updateData.Data, valueobject.TodoItemsDataUpdate{
		Param: valueobject.TodoItems{
			StandardKey: entity.StandardKey{
				ID: payload.ID,
			},
		},
		Body: valueobject.TodoItems{
			TodoItems: entity.TodoItems{
				TodoItemsKey: entity.TodoItemsKey{
					ActivityGroupID: payload.ActivityGroupID,
				},
				Title:    payload.Title,
				IsActive: payload.IsActive,
				Priority: payload.Priority,
			},
			Time: entity.Time{
				UpdatedAt: &now,
			},
		},
	})

	queryConfig, err := todoItems.ProcessUpdate(updateData)

	if err != nil {
		return
	}

	return todoItems.mysqlRepository.Exec(queryConfig...)
}

func (todoItems todoItemsUsecase) Delete(payload valueobject.TodoItems) (err error) {
	var deleteData valueobject.TodoItemsPayloadDelete
	deleteData.Param = append(deleteData.Param, valueobject.TodoItems{
		StandardKey: entity.StandardKey{
			ID: payload.ID,
		},
	})
	queryConfig, err := todoItems.ProcessDelete(deleteData)

	if err != nil {
		return
	}

	return todoItems.mysqlRepository.Exec(queryConfig...)
}
