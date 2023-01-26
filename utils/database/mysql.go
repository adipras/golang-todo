package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func SetupMysqlDatabaseConnection() (db *sql.DB, err error) {
	var (
		driver   = "mysql"
		username = os.Getenv("MYSQL_USER")
		password = os.Getenv("MYSQL_PASSWORD")
		host     = os.Getenv("MYSQL_HOST")
		port     = os.Getenv("MYSQL_PORT")
		name     = os.Getenv("MYSQL_DBNAME")
		location = "loc=Asia%2FJakarta"
	)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&%s", username, password, host, port, name, location)

	db, err = sql.Open(driver, connection)

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(100 * time.Millisecond)

	return
}

func GenerateID(db *sql.DB) (id uint64, err error) {
	query := db.QueryRow(`SELECT UUID_SHORT()`)
	err = query.Scan(&id)
	return
}

func GenerateUUID(db *sql.DB) (uuid string, err error) {
	query := db.QueryRow(`SELECT UUID()`)
	err = query.Scan(&uuid)
	return
}
