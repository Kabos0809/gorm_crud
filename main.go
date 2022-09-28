package main

import (
	"github.com/kabos0809/gorm_crud/Models"
	"github.com/kabos0809/gorm_crud/Routes"
	"github.com/kabos0809/gorm_crud/Database"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/gin-gonic/gin"
	
	"fmt"
)

func main() {
	gin.SetMode(gin.DebugMode)
	dsn := Database.DbUrl()
	fmt.Printf("%s", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Models.Article{})

	controller := Models.Model{Db: db}

	router := Routes.SetRouter(controller)

	router.Run(":8080")
}