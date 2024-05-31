package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func initDB(){
	sql.Open("sqlite3", "api.db")
}