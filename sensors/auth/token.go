package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//TokenDetails details
type TokenDetails struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
}

//TokenClaims claims
type TokenClaims struct {
	User TokenDetails `json:"user"`

	jwt.StandardClaims
}

func generateToken(user *User) *jwt.Token {
	expirationTime := time.Now().Add(72 * time.Hour).Unix()
	issuedAt := time.Now().Unix()
	claims := TokenClaims{
		User: TokenDetails{
			UserID: user.ID.String(),
			Email:  user.Email,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			IssuedAt:  issuedAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token
}

func getTokenString(secret string, token *jwt.Token) (string, error) {
	str, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", ErrTokenParsing{message: err.Error()}
	}
	return str, nil
}

//ParseToken parse JWT token
func ParseToken(token, secret string, claims *TokenClaims) (*jwt.Token, error) {
	tok, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	return tok, err
}

//ValidateToken valid JWT token
func ValidateToken(tok *jwt.Token) bool {
	if !tok.Valid {
		return false
	}
	return true
}
