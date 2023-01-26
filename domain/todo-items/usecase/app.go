package usecase

import (
	todoItems "golang-todo/domain/todo-items"
)

/*
*
the struct of usecases
*/
type todoItemsUsecase struct {
	mysqlRepository todoItems.MysqlRepository
}

/*
*
the initiator function for usecase
*/
func NewTodoItemsUsecase(mysql todoItems.MysqlRepository) todoItems.Usecase {
	return &todoItemsUsecase{
		mysqlRepository: mysql,
	}
}
