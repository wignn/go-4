package service

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wignn/go-with-mongoDb/config"
	"github.com/wignn/go-with-mongoDb/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *fiber.Ctx) error {
	var users []dto.User
	collection := config.GetCollection("users")

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Error finding users",
			"details": err.Error(),
		})
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user dto.User
		if err := cursor.Decode(&user); err != nil {
			log.Fatalf("Error decoding user: %v", err)
			return c.Status(500).JSON(fiber.Map{
				"error":   "Error decoding user",
				"details": err.Error(),
			})
		}

		users = append(users, user)

	}

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(dto.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Cannot parse JSON",
			"details": err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Error hashing password",
			"details": err.Error(),
		})
	}
	user.Password = string(hashedPassword)
	collection := config.GetCollection("users")
	insertResult,err := collection.InsertOne(context.Background(), user)


	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Error creating user",
			"details": err.Error(),
		})
	}

	user.ID = insertResult.InsertedID.(primitive.ObjectID)
	return c.Status(200).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {

	return nil
}

func DeleteUser(c *fiber.Ctx) error {
	return nil
}
