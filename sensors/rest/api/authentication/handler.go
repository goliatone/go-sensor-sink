package authentication

import (
	"fmt"
	"sensors/auth"

	"github.com/gofiber/fiber"
	"github.com/twinj/uuid"
)

//Login will return a login handler
func Login(domain auth.Interactor) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		var params LoginParams
		_ = ctx.BodyParser(&params)

		if params.Identifier == "" {
			err := ErrResponse(ErrInvalidParams{
				message: fmt.Sprintf("provide valid identifier to sign in"),
			})
			_ = ctx.Status(err.Status).JSON(err)
			return
		}

		signedUser, err := domain.AuthenticateByEmail(params.Identifier, params.Password)
		if err != nil {
			re := ErrResponse(err)
			_ = ctx.Status(re.Status).JSON(re)
			return
		}

		ctx.Status(fiber.StatusOK).JSON(signedUser)
	}
}

//Register will create a new user and add it to DB
func Register(domain auth.Interactor) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		var params RegistrationParams
		_ = ctx.BodyParser(&params)

		_, err := ValidateRegisterParams(&params)
		if err != nil {
			e := ErrResponse(err)
			_ = ctx.Status(e.Status).JSON(e)
			return
		}

		newUser := createUserObject(params)
		u, err := domain.Register(newUser)
		if err != nil {
			errHTTP := ErrResponse(err)
			_ = ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}

		_ = ctx.JSON(auth.RegistrationResponse(&u))

	}
}

func createUserObject(params RegistrationParams) *auth.User {
	var u = auth.User{
		ID:       uuid.NewV4(),
		Email:    params.Email,
		Password: params.Password,
	}
	return &u
}
