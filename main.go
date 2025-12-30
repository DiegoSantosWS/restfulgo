package main

import (
	"fmt"
	"log"
	"os"
	"restfulgo/connection"
	"restfulgo/pkgs/auth"

	route "restfulgo/routers"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {

	//buscando a chave gerada para token
	key, err := os.ReadFile("secret.str")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	key = []byte(key) //converte para byte

	tokenstring := auth.GenerateToken(key) //recebendo a string do token
	//verifica se a string estar correta e retornar a chave.
	token, _ := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	//verifica se token é verdadeiro
	if token.Valid == true {
		err := connection.GetConnection()
		if err != nil {
			fmt.Println("Erro ao abrir banco de dandos teste: ", err.Error())
			return
		}
		route.Routers()
	}
}
