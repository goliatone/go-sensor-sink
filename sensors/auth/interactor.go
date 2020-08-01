package auth

import (
	"fmt"
	"log"
	"sensors"

	"golang.org/x/crypto/bcrypt"
)

//Interactor implements interactions
type Interactor interface {
	AuthenticateByEmail(email, password string) (SignedUser, error)
	Register(*User) (User, error)
}

type interactor struct {
	config     sensors.Auth
	repository Repository
}

//NewInteractor interactor
func NewInteractor(config sensors.Auth, authRepo Repository) Interactor {
	return &interactor{
		config:     config,
		repository: authRepo,
	}
}

func (h interactor) AuthenticateByEmail(email, password string) (SignedUser, error) {
	user, err := h.repository.GetByEmail(email)
	if err != nil {
		return SignedUser{}, err
	}

	token, err := h.authenticateUser(&user, password)
	if err != nil {
		return SignedUser{}, err
	}

	return SignedUser{UserID: user.ID.String(), Token: token}, nil
}

func (h interactor) authenticateUser(user *User, password string) (string, error) {
	if isValidPassword := h.comparePasswords(user.Password, []byte(password)); !isValidPassword {
		msg := fmt.Sprintf("user or password invalid")
		return "", ErrUnauthorized{Message: msg}
	}

	tok := generateToken(user)
	token, err := getTokenString(h.config.JWTSecret, tok)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (h interactor) comparePasswords(hash string, password []byte) bool {
	byteHash := []byte(hash)

	err := bcrypt.CompareHashAndPassword(byteHash, password)
	if err != nil {
		log.Printf("err comparing passwords %v", err)
		return false
	}
	return true
}

func (h interactor) Register(user *User) (User, error) {
	passHash, err := h.hashUserPassword([]byte(user.Password))
	if err != nil {
		return User{}, err
	}

	user.Password = passHash
	u, err := h.repository.Add(*user)
	if err != nil {
		return User{}, err
	}

	// h.postNewUserToChannel(&u)
	return u, nil
}

func (h interactor) hashUserPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, 8)
	if err != nil {
		hashErr := ErrHashPassword{password: string(password), message: err.Error()}
		return "", &hashErr
	}

	return string(hash), nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////

//Login interactor checks for user
// func Login(c *fiber.Ctx) {
// 	var input LoginInput
// 	var au User

// 	if err := c.BodyParser(&input); err != nil {
// 		//TODO: log error but do not return in message
// 		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
// 		return
// 	}

// 	identity := input.Identity
// 	pass := input.Password

// 	user, err := repository.GetByUsername(identity)
// 	if err != nil {
// 		//TODO: log error but do not return in message
// 		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error during login", "data": err})
// 		return
// 	}

// 	if !CheckPasswordHash(pass, user.Password) {
// 		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid authorization", "data": nil})
// 		return
// 	}

// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["id"] = user.ID
// 	claims["username"] = user.Username
// 	claims["exp"] = time.Now().Add(time.Hour * 732).Unix()

// 	t, err := token.SignedString([]byte("SECRET"))
// 	if err != nil {
// 		c.SendStatus(fiber.StatusInternalServerError)
// 		return
// 	}
// 	c.JSON(fiber.Map{"status": "success", "message": "Success Login", "data": t})
// }

//CheckPasswordHash compare hash password
// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }
