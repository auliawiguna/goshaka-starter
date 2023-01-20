package database

import (
	"errors"
	model "goshaka/app/models"
	appConfig "goshaka/configs"
	appHelper "goshaka/helpers"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error
	dbConnURL, _ := appHelper.ConnectionURLBuilder(appConfig.GetEnv("DB_DRIVER"))

	switch appConfig.GetEnv("DB_DRIVER") {
	case "mysql":
		DB, err = gorm.Open(mysql.Open(dbConnURL), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})
	case "postgres":
		DB, err = gorm.Open(postgres.Open(dbConnURL), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})
	}

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&model.Note{})
	if err = DB.AutoMigrate(&model.User{}); err == nil && DB.Migrator().HasTable(&model.User{}) {
		if err := DB.First(&model.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			DB.Create(&model.User{
				Username:    "aulia",
				Email:       "aulia@goshaka.id",
				Password:    "shaka321",
				ValidatedAt: time.Date(2023, 1, 1, 10, 10, 10, 0, time.UTC),
			})
		}
	}

	return nil
}
