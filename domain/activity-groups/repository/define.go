package repository

import (
	"database/sql"

	activityGroups "golang-todo/domain/activity-groups"
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
const MYSQL_TABLE = "activities"
const ORACLE_TABLE = "activities"

/*
*
the struct of mysql
*/
type mysqlActivityGroupsRepository struct {
	sqlDB *sql.DB
}

/*
*
the initiator function for mysql
*/
func NewMysqlActivityGroupsRepository(databaseConnection *sql.DB) activityGroups.MysqlRepository {
	return &mysqlActivityGroupsRepository{databaseConnection}
}

/*
*
the struct of oracle
*/
type oracleActivityGroupsRepository struct {
	sqlDB *sql.DB
}

/*
*
the initiator function for oracle
*/
func NewOracleActivityGroupsRepository(databaseConnection *sql.DB) activityGroups.OracleRepository {
	return &oracleActivityGroupsRepository{databaseConnection}
}
