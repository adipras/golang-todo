package main

import (
	"github.com/gin-gonic/gin"

	"database/sql"
	"log"
	"os"

	"golang-todo/utils/database"

	_boilerplateHttpDeliver "golang-todo/domain/boilerplate/delivery/http"
	_boilerplateRepository "golang-todo/domain/boilerplate/repository"
	_boilerplateUsecase "golang-todo/domain/boilerplate/usecase"

	_activityGroupsHttpDeliver "golang-todo/domain/activity-groups/delivery/http"
	_activityGroupsRepository "golang-todo/domain/activity-groups/repository"
	_activityGroupsUsecase "golang-todo/domain/activity-groups/usecase"

	_todoItemsHttpDeliver "golang-todo/domain/todo-items/delivery/http"
	_todoItemsRepository "golang-todo/domain/todo-items/repository"
	_todoItemsUsecase "golang-todo/domain/todo-items/usecase"
)

func main() {
	routers := gin.Default()

	mysql := ConnectMySQL()
	//orcl:= ConnectOracle()

	boilerplateMysqlRepository := _boilerplateRepository.NewMysqlBoilerplateRepository(mysql)
	boilerplateUsecase := _boilerplateUsecase.NewBoilerplateUsecase(boilerplateMysqlRepository)
	_boilerplateHttpDeliver.NewBoilerplateHttpHandler(boilerplateUsecase, routers)

	activityGroupsMysqlRepository := _activityGroupsRepository.NewMysqlActivityGroupsRepository(mysql)
	activityGroupsUsecase := _activityGroupsUsecase.NewActivityGroupsUsecase(activityGroupsMysqlRepository)
	_activityGroupsHttpDeliver.NewActivityGroupsHttpHandler(activityGroupsUsecase, routers)

	todoItemsMysqlRepository := _todoItemsRepository.NewMysqlTodoItemsRepository(mysql)
	todoItemsUsecase := _todoItemsUsecase.NewTodoItemsUsecase(todoItemsMysqlRepository)
	_todoItemsHttpDeliver.NewTodoItemsHttpHandler(todoItemsUsecase, routers)

	routers.Run(":" + os.Getenv("PORT"))
}

func ConnectMySQL() (mysql *sql.DB) {
	mysql, err := database.SetupMysqlDatabaseConnection()

	if err != nil {
		log.Fatal(err.Error())
	}

	return
}
