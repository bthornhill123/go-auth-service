package auth

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

// TODO - put this environment variable into docker-compose
// var signingKey = os.Get("JWT_SIGN_KEY")
var signingKey = "secret-key"

func Authenticate(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.JSON(fiber.Map{
			"authenticated": "false",
		})
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.Status(fiber.StatusUnauthorized)
		} else {
			c.Status(fiber.StatusBadRequest)
		}

		return c.JSON(fiber.Map{
			"authenticated": "false",
		})
	}

	if !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"authenticated": "false",
		})
	}

	return c.JSON(fiber.Map{
		"authenticated": "true",
	})
}

// Middleware
// func IsAuthorized(endpoint func(c *fiber.Ctx)) error {
// 	tokenString := "tokenString"
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
// 		_, ok := token.Method
// 		if !ok {
// 			return nil, fmt.Errorf("Could not authorize")
// 		}
// 		return signingKey, nil
// 	})

// 	if err != nil {
// 		return c.JSON(fiber.Map{
// 			"authenticated": "false",
// 		})
// 	}

// 	if token.Valid {
// 		endpoint(c)
// 	}
// }
