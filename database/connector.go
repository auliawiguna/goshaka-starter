package database

import (
	model "goshaka/app/models"
	appConfig "goshaka/configs"
	appHelper "goshaka/helpers"

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

	return nil
}
