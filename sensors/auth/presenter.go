package auth

//SignedUser user representation
type SignedUser struct {
	UserID string `json:"userId"`
	Token  string `json:"token"`
}

//RegistrationResponse builds response for registration request
func RegistrationResponse(user *User) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": "User registered",
		"user": map[string]interface{}{
			"id":    user.ID.String(),
			"email": user.Email,
		},
	}
}