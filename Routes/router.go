package Routes

import (
	"github.com/kabos0809/gorm_crud/Controllers"
	"github.com/kabos0809/gorm_crud/Models"

	"github.com/gin-gonic/gin"
)

func SetRouter(controller Models.Model) *gin.Engine{
	r := gin.Default()
	v1 := r.Group("/v1")
	C := Controllers.Controller{Model: controller}
	{
		v1.GET("/article/list", C.GetArticle)
		v1.GET("/article/detail/:id", C.GetArticleById)
		v1.POST("/article/create", C.CreateArticle)
		v1.PUT("/article/update/:id", C.UpdateArticle)
		v1.DELETE("/article/delete/:id", C.DeleteArticle)
	}

	return r
}