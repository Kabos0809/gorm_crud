package Models

import "time"

type Article struct {
	ID uint `json:"id" gorm:"AUTO_INCREMENT; primaryKey;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreatetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreatetime"`
	Title string `json:"title" gorm:"default:Unknown; not null;"`
	Text uint `json:"text" gorm:"default:18; not null;"`
}

func (b *Article) TableName() string {
	return "article"
}
