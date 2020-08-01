package auth

import (
	"fmt"
	"sensors/storage"
)

//SearchParameters handles querying User
type SearchParameters struct {
}

//Repository is the interface to handle User persistence
type Repository interface {
	Add(User) (User, error)
	Delete(User) error
	// Get(SearchParameters) ([]User, error)
	GetByEmail(string) (User, error)
	GetByUsername(string) (User, error)
}

type repository struct {
	database *storage.Database
}

//NewRepository creates a new User repository instance
func NewRepository(db *storage.Database) Repository {
	return &repository{database: db}
}

//Add a new user record to the DB

func (r repository) Add(user User) (User, error) {
	result := r.database.Create(&user)
	if err := result.Error; err != nil {
		return User{}, NewErrUnexpected(err)
	}
	return user, nil
}

func (r repository) Delete(user User) error {
	result := r.database.Delete(&user)
	if err := result.Error; err != nil {
		return NewErrUnexpected(result.Error)
	}
	return nil
}

//GetByEmail will fetch a user by email
func (r repository) GetByEmail(email string) (User, error) {
	var user User

	result := r.database.Where("email = ?", email).First(&user)
	if result.RecordNotFound() {
		msg := fmt.Sprintf("User with email \"%s\" not found", email)
		return User{}, NewErrUserNotFound(msg)
	}

	if err := result.Error; err != nil {
		return User{}, err
	}

	return user, nil
}

//GetByUsername will fetch a user by username
func (r repository) GetByUsername(username string) (User, error) {
	var user User

	result := r.database.Where("username = ?", username).First(&user)
	if result.RecordNotFound() {
		msg := fmt.Sprintf("User with username \"%s\" not found", username)
		return User{}, NewErrUserNotFound(msg)
	}

	if err := result.Error; err != nil {
		return User{}, err
	}

	return user, nil
}
