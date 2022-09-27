package Controller

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/kabos0809/gorm_crud/Models"
	"github.com/kabos0809/gorm_crud/Database"
)

func ConnectDB() *gorm.DB {
	db := Database.buildDBConfig()
	dsn := Database.DbUrl(db)
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return Db
}

//CRUD処理
func CreateUser() {
	db := ConnectDB()
	user := Models.User{Name: ""}
}