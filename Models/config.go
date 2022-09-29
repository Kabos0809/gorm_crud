package Models

import (
	"gorm.io/gorm"
)

var err error

type ModelInterface interface {
	GetArticle() (*[]Article, error)
	GetArticleById(id uint64) (*Article, error)
	CreateArticle(article *Article) error 
	UpdateArticle(article *Article) error
	DeleteArticle(id uint64) error
}

type Model struct {
	Db *gorm.DB
}