package repository

import (
	"context"
	"nozzlium/kepo_backend/data/entity"
	"nozzlium/kepo_backend/helper"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newTestDB() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:kenshoryureppa@tcp(127.0.0.1:3306)/kepo_backend_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)
	return db
}

var repo NotificationRepository = &NotificationRepositoryImpl{}
var db = newTestDB()

func TestInsert(t *testing.T) {
	notif := entity.Notification{
		UserID:     1,
		QuestionID: 1,
		NotifType:  "ANS",
		Headline:   "test headline",
		Preview:    "test preview",
	}
	res, err := repo.Create(context.Background(), db, notif)

	assert.Nil(t, err)
	assert.NotEqual(t, uint(0), res.ID)
}

func TestRead(t *testing.T) {
	notif := entity.Notification{
		ID:     2,
		IsRead: true,
	}
	res, err := repo.Read(context.Background(), db, notif)

	assert.Nil(t, err)
	assert.True(t, res.IsRead)
}
