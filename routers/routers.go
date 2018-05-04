package routers

import (
	"log"
	"net/http"

	m "github.com/DiegoSantosWS/restfulgo/models"
	"github.com/gorilla/mux"
)

//Routers Definição das rotas usadas pelo sistema
func Routers() {
	r := mux.NewRouter()
	//routers of the clients
	r.HandleFunc("/v1/", m.MyHome).Methods("GET")
	r.HandleFunc("/v1/products", m.CtrProdutos).Methods("GET")
	r.HandleFunc("/v1/products/{id}", m.CtrProdutos).Methods("GET")
	r.HandleFunc("/v1/products/", m.CtrProdutos).Methods("POST")
	r.HandleFunc("/v1/products/{id}", m.CtrProdutos).Methods("PUT")
	r.HandleFunc("/v1/products/{id}", m.CtrProdutos).Methods("DELETE")
	//routers of the clients
	r.HandleFunc("/v1/clients", m.CtrClients).Methods("GET")
	r.HandleFunc("/v1/clients/{id}", m.CtrClients).Methods("GET")
	r.HandleFunc("/v1/clients/", m.CtrClients).Methods("POST")
	r.HandleFunc("/v1/clients/{id}", m.CtrClients).Methods("PUT")
	r.HandleFunc("/v1/clients/{id}", m.CtrClients).Methods("DELETE")

	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatal("Erro ao instaciar o servidor: ", err.Error())
	}

}
