package connection

import (
	"fmt"
	"log"
	/**
	github.com/go-sql-driver/mysql not is used in apllication directamente
	*/
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//Conf Info from config file
type Conf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	DB   string `yaml:"bd"`
}

//Db é um ponteiro do pacote sqlx
var Db *sqlx.DB

//Connection CONNECTION WITH DATABASE
func Connection(conf *Conf) (err error) {
	fmt.Println(conf)
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
