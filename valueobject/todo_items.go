package valueobject

import "golang-todo/entity"

type TodoItems struct {
	entity.TodoItems
	entity.StandardKey
	entity.Time
}

type TodoItemsPayloadInsert struct {
	Data []TodoItems `json:"data" binding:"required"`
	User string
}

type TodoItemsPayloadUpdate struct {
	Data []TodoItemsDataUpdate `json:"data" binding:"required"`
	User string
}

type TodoItemsDataUpdate struct {
	Param TodoItems `json:"param" binding:"required"`
	Body  TodoItems `json:"body" binding:"required"`
}

type TodoItemsPayloadDelete struct {
	Param []TodoItems `json:"param" binding:"required"`
}
