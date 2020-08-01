package auth

import "github.com/twinj/uuid"

//LoginInput login payload
type LoginInput struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
