package Authenticator

import (
	"e_market/src/Config"
	"e_market/src/Models"
	"time"

	"log"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(Config.JWT_secret_key)

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(user Models.User) (string, bool) {
	expirationTime := time.Now().Add(time.Minute * 30)

	claims := &JWTClaim{
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // MAYBE US SHA256 instead
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Error while generating JWT: ", err)
		return "", false
	}
	return tokenString, true
}

// If token is invalid, return error, otherwise return nothing
func ValidateJWT(signedToken string) (claims *JWTClaim, isValid bool) {
	parsedClaims := JWTClaim{}
	token, err := jwt.ParseWithClaims(signedToken, &parsedClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		log.Println("Unable to validate token: ", err)
		return claims, false
	}

	if claims, ok := token.Claims.(*JWTClaim); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return claims, false
		}
		return claims, true
	}
	return claims, false
}