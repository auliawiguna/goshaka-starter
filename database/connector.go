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
	if err = DB.AutoMigrate(&model.Role{}); err == nil && DB.Migrator().HasTable(&model.Role{}) {
		if err := DB.First(&model.Role{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			var roles = []model.Role{{
				Name:    "admin",
				Display: "Super Admin (Developer)",
			}, {
				Name:    "user",
				Display: "User",
			}}
			DB.Create(&roles)
		}
	}
	if err = DB.AutoMigrate(&model.Permission{}); err == nil && DB.Migrator().HasTable(&model.Permission{}) {
		if err := DB.First(&model.Permission{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			var permissions = []model.Permission{{
				Name:    "role-create",
				Display: "Create role",
			}, {
				Name:    "role-read",
				Display: "Index role",
			}, {
				Name:    "role-update",
				Display: "Update role",
			}, {
				Name:    "role-delete",
				Display: "Delete role",
			}, {
				Name:    "permission-create",
				Display: "Create permission",
			}, {
				Name:    "permission-read",
				Display: "Index permission",
			}, {
				Name:    "permission-update",
				Display: "Update permission",
			}, {
				Name:    "permission-delete",
				Display: "Delete permission",
			}, {
				Name:    "user-create",
				Display: "Create User",
			}, {
				Name:    "user-read",
				Display: "Index User",
			}, {
				Name:    "user-update",
				Display: "Update User",
			}, {
				Name:    "user-delete",
				Display: "Delete User",
			}}
			DB.Create(&permissions)
		}
	}
	if err = DB.AutoMigrate(&model.RoleUser{}); err == nil && DB.Migrator().HasTable(&model.RoleUser{}) {
		if err := DB.First(&model.RoleUser{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			DB.Create(&model.RoleUser{
				UserId: 1,
				RoleId: 1,
			})
		}
	}
	if err = DB.AutoMigrate(&model.PermissionRole{}); err == nil && DB.Migrator().HasTable(&model.PermissionRole{}) {
		if err := DB.First(&model.PermissionRole{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			var permissions []*model.Permission
			DB.Find(&permissions)
			for _, p := range permissions {
				DB.Create(&model.PermissionRole{
					PermissionId: p.ID,
					RoleId:       1,
				})

			}
		}
	}
	DB.AutoMigrate(&model.UserToken{})
	DB.AutoMigrate(&model.ChangeEmail{})

	return nil
}
