package jwt

import "github.com/golang-jwt/jwt"

const JWT_SCRET = "123"

func CreateToken(id int, username string) (string, error) {
	claims := jwt.MapClaims{}

	claims["id"] = id
	claims["username"] = username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWT_SCRET))
}