package sqlite

import (
	"database/sql"
	"go-notification-api/internal/logs"
)

func OpenConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./notifications.db")
	if err != nil {
		panic(err)
	}
	return db
}

func CloseConnection(Db *sql.DB) {
	err := Db.Close()
	if err != nil {
		logs.DbLogError(err)
		return
	}
}
