package Router

import (
	"github.com/kabos0809/gorm_crud/Controller"
	"github.com/kabos0809/gorm_crud/Models"

	"github.com/gin-gonic/gin"
)

func SetRouter(controller Models.Model) *gin.Engine{
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/article/list", Controller.Controller{Model: controller}.GetArticle)
		v1.GET("/article/list/:id", Controller.Controller{Model: controller}.GetArticleById)
		v1.POST("/article/create", Controller.Controller{Model: controller}.CreateArticle)
		v1.PUT("/article/update/:id", Controller.Controller{Model: controller}.UpdateArticle)
		v1.DELETE("/article/delete/:id", Controller.Controller{Model: controller}.DeleteArticle)
	}

	return r
}