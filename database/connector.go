package database

import (
	appConfig "goshaka/configs"
	appHelper "goshaka/helpers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error

	dbConnURL, _ := appHelper.ConnectionURLBuilder(appConfig.GetEnv("DB_TYPE"))

	DB, err = gorm.Open(mysql.Open(dbConnURL), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	return nil
}
