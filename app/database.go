package app

import (
	"nozzlium/kepo_backend/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewTestDB() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:kenshoryureppa@tcp(127.0.0.1:3306)/kepo_backend_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)
	return db
}

func NewDB() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:Admin123@tcp(127.0.0.1:3306)/kepo_backend?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)
	return db
}

/*
func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Admin123@tcp(localhost:3306)/kepo_backend?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(3 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	return db
}
*/
