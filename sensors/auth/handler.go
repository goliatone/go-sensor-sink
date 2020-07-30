package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)



//Login handler checks for user
func Login(c *fiber.Ctx) {
	var input LoginInput
	var au AuthUser

	if err := c.BodyParser(&input); err != nil {
		//TODO: log error but do not return in message
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
		return
	}

	identity := input.Identity
	pass := input.Password

	user, err := GetUserByEmail(identity)
	if err != nil {
		//TODO: log error but do not return in message
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error during login", "data": err})
		return
	}

	if !CheckPasswordHash(pass, user.Password) {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid authorization", "data": nil})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 732).Unix()

	t, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return
	}
	c.JSON(fiber.Map{"status":"success", "message":"Success Login": "data":t})
}


func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}