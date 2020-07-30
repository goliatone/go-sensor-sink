package auth

import (
	"fmt"
	"sensors/storage"
)

//SearchParameters handles querying AuthUser
type SearchParameters struct {
}

//Repository is the interface to handle AuthUser persistence
type Repository interface {
	Add(AuthUser) (AuthUser, error)
	Delete(AuthUser) error
	// Get(SearchParameters) ([]AuthUser, error)
	GetByEmail(string) (AuthUser, error)
	GetByUsername(string) (AuthUser, error)
}

type repository struct {
	database *storage.Database
}

//NewRepository creates a new AuthUser repository instance
func NewRepository(db *storage.Database) Repository {
	return &repository{database: db}
}

//Add a new user record to the DB

func (r repository) Add(user AuthUser) (AuthUser, error) {
	result := r.database.Create(&user)
	if err := result.Error; err != nil {
		return AuthUser{}, NewErrUnexpected(err)
	}
	return user, nil
}

func (r repository) Delete(user AuthUser) error {
	result := r.database.Delete(&user)
	if err := result.Error; err != nil {
		return NewErrUnexpected(result.Error)
	}
	return nil
}

//GetByEmail will fetch a user by email
func (r repository) GetByEmail(email string) (AuthUser, error) {
	var user AuthUser

	result := r.database.Where("email = ?", email).First(&user)
	if result.RecordNotFound() {
		msg := fmt.Sprintf("AuthUser with email \"%s\" not found", email)
		return AuthUser{}, NewErrUserNotFound(msg)
	}

	if err := result.Error; err != nil {
		return AuthUser{}, err
	}

	return user, nil
}

//GetByUsername will fetch a user by username
func (r repository) GetByUsername(username string) (AuthUser, error) {
	var user AuthUser

	result := r.database.Where("username = ?", username).First(&user)
	if result.RecordNotFound() {
		msg := fmt.Sprintf("AuthUser with username \"%s\" not found", username)
		return AuthUser{}, NewErrUserNotFound(msg)
	}

	if err := result.Error; err != nil {
		return AuthUser{}, err
	}

	return user, nil
}
