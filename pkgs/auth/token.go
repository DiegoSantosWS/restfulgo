package auth

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateToken Gerando o token JWT
func GenerateToken(key []byte) string {
	secretKey := []byte(key)
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		Issuer:    "APId2js898ilsje6272g726g072gso",
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //Incorporar informações do usuário ao token
	tokenstring, err := token.SignedString(secretKey)          // token -> string. Only server knows this secret (wsitesb).
	if err != nil {
		log.Fatalln(err)
	}

	return tokenstring
}
