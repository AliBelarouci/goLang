package database

import (
	"fmt"
	"log"
	"os"

	"datascore/challenge/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// implemetation connection for database

func ConnectDb() {
	// read envernment parametter from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	// generate connection string
	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=5430 sslmode=disable TimeZone=Asia/Shanghai",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// )
	// open connection by using gorm ORM
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_SOURCE")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	// check if the connection is correct
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	fmt.Println("Data base connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	// migrate for generating customer table in server database
	db.AutoMigrate(&models.Customer{})
	// migrate for generating file table in server database
	db.AutoMigrate(&models.File{})

	DB = Dbinstance{
		Db: db,
	}
}
