package auth

import (
	"fmt"
	"log"
	"sensors"
	"sensors/data"

	"golang.org/x/crypto/bcrypt"
)

//Interactor implements interactions
type Interactor interface {
	AuthenticateByEmail(email, password string) (SignedUser, error)
	Register(*User) (User, error)
}

type interactor struct {
	config      sensors.Auth
	repository  Repository
	userChannel data.ChanNewUsers
}

//NewInteractor interactor
func NewInteractor(config sensors.Auth, authRepo Repository, usersChan data.ChanNewUsers) Interactor {
	return &interactor{
		config:      config,
		repository:  authRepo,
		userChannel: usersChan,
	}
}

func (i interactor) AuthenticateByEmail(email, password string) (SignedUser, error) {
	user, err := i.repository.GetByEmail(email)
	if err != nil {
		return SignedUser{}, err
	}

	token, err := i.authenticateUser(&user, password)
	if err != nil {
		return SignedUser{}, err
	}

	return SignedUser{UserID: user.ID.String(), Token: token}, nil
}

func (i interactor) authenticateUser(user *User, password string) (string, error) {
	if isValidPassword := i.comparePasswords(user.Password, []byte(password)); !isValidPassword {
		msg := fmt.Sprintf("user or password invalid")
		return "", ErrUnauthorized{Message: msg}
	}

	tok := generateToken(user)
	token, err := getTokenString(i.config.JWTSecret, tok)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (i interactor) comparePasswords(hash string, password []byte) bool {
	byteHash := []byte(hash)

	err := bcrypt.CompareHashAndPassword(byteHash, password)
	if err != nil {
		log.Printf("err comparing passwords %v", err)
		return false
	}
	return true
}

func (i interactor) Register(user *User) (User, error) {
	passHash, err := i.hashUserPassword([]byte(user.Password))
	if err != nil {
		return User{}, err
	}

	user.Password = passHash
	u, err := i.repository.Add(*user)
	if err != nil {
		return User{}, err
	}

	i.postNewUserToChannel(&u)
	return u, nil
}

func (i interactor) hashUserPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, 8)
	if err != nil {
		hashErr := ErrHashPassword{password: string(password), message: err.Error()}
		return "", &hashErr
	}

	return string(hash), nil
}

func (i interactor) postNewUserToChannel(user *User) {
	u := parseToNewUser(*user)
	go func() { i.userChannel.Writer <- u }()
}
