package main

import (
	"fmt"
	"io/ioutil"
	"log"

	conect "github.com/DiegoSantosWS/restfulgo/connection"
	funcs "github.com/DiegoSantosWS/restfulgo/pakgs"
	route "github.com/DiegoSantosWS/restfulgo/routers"
	jwt "github.com/dgrijalva/jwt-go"
	yaml "gopkg.in/yaml.v2"
)

func main() {

	yamlFile, err := ioutil.ReadFile("db.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	conf := &conect.Conf{}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Printf("yamlFile.Get unmarshal err   #%v ", err)
	}
	//buscando a chave gerada para token
	key, err := ioutil.ReadFile("secret.str")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	key = []byte(key) //converte para byte

	tokenstring := funcs.GenerateToken(key) //recebendo a string do token
	//verifica se a string estar correta e retornar a chave.
	token, _ := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	//verifica se token é verdadeiro
	if token.Valid == true {
		err := conect.Connection(conf)
		if err != nil {
			fmt.Println("Erro ao abrir banco de dandos: ", err.Error())
			return
		}
		route.Routers()
	}
}
