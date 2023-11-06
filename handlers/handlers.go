package handler

import (
	"rest_api/database"
	"rest_api/models"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, world!")
}

func InsertProduct(c *fiber.Ctx) error {
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error inserting product",
			"error":   err.Error(),
		})
	}

	if err := database.Db.DB.Create(&product).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error inserting product",
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(product)
}

func FetchAllProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	err := database.Db.DB.Find(&products).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error Fetching all product",
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success Fetching all product",
		"error":   "No error",
		"data":    products,
	})
}

func FetchProduct(c *fiber.Ctx) error {
	product := models.Product{}

	idParams := c.Params("id")

	if err := database.Db.DB.Where("id = ?", idParams).Take(&product).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error Fetching product",
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success Fetching all product",
		"error":   "No error",
		"data":    product,
	})

}

func DeleteProduct(c *fiber.Ctx) error {
	product := models.Product{}

	idParams := c.Params("id")

	go func() {
		database.Db.DB.Delete(&product, idParams)
	}()

	return c.Status(200).JSON(fiber.Map{
		"message": "Success Delete Products",
		"error":   "No error",
	})

}
