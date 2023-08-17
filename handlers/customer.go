package handlers

import (
	"datascore/challenge/database"
	"datascore/challenge/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SayHello(c *fiber.Ctx) error {

	return c.SendString("Hello, World!")
}

func AddNewCustomerFromCsvRecord(record []string) *gorm.DB {
	customer := new(models.Customer)
	customer.CallTime = record[0]
	customer.CallDisposition = record[1]
	customer.Phone = record[2]
	customer.FirstName = record[3]
	customer.LastName = record[4]
	customer.Zip = record[5]
	return database.DB.Db.Create(&customer)
	//return customer
}
func ListCustomers(c *fiber.Ctx) error {
	customers := []models.Customer{}
	database.DB.Db.Find(&customers)
	return c.Status(200).JSON(customers)
}
func ListFiles(c *fiber.Ctx) error {
	files := []models.File{}
	database.DB.Db.Find(&files)
	return c.Status(200).JSON(files)
}

// Add new customer from Front end post
func AddNewCustomer(c *fiber.Ctx) error {
	customers := new(models.Customer)
	if err := c.BodyParser(customers); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&customers)
	return c.Status(200).JSON(customers)
}
func DelletAll(c *fiber.Ctx) error {

	database.DB.Db.Where("1=1").Delete(&models.File{})
	database.DB.Db.Where("1=1").Delete(&models.Customer{})
	return c.SendString("Delete All")
}
