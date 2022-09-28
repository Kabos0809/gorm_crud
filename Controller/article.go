package Controller

import (
	"net/http"
	"strconv"
	
	"github.com/kabos0809/gorm_crud/Models"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Model Models.ModelInterface
}

//Create Article
func (c Controller) CreateArticle(context *gin.Context) {
	var article Models.Article
	if err := context.BindJSON(&article); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status":"BadRequest"})
		return
	}
	err := c.Model.CreateArticle(&article)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{})
	}
}

//Fetch all article
func (c Controller) GetArticle(context *gin.Context) {
	r, err := c.Model.GetArticle()
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, r)
	}
}

//Update article
func (c Controller) UpdateArticle(context *gin.Context) {
	var article *Models.Article
	id := context.Params.ByName("id")
	idUint, _:= strconv.ParseUint(id, 10, 64)
	r, err := c.Model.GetArticleById(idUint)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		return
	}
	context.BindJSON(&article)
	r.Title = article.Title
	r.Text = article.Text
	err = c.Model.UpdateArticle(r)
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, r)
	}
}

//Delete article
func (c Controller) DeleteArticle(context *gin.Context) {
	id := context.Params.ByName("id")
	idUint, _:= strconv.ParseUint(id, 10, 64)
	err := c.Model.DeleteArticle(idUint)
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}