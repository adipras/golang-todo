package http

import (
	"github.com/gin-gonic/gin"

	todoItems "golang-todo/domain/todo-items"
)

type HttpTodoItemsHandler struct {
	todoItemsUsecase todoItems.Usecase
}

func NewTodoItemsHttpHandler(todoItems todoItems.Usecase, httpRouter *gin.Engine) {
	handler := &HttpTodoItemsHandler{
		todoItemsUsecase: todoItems,
	}

	public := httpRouter.Group("/public/api/v1")
	public.GET("/todo-items", handler.GetAll)
	public.GET("/todo-items/:id", handler.GetByID)
	public.POST("/todo-items", handler.Store)
	public.PATCH("/todo-items/:id", handler.Update)
	public.DELETE("/todo-items/:id", handler.Delete)
}
