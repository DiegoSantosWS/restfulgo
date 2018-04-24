package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	ID          int     `json:"id" db:"id"`
	Nome        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Stock       float64 `json:"stock" db:"stock"`
	Width       float64 `json:"width" db:"width"`
	Height      float64 `json:"heigth" db:"height"`
	Amount      float64 `json:"amount" db:"amount"`
	Weight      float64 `json:"weight" db:"weight"`
	Price       float64 `json:"price" db:"price"`
	Discount    float64 `json:"discount" db:"discount"`
	Promotion   float64 `json:"promotion" db:"promotion"`
}

type Message struct {
	Executed  bool
	Message   string
	IDProduct int
}

//ApisV1Produtos teste teste teste teste
func ApisV1Produtos(w http.ResponseWriter, r *http.Request) {
	//uri := r.RequestURI
	switch r.Method {
	case "GET":
		Produtos(w, r)
	case "POST":
		Produtos(w, r)
	case "PUT":
		Produtos(w, r)
	case "DELETE":
		Produtos(w, r)
	}
}

//MyHome teste teste teste teste
func MyHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "METODO USADO <%s>:\r\n", r.Method)
	fmt.Fprintf(w, "URI DO USADA <%s>: \r\n", r.URL.Path)
	fmt.Fprintf(w, "MINHA HOME")
}

//Produtos teste teste teste teste
func Produtos(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)
	id, _ := strconv.Atoi(code["id"])

	switch r.Method {
	case "POST":
		insertProducts(w, r)
		break
	case "DELETE":
		deleteProducts(id, w, r)
		break
	case "PUT":
		updateProducts(id, w, r)
		break
	case "GET":
		if id != 0 {
			listProductsByID(id, w, r)
		} else {
			listProducts(w, r)
		}
		break
	}
}

func listProducts(w http.ResponseWriter, r *http.Request) {
	p := Products{}
	sql := "SELECT * FROM products "
	res, err := Db.Queryx(sql)
	if err != nil {
		log.Fatal("ERROR: listar produtos: ", err.Error())
	}
	defer res.Close()
	var prod []Products
	for res.Next() {
		err := res.StructScan(&p)
		if err != nil {
			log.Fatal("ERROR: scan produtos: ", err.Error())

		}
		prod = append(prod, p)
	}
	prodJSON, err := json.Marshal(prod)
	if err != nil {
		log.Fatal("ERROR: json produtos", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(prodJSON)
}

func listProductsByID(id int, w http.ResponseWriter, r *http.Request) {
	p := Products{}
	sql := "SELECT * FROM products WHERE id = ?"
	res, err := Db.Queryx(sql, id)
	if err != nil {
		log.Fatal("ERROR: listar produtos: ", err.Error())
		return
	}
	defer res.Close()
	var prod []Products
	for res.Next() {
		err := res.StructScan(&p)
		if err != nil {
			log.Fatal("ERROR: scan produtos: ", err.Error())
			return

		}
		prod = append(prod, p)
	}
	prodJSON, err := json.Marshal(prod)
	if err != nil {
		log.Fatal("ERROR: json produtos", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(prodJSON)
}

func deleteProducts(id int, w http.ResponseWriter, r *http.Request) {
	sql := "DELETE FROM products WHERE id = ? "
	rows, err := Db.Exec(sql, id)
	if err != nil {
		log.Fatal("ERRO: erro ao deletar produto: ", err.Error())
	}

	linhas, err := rows.RowsAffected()
	if err != nil {
		log.Fatal("ERRO: erro ao deletar produto inexistente: ", err.Error())
	}
	fmt.Println(linhas)
	d := Message{}
	if linhas != 0 {
		d = Message{true, "Product deleted successfully.", id}
	} else {
		d = Message{false, "Product not deleted or not localized.", id}
	}
	prodDelJSON, err := json.Marshal(d)
	if err != nil {
		log.Fatal("ERROR: json produtos", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(prodDelJSON)

}

func updateProducts(id int, w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	price := r.FormValue("price")

	sql := "UPDATE products SET name = ?, price = ? WHERE id = ? "
	rows, err := Db.Exec(sql, name, price, id)
	if err != nil {
		log.Fatal("ERRO: erro ao alterar produto: ", err.Error())
	}

	linhas, err := rows.RowsAffected()
	if err != nil {
		log.Fatal("ERRO: erro ao alterar produto inexistente: ", err.Error())
	}
	fmt.Println(linhas)
	d := Message{}
	if linhas != 0 {
		d = Message{true, "Product altered successfully.", id}
	} else {
		d = Message{false, "Product not altered or not localized.", id}
	}
	prodUpJSON, err := json.Marshal(d)
	if err != nil {
		log.Fatal("ERROR: json produtos", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(prodUpJSON)
}

func insertProducts(w http.ResponseWriter, r *http.Request) {

	Name := r.FormValue("name")
	Description := r.FormValue("description")
	Stock := r.FormValue("stock")
	Width := r.FormValue("width")
	Height := r.FormValue("heigth")
	fmt.Println(Height)
	Amount := r.FormValue("amount")
	Weight := r.FormValue("weight")
	Price := r.FormValue("price")
	Discount := r.FormValue("discount")
	Promotion := r.FormValue("promotion")

	sql := "INSERT products SET "
	sql += "name = ?,"
	sql += "description = ?,"
	sql += "stock = ?,"
	sql += "width = ?,"
	sql += "height = ?,"
	sql += "amount = ?,"
	sql += "weight = ?,"
	sql += "price = ?,"
	sql += "discount = ?,"
	sql += "promotion = ? "
	rows, err := Db.Exec(sql, Name, Description, Stock, Width, Height, Amount, Weight, Price, Discount, Promotion)
	if err != nil {
		log.Fatal("Erro ao inserir um novo produtos: ", err.Error())
		return
	}

	linhas, err := rows.RowsAffected()
	d := Message{}
	lastID, _ := rows.LastInsertId()
	if linhas != 0 {
		d = Message{true, "Product inserted successfully.", int(lastID)}
	} else {
		d = Message{false, "Product not inserted.", int(lastID)}
	}
	prodInsertJSON, err := json.Marshal(d)
	if err != nil {
		log.Fatal("ERROR: json produtos", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(prodInsertJSON)

}
