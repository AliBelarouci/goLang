package handlers

import (
	"bytes"
	"io"

	"crypto/sha256"
	"datascore/challenge/database"
	"datascore/challenge/models"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Hash-based comparison: When a file is uploaded, generate a unique hash for it
// (such as MD5 or SHA-256), and store the hash in a database along with the file's
// information. Then, before allowing a new file to be uploaded, compare its hash to the hashes in the
func UploadAndTrackCsvSHA256(c *fiber.Ctx) error {
	// Get the file from the form data
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	//Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	defer src.Close()
	// Create a buffer to store the contents of the file
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, src)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	fileBytes := buf.Bytes()
	hash := sha256.Sum256([]byte(fileBytes))
	hashString := hex.EncodeToString(hash[:])
	var exists bool
	// Find in the database for file have the same Hash
	if err := database.DB.Db.Model(models.File{}).
		Select("count(*) > 0").
		Where("Hash = ?", hashString).
		Find(&exists).Error; err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	// prevent duplicte files... file hase the same Hash
	if exists {
		return c.Status(http.StatusForbidden).SendString(fmt.Sprintf("The file %s already loaded the same content of this file it uploaded", file.Filename))
	}
	// Parse the contents of the file as a CSV
	r := csv.NewReader(buf)
	records, err := r.ReadAll()
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
	fileobj := new(models.File)
	fileobj.Filename = file.Filename
	fileobj.ImportedDate = time.Now().GoString()
	fileobj.TotalRowCount = records_counts
	fileobj.ImportedRowCount = ImportedRowCount
	fileobj.Hash = hashString

	database.DB.Db.Create(&fileobj)
	return c.SendString(fmt.Sprintf("the file %s uploaded with %d total rows and %d rows imported",
		file.Filename,
		uint64(records_counts),
		uint64(ImportedRowCount),
	))

}

// upload file for save recordes in database and  tracking files prevent the same name
func UploadAndTrackCsv(c *fiber.Ctx) error {
	// Get the file from the form data
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	var exists bool
	// Find in the database for file have the same name
	if err := database.DB.Db.Model(models.File{}).
		Select("count(*) > 0").
		Where("Filename = ?", file.Filename).
		Find(&exists).Error; err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	// prevent duplicte files
	if exists {
		return c.Status(http.StatusForbidden).SendString(fmt.Sprintf("The file %s already loaded", file.Filename))
	}
	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	defer src.Close()
	// Create a buffer to store the contents of the file
	buf := new(bytes.Buffer)
	buf.ReadFrom(src)

	// Parse the contents of the file as a CSV
	r := csv.NewReader(buf)
	records, err := r.ReadAll()
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
		//customer := AddNewCustomerFromCsvRecord(record)

		if dbc := AddNewCustomerFromCsvRecord(record); dbc.Error == nil {
			ImportedRowCount++
		}
	}
	fileobj := new(models.File)
	fileobj.Filename = file.Filename
	fileobj.ImportedDate = time.Now().GoString()
	fileobj.TotalRowCount = records_counts
	fileobj.ImportedRowCount = ImportedRowCount

	database.DB.Db.Create(&fileobj)
	return c.SendString(fmt.Sprintf("the file %s uploaded with %d total rows and %d rows imported",
		file.Filename,
		uint64(records_counts),
		uint64(ImportedRowCount),
	))
}

// Add list of customers from csv file withou traking
func AddCustomersFromCsvFile(c *fiber.Ctx) error {
	// Get the file from the form data
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	defer src.Close()
	// Create a buffer to store the contents of the file
	buf := new(bytes.Buffer)
	buf.ReadFrom(src)

	// Parse the contents of the file as a CSV
	r := csv.NewReader(buf)
	records, err := r.ReadAll()
	if err != nil {
		c.Status(http.StatusBadRequest).SendString(err.Error())
		return err
	}
	// remove the header of csv file (first row)
	records = records[1:]
	// Do something with the records
	for _, record := range records {
		AddNewCustomerFromCsvRecord(record)

	}
	return c.SendString("File uploaded successfully")
}
