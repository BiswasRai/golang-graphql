package database

import (
	"fmt"
	"log"
	"os"

	"github.com/biswasRai/golang-graphql/graph/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() *gorm.DB {
	godotenv.Load(".env")

	// Set up the database connection string.
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Open a database connection.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode((logger.Info)),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		os.Exit(2)
	}

	log.Println("Successfully connected to the database!")

	db.AutoMigrate(&model.Todo{})

	DB = Dbinstance{
		Db: db,
	}

	return db

}
