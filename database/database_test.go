package database_test

import (
	"datascore/challenge/models"
	//"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbinstanceTest struct {
	DbTesting *gorm.DB
}

var DBTest DbinstanceTest

func TestConnection(t *testing.T) {
	err := godotenv.Load("./../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=5430 sslmode=disable TimeZone=Asia/Shanghai",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// )
	// open connection by using gorm ORM
	DBTest = DbinstanceTest{}
	DBTest.DbTesting, err = gorm.Open(postgres.Open(os.Getenv("DB_SOURCE")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	DBTest.DbTesting.Logger = logger.Default.LogMode(logger.Info)
	// migrate for generating customer table in server database
	if err := DBTest.DbTesting.AutoMigrate(&models.Customer{}); err != nil {
		log.Fatal("Failed to cretate Table Customer. \n", err)
	}
	// migrate for generating file table in server database
	if err := DBTest.DbTesting.AutoMigrate(&models.File{}); err != nil {
		log.Fatal("Failed to cretate Table File. \n", err)
	}

	//os.Exit(m.Run())
}
func TestFile(t *testing.T) {
	file := new(models.File)
	//file.Filename = "test_1.csv"
	file.ImportedDate = "10/10/22 12:14:56"
	file.TotalRowCount = 7
	file.ImportedRowCount = 5
	DBTest.DbTesting.Create(&file)
	var fileCreated models.File
	DBTest.DbTesting.Find(&fileCreated, file.ID)
	require.NotEmpty(t, file)
	require.NotZero(t, fileCreated.ID)
	require.NotZero(t, fileCreated.CreatedAt)
	require.Equal(t, fileCreated.Filename, file.Filename)
	require.Equal(t, fileCreated.ImportedDate, file.ImportedDate)
	require.Equal(t, fileCreated.TotalRowCount, file.TotalRowCount)
	require.Equal(t, fileCreated.ImportedRowCount, file.ImportedRowCount)
	require.GreaterOrEqual(t, fileCreated.TotalRowCount, fileCreated.ImportedRowCount)
}
func TestCreatCustomer(t *testing.T) {
	customer := new(models.Customer)
	customer.CallTime = "10/10/22 12:14:56"
	customer.CallDisposition = "AA"
	customer.Phone = "0658695528"
	customer.FirstName = "Jone"
	customer.LastName = "Andres"
	customer.Address1 = "Any address1"
	customer.Address2 = "Any address2"
	customer.State = "nivada"
	customer.Zip = "A69008"

	DBTest.DbTesting.Create(&customer)

	var customerCreated models.Customer
	DBTest.DbTesting.First(&customerCreated, customer.ID)
	require.NotEmpty(t, customer)
	require.NotZero(t, customerCreated.ID)
	require.NotZero(t, customerCreated.CreatedAt)
	require.Equal(t, customer.CallTime, customerCreated.CallTime)
	require.Equal(t, customer.CallDisposition, customerCreated.CallDisposition)
	require.Equal(t, customer.Phone, customerCreated.Phone)
	require.Equal(t, customer.FirstName, customerCreated.FirstName)
	require.Equal(t, customer.LastName, customerCreated.LastName)
	require.Equal(t, customer.Address1, customerCreated.Address1)
	require.Equal(t, customer.Address2, customerCreated.Address2)
	require.Equal(t, customer.State, customerCreated.State)
	require.Equal(t, customer.Zip, customerCreated.Zip)

}
