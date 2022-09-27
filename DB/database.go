package Database

import (
	"fmt"
	"os"
	"gorm.io/gorm"
)

type DbConfig struct{
	DB string
	User string
	Pass string
	Host string
	TZ string
	Port string
}

func buildDBConfig() *DbConfig {
	dbConfig := DbConfig{
		DB = os.Getenv("POSTGRES_DB")
		User = os.Getenv("POSTGRES_USER")
		Pass = os.Getenv("POSTGRES_PASS")
		Host = os.Getenv("POSTGRES_HOST")
		TZ = os.Getenv("TZ")
		Port = os.Getenv("POSTGRES_PORT")
	}
	return &dbConfig
}

func DbUrl(dbConfig *DbConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", 
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.DB,
		dbConfig.Port,
		dbConfig.TZ
	)
}