package service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/wignn/go-with-mongoDb/config"
	"github.com/wignn/go-with-mongoDb/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBook(c *fiber.Ctx) error {
	book := new(dto.Book)

	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Cannot parse JSON",
			"details": err.Error(),
		})
	}

	collection := config.GetCollection("book")
	result, err := collection.InsertOne(context.Background(), book)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Cannot insert book",
			"details": err.Error(),
		})
	}

	book.ID = result.InsertedID.(primitive.ObjectID)
	return c.Status(200).JSON(result)
}

func GetBooks(c *fiber.Ctx) error {
	var books []dto.Book
	collection := config.GetCollection("book")

	cursor, err := collection.Find(context.Background(), nil)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Error finding books",
			"details": err.Error(),
		})
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var book dto.Book
		if err := cursor.Decode(&book); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error":   "Error decoding book",
				"details": err.Error(),
			})
		}

		books = append(books, book)
	}

	return c.Status(200).JSON(books)
}


func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Cannot parse ID",
			"details": err.Error(),
		})
	}
	collection := config.GetCollection("book")
	result,err := collection.DeleteOne(context.Background(),primitive.M{"_id":objectId})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Cannot delete book",
			"details": err.Error(),
		})
	}

	if(result.DeletedCount==0){
		return c.Status(404).JSON(fiber.Map{
			"error":   "Book not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}