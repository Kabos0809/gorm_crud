package Database

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
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
	err := godotenv.Load("../env")
	if err != nil {
		fmt.Println(err)	
	}
	dbConfig := DbConfig{
		DB: os.Getenv("POSTGRES_DB"),
		User: os.Getenv("POSTGRES_USER"),
		Pass: os.Getenv("POSTGRES_PASSWORD"),
		Host: "localhost",
		TZ: os.Getenv("TZ"),
		Port: "5432",
	}
	
	return &dbConfig
}

func DbUrl() string {
	dbConfig := buildDBConfig()
	return fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=disable TimeZone=%s", 
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.DB,
		dbConfig.Port,
		dbConfig.TZ,
	)
}