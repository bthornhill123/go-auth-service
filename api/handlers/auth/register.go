package auth

import (
	"github.com/bthornhill123/go-auth-service/database"
	"github.com/bthornhill123/go-auth-service/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	var existingUser models.User
	database.DB.Where("email = ?", data["email"]).First(&existingUser)
	if existingUser.ID != 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "email previously registered"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 10)
	user := models.User{
		Email:    data["email"],
		Password: string(hashedPassword),
	}
	database.DB.Create(&user)
	return c.JSON(user)
}
