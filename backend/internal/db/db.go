package db

import (
	"comp4913-backend/internal/models"
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

func New(config ...string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v port=%v dbname=%v", config[0], config[1], config[2], config[3], config[4])
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.Vote{},
	); err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func GetDBConfig() []string {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPw := os.Getenv("POSTGRES_PASSWORD")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB_NAME")
	return []string{dbHost, dbUser, dbPw, dbPort, dbName}
}
