package Models

type User struct {
	ID uint `gorm:"AUTO_INCREMENT; primaryKey;"`
	Name string `gorm:"default:Unknown; not null;"`
	Age uint `gorm:"default:18; not null;"`
}

func (b *User) TableName() string {
	return "user"
}
