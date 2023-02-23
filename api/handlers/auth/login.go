package auth

import (
	"time"

	"github.com/bthornhill123/go-auth-service/database"
	"github.com/bthornhill123/go-auth-service/models"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{"message": "user not found"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))
	if err != nil {
		return c.JSON(fiber.Map{"message": "login unsuccessful"})
	}

	expireTime := time.Now().Add(time.Hour)
	token, signError := GenerateJwtToken(user.Email, expireTime)
	if signError != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "login unsuccessful"})
	}

	SetJwtCookie(c, token, expireTime)

	return c.JSON(fiber.Map{"message": "login successful"})
}

func GenerateJwtToken(email string, expireTime time.Time) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    email,
		ExpiresAt: jwt.At(expireTime),
	})
	return claims.SignedString([]byte(signingKey))
}

func SetJwtCookie(c *fiber.Ctx, token string, expireTime time.Time) {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  expireTime,
		HTTPOnly: true,
	})
}
