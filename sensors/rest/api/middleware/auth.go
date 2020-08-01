package middleware

import (
	"net/http"
	"reflect"
	"sensors/auth"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

//ErrHTTP is an error struct
type ErrHTTP struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

//NewErrHTTP creates a new JSON error from an Error
func NewErrHTTP(err error) ErrHTTP {
	switch err.(type) {
	case auth.ErrUnauthorized:
		return ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
			Message: err.Error(),
			Status:  http.StatusUnauthorized,
		}
	default:
		return ErrHTTP{
			Error:  reflect.TypeOf(err).Name(),
			Mesage: err.Error(),
			Status: http.StatusUnauthorized,
		}
	}
}

//AuthByBearertoken middleware function to auth users
func AuthByBearertoken(secret string) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		header := ctx.Get("Authorization")
		if header == "" {
			err := NewErrHTTP(auth.ErrUnauthorized{Message: "Authorization header not set"})
			_ = ctx.Status(err.Status).JSON(err)
			return
		}

		bearer := strings.Split(header, " ")
		if len(bearer) < 2 || bearer[1] == "" {
			err := NewErrHTTP(auth.ErrUnauthorized{Message: "Authorization header not set"})
			_ = ctx.Status(err.Status).JSON(err)
			return
		}

		var claims auth.TokenClaims
		token, err := auth.ParseToken(bearer[1], secret, &claims)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				errUnauthorized := NewErrHTTP(auth.ErrUnauthorized{Message: "Invalid signature"})
				_ = ctx.Status(errUnauthorized.Status).JSON(errUnauthorized)
				return
			}

			errUnauthorized := NewErrHTTP(auth.ErrUnauthorized{Message: "Expired or invalid token"})
			_ = ctx.Status(errUnauthorized.Status).JSON(errUnauthorized)
			return
		}

		if valid := auth.ValidateToken(token); !valid {
			err := NewErrHTTP(auth.ErrUnauthorized{Message: "Invalid token"})
			_ = ctx.Status(err.Status).JSON(err)
			return
		}

		userDetails := map[string]string{
			"userId": claims.User.UserID,
			"email":  claims.User.Email,
		}

		ctx.Locals("user", userDetails)

		ctx.Next()
	}
}
