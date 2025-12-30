package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	conect "github.com/DiegoSantosWS/restfulgo/connection"
	"github.com/gorilla/mux"
)

// MyHome teste teste teste teste
func MyHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "METODO USADO <%s>:\r\n", r.Method)
	fmt.Fprintf(w, "URI DO USADA <%s>: \r\n", r.URL.Path)
	fmt.Fprintf(w, "MINHA HOME")
}

// ControllerProdutos controller de chamadas
func ControllerProdutos(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)
	id, _ := strconv.Atoi(code["id"])

	switch r.Method {
	case "POST":
		CreateProduct(w, r)
	case "DELETE":
		DeleteProduct(id, w, r)
	case "PUT":
		UpdateProduct(id, w, r)
	case "GET":
		if id != 0 {
			GetlistProductByID(id, w, r)
		} else {
			GetListProduct(w, r)
		}
	}
}

// GetListProduct list all products resgistered in your databases
func GetListProduct(w http.ResponseWriter, r *http.Request) {
	p := Product{}
	sql := "SELECT * FROM products "
	res, err := conect.Db.Queryx(sql)
	if err != nil {
		log.Fatal("ERROR: listar produtos: ", err.Error())
	}
	defer res.Close()

	var prod []Product
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

// GetlistProductByID list an product by your code
func GetlistProductByID(id int, w http.ResponseWriter, r *http.Request) {
	p := Product{}
	sql := "SELECT * FROM products WHERE id = ?"
	res, err := conect.Db.Queryx(sql, id)
	if err != nil {
		log.Fatal("ERROR: listar produtos: ", err.Error())
		return
	}
	defer res.Close()

	var prod []Product
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

// DeleteProduct exclude a product
func DeleteProduct(id int, w http.ResponseWriter, r *http.Request) {
	sql := "DELETE FROM products WHERE id = ? "
	rows, err := conect.Db.Exec(sql, id)
	if err != nil {
		log.Fatal("ERRO: erro ao deletar produto: ", err.Error())
	}

	linhas, err := rows.RowsAffected()
	if err != nil {
		log.Fatal("ERRO: erro ao deletar produto inexistente: ", err.Error())
	}

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

// UpdateProduct altereding a product
func UpdateProduct(id int, w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	price := r.FormValue("price")

	sql := "UPDATE products SET name = ?, price = ? WHERE id = ? "
	rows, err := conect.Db.Exec(sql, name, price, id)
	if err != nil {
		log.Fatal("ERRO: erro ao alterar produto: ", err.Error())
	}

	linhas, err := rows.RowsAffected()
	if err != nil {
		log.Fatal("ERRO: erro ao alterar produto inexistente: ", err.Error())
	}

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

// CreateProduct create new product in databases
func CreateProduct(w http.ResponseWriter, r *http.Request) {

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
	rows, err := conect.Db.Exec(sql, Name, Description, Stock, Width, Height, Amount, Weight, Price, Discount, Promotion)
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

// ControllerClients controller de chamadas
func ControllerClients(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)
	id, _ := strconv.Atoi(code["id"])

	switch r.Method {
	case "POST":
		CreateClient(w, r)
	case "DELETE":
		DeleteClient(id, w, r)
	case "PUT":
		UpdateClient(id, w, r)
	case "GET":
		if id != 0 {
			GetClientByID(id, w, r)
		} else {
			GetListClients(w, r)
		}
	}
}

// GetListClients list all clients resgistered in your databases
func GetListClients(w http.ResponseWriter, r *http.Request) {
	sql := "SELECT * FROM clients "
	res, err := conect.Db.Queryx(sql)
	if err != nil {
		log.Fatal("ERROR: listar clients: ", err.Error())
	}
	defer res.Close()
	c := Clients{}
	var cli []Clients
	for res.Next() {
		err := res.Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.Status, &c.Date)
		if err != nil {
			log.Fatal("ERROR: scan clients: ", err.Error())

		}
		cli = append(cli, c)
	}

	cliJSON, err := json.Marshal(cli)
	if err != nil {
		log.Fatal("ERROR: json clients", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(cliJSON)
}

// GetClientByID get a client by your code
func GetClientByID(id int, w http.ResponseWriter, r *http.Request) {

	c := Clients{}
	sql := "SELECT c.*, cEnd.id, cEnd.idClients, cEnd.address, cEnd.number, cEnd.city, cEnd.neighborhood, cEnd.country, cEnd.state FROM clients as c LEFT JOIN clients_address AS cEnd ON c.id = cEnd.idClients WHERE c.id = ?"
	res, err := conect.Db.Queryx(sql, id)
	if err != nil {
		log.Fatal("ERROR: scan clients: ", err.Error())
		return
	}
	defer res.Close()

	for res.Next() {
		end := AddressClient{}

		err := res.Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.Status, &c.Date, &end.ID, &end.Address, &end.City, &end.Country, &end.IDclients, &end.Neighborhood, &end.Number, &end.State)
		if err != nil {
			log.Fatal("ERROR: scan clients: ", err.Error())

		}
		c.Address = append(c.Address, end)
	}

	cliJSON, err := json.Marshal(c)
	if err != nil {
		log.Fatal("ERROR: json clients by id", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(cliJSON)
}

// DeleteClient delete a client
func DeleteClient(id int, w http.ResponseWriter, r *http.Request) {
	sql := "DELETE FROM clients WHERE id = ? "
	rows, err := conect.Db.Exec(sql, id)
	if err != nil {
		log.Fatal("ERRO: erro ao deletar clients: ", err.Error())
	}

	linhas, err := rows.RowsAffected()
	if err != nil {
		log.Fatal("ERRO: erro ao deletar clients inexistente: ", err.Error())
	}
	fmt.Println(linhas)
	msg := Message{}
	if linhas != 0 {
		msg = Message{true, "Client deleted successfully.", id}
	} else {
		msg = Message{false, "Client not deleted or not localized.", id}
	}
	cliDelJSON, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("ERROR: json produtos", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(cliDelJSON)
}

// UpdateClient client update
func UpdateClient(id int, w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	status := r.FormValue("status")

	sql := "UPDATE clients SET name = ?, email = ?, phone = ?, status = ? WHERE id = ? "
	rows, err := conect.Db.Exec(sql, name, email, phone, status, id)
	if err != nil {
		log.Fatal("ERRO: erro ao alterar client: ", err.Error())
	}

	linhas, err := rows.RowsAffected()
	if err != nil {
		log.Fatal("ERRO: erro ao alterar client inexistente: ", err.Error())
	}
	msg := Message{}
	if linhas != 0 {
		msg = Message{true, "Client altered successfully.", id}
	} else {
		msg = Message{false, "Client not altered or not localized.", id}
	}
	cliUpJSON, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("ERROR: json produtos", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(cliUpJSON)
}

// CreateClient create new client in databases
func CreateClient(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	status := r.FormValue("status")

	sql := "INSERT clients SET "
	sql += "name = ?,"
	sql += "email = ?,"
	sql += "phone = ?,"
	sql += "status = ?,"
	sql += "date = ? "
	rows, err := conect.Db.Exec(sql, name, email, phone, status, time.Now().Format("2006-01-02"))
	if err != nil {
		log.Fatal("Erro ao inserir um novo client: ", err.Error())
		return
	}

	linhas, err := rows.RowsAffected()
	msg := Message{}
	lastID, _ := rows.LastInsertId()
	if linhas != 0 {
		msg = Message{true, "Client created successfully.", int(lastID)}
	} else {
		msg = Message{false, "Client not created.", int(lastID)}
	}
	cliInsertJSON, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("ERROR: json produtos", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(cliInsertJSON)
}
