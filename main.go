package main

import (
	"github.com/kabos0809/gorm_crud/Models"
	"github.com/kabos0809/gorm_crud/Routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	db, err := Models.ConnectDB()
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Models.Article{})

	controller := Models.Model{Db: db}

	router := Router.SetRouter(controller)

	router.Run(":8080")
}