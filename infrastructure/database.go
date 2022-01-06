package infrastructure

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB       *sqlx.DB
	DBReader *sqlx.DB
)

func init() {
	DB = ConnectDB(true)
	DBReader = ConnectDB(false)
}

func ConnectDB(writer bool) *sqlx.DB {
	username := os.Getenv("POKERROOM_DATABASE_USERNAME")
	password := os.Getenv("POKERROOM_DATABASE_PASSWORD")
	host := os.Getenv("POKERROOM_DATABASE_WRITER_HOST")
	if !writer {
		host = os.Getenv("POKERROOM_DATABASE_READER_HOST")
	}
	database := os.Getenv("POKERROOM_DATABASE")

	dbConfigStr := username + ":" + password + "@tcp(" + host + ":3306)/" + database + "?parseTime=true"
	db, err := sqlx.Connect("mysql", dbConfigStr)
	if err != nil {
		panic(err)
	}

	return db
}
