package main

import (
	"fmt"
	"os"

	//"fmt"
	"datascore/challenge/database"
	"datascore/challenge/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var App *fiber.App

func main() {
	err := godotenv.Load()
	database.ConnectDb()

	if err != nil {
		fmt.Println("Error loading .env file")
	}
	App := fiber.New()
	//  read from Bucket in Google Cloud Storage.
	App.Get("/google_cloud_bucket_file_csv", handlers.ReadFromGCP_Storage)

	// upload file   tracking files using SHA-256
	App.Post("/upload_and_track_file_csv_SHA_256", handlers.UploadAndTrackCsvSHA256)
	// upload file   tracking files using SHA-256
	App.Post("/upload_and_track_file_csv", handlers.UploadAndTrackCsv)

	// Add list of customers from csv file withou traking
	App.Post("/upload_csv", handlers.AddCustomersFromCsvFile)

	// Get all customers saved on database
	App.Get("/customers", handlers.ListCustomers)
	// Get all Files list saved on database
	App.Get("/files", handlers.ListFiles)
	// Add new customer from Front end post
	App.Post("/customer", handlers.AddNewCustomer)

	App.Get("/delete_all", handlers.DelletAll)

	App.Get("/hello", handlers.SayHello)
	port := os.Getenv("HTTP_SERVER_PORT")
	App.Listen(port)
}
