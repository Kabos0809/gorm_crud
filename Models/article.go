package Models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
//CRUD処理

//Insert new article
func (m Model) CreateArticle(article *Article) error {
	tx := m.Db.Begin()
	err = tx.Create(article).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

//Fetch all article
func (m Model) GetArticle() ([]*Article, error) {
	var article []*Article
	tx := m.Db.Begin()
	err = tx.Find(article).Error
	if err != nil {
		tx.Rollback()
		return article, err
	}
	tx.Commit()
	return article, err
}

//Fetch article by id
func (m Model) GetArticleById(id uint) (*Article, error) {
	var article Article
	tx := m.Db.Begin()
	err = tx.Where("id = ?", id).Find(&article).Error
	if err != nil {
		tx.Rollback()
		return &article, err
	}
	tx.Commit()
	return &article, err
}

//Update article
func (m Model) UpdateArticle(article *Article) error {
	tx := m.Db.Begin()
	err := tx.Save(article).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

//Delete article
func (m Model) DeleteArticle(id uint) error {
	tx := m.Db.Begin()
	err := tx.Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}