package service

import (
	"context"
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/wignn/go-with-mongoDb/config"
	"github.com/wignn/go-with-mongoDb/dto"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(c *fiber.Ctx) error {
	var users []dto.User
	collection := config.GetCollection("users")

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error finding users",
			"details": err.Error(),
		})
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user dto.User
		if err := cursor.Decode(&user); err != nil {
			log.Fatalf("Error decoding user: %v", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Error decoding user",
				"details": err.Error(),
			})
		}

		users = append(users, user)

	}

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {

	return nil
}

func UpdateUser(c *fiber.Ctx) error {

	return nil
}

func DeleteUser(c *fiber.Ctx) error {
	return nil
}
