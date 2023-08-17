package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"encoding/csv"

	"cloud.google.com/go/storage"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

func ReadFromGCP_Storage(c *fiber.Ctx) error {
	ctx := context.Background()
	// Read the JSON credentials file
	jsonFile, err := os.Open("./credentials.json")
	if err != nil {
		log.Fatalf("Failed to open the credentials file: %v", err)
	}
	defer jsonFile.Close()
	// Create a buffer to store the contents of the file
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, jsonFile)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	fileBytes := buf.Bytes()
	// Get the credentials from the JSON file
	creds, err := google.CredentialsFromJSON(ctx, fileBytes, storage.ScopeReadOnly)
	if err != nil {
		log.Fatalf("Failed to get credentials from the JSON file: %v", err)
	}
	// Create a new storage client
	client, err := storage.NewClient(ctx, option.WithCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to create a new storage client: %v", err)
	}

	// Get a reference to the desired bucket
	bucket := client.Bucket("alibucket123")

	// Get a reference to the desired file in the bucket
	obj := bucket.Object("test-1.csv")

	// Get a reader for the file
	reader, err := obj.NewReader(ctx)
	if err != nil {
		log.Fatalf("Failed to get a reader for the file: %v", err)
	}
	defer reader.Close()
	// Use the csv package to read the file
	csvReader := csv.NewReader(reader)
	records, err := csvReader.ReadAll()
	if err != nil {
		c.Status(http.StatusBadRequest).SendString(err.Error())
		return err
	}
	//remove the header of csv file (first row)
	records = records[1:]
	var records_counts int64 = 0
	var ImportedRowCount int64 = 0
	for _, record := range records {
		records_counts++
		dbc := AddNewCustomerFromCsvRecord(record)
		if dbc.Error == nil {
			ImportedRowCount++
		}
	}
	return c.SendString(fmt.Sprintf("the file test-1.csv uploaded from bucket in GCP with %d total rows and %d rows imported",
		uint64(records_counts),
		uint64(ImportedRowCount),
	))

}
