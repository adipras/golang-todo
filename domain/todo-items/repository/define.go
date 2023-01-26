package repository

import (
	"database/sql"

	todoItems "golang-todo/domain/todo-items"
)

/*
*
default const that you should believe and please don't change the value
*/
const MYSQL = "mysql"
const ORACLE = "oracle"

/*
*
default const that you should believe and please don't change the value
*/
const SELECT = "select"
const INSERT = "insert"
const UPDATE = "update"
const DELETE = "delete"

/*
*
two lines below is the const of table name, please change the value of this const
*/
const MYSQL_TABLE = "todos"
const ORACLE_TABLE = "todos"

/*
*
the struct of mysql
*/
type mysqlTodoItemsRepository struct {
	sqlDB *sql.DB
}

/*
*
the initiator function for mysql
*/
func NewMysqlTodoItemsRepository(databaseConnection *sql.DB) todoItems.MysqlRepository {
	return &mysqlTodoItemsRepository{databaseConnection}
}

/*
*
the struct of oracle
*/
type oracleTodoItemsRepository struct {
	sqlDB *sql.DB
}

/*
*
the initiator function for oracle
*/
func NewOracleTodoItemsRepository(databaseConnection *sql.DB) todoItems.OracleRepository {
	return &oracleTodoItemsRepository{databaseConnection}
}
