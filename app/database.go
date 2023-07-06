package app

import (
	"database/sql"
	"nozzlium/kepo_backend/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Admin123@tcp(localhost:3306)/kepo_backend?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(3 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	return db
}
