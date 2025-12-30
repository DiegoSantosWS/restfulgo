package connection

import (
	"fmt"
	"log"
	"os"

	/**
	github.com/go-sql-driver/mysql not is used in apllication directamente
	*/
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Conf Info from config file
type Conf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	DB   string `yaml:"bd"`
}

// Db é um ponteiro do pacote sqlx
var Db *sqlx.DB

// GetConnection CONNECTION WITH DATABASE
func GetConnection() (err error) {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"), "3306", os.Getenv("MYSQL_DATABASE"))
	Db, err = sqlx.Open("mysql", uri)
	if err != nil {
		log.Fatal("ERRO ao conectar com banco de dados teste con: ", err.Error())
		return
	}

	if err = Db.Ping(); err != nil {
		log.Fatalf("ERRO ao conectar com banco de dados teste ping: %s", err)
	}
	return
}
