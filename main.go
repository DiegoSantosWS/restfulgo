package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	yaml "gopkg.in/yaml.v2"
	/**
	github.com/go-sql-driver/mysql not is used in apllication directamente
	*/
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Info from config file
type conf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	DB   string `yaml:"bd"`
}

//Db é um ponteiro do pacote sqlx
var Db *sqlx.DB

func main() {

	yamlFile, err := ioutil.ReadFile("db.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	conf := &conf{}
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

	tokenstring := generateToken(key) //recebendo a string do token
	//verifica se a string estar correta e retornar a chave.
	token, _ := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	//verifica se token é verdadeiro
	if token.Valid == true {
		err := conf.connection()
		if err != nil {
			fmt.Println("Erro ao abrir banco de dandos: ", err.Error())
			return
		}
		routers()
	}
}

//CONNECTION WITH DATABASE
func (conf *conf) connection() (err error) {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.User, conf.Pass, conf.Host, conf.Port, conf.DB)
	Db, err = sqlx.Open("mysql", uri)
	if err != nil {
		log.Fatal("ERRO ao conectar com banco de dados: ", err.Error())
		return
	}

	if err = Db.Ping(); err != nil {
		log.Fatalf("ERRO ao conectar com banco de dados: %s", err)
	}
	return
}

//Gerando o token JWT
func generateToken(key []byte) string {
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

//Definição das rotas usadas pelo sistema
func routers() {
	r := mux.NewRouter()

	r.HandleFunc("/v1/", MyHome).Methods("GET")
	r.HandleFunc("/v1/products", ApisV1Produtos).Methods("GET")
	r.HandleFunc("/v1/products/{id}", ApisV1Produtos).Methods("GET")
	r.HandleFunc("/v1/products/", ApisV1Produtos).Methods("POST")
	r.HandleFunc("/v1/products/{id}", ApisV1Produtos).Methods("PUT")
	r.HandleFunc("/v1/products/{id}", ApisV1Produtos).Methods("DELETE")

	err := http.ListenAndServe(":4000", r)
	if err != nil {
		log.Fatal("Erro ao instaciar o servidor: ", err.Error())
	}
}
