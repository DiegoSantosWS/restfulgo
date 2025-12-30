package routers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	m "github.com/DiegoSantosWS/restfulgo/models"
	"github.com/gorilla/mux"
)

// Routers Definição das rotas usadas pelo sistema
func Routers() {
	r := mux.NewRouter()
	//routers of the clients
	r.HandleFunc("/v1/", m.MyHome).Methods("GET")
	r.HandleFunc("/v1/products", m.ControllerProdutos).Methods("GET")
	r.HandleFunc("/v1/products/{id}", m.ControllerProdutos).Methods("GET")
	r.HandleFunc("/v1/products/", m.ControllerProdutos).Methods("POST")
	r.HandleFunc("/v1/products/{id}", m.ControllerProdutos).Methods("PUT")
	r.HandleFunc("/v1/products/{id}", m.ControllerProdutos).Methods("DELETE")
	//routers of the clients
	r.HandleFunc("/v1/clients", m.ControllerClients).Methods("GET")
	r.HandleFunc("/v1/clients/{id}", m.ControllerClients).Methods("GET")
	r.HandleFunc("/v1/clients/", m.ControllerClients).Methods("POST")
	r.HandleFunc("/v1/clients/{id}", m.ControllerClients).Methods("PUT")
	r.HandleFunc("/v1/clients/{id}", m.ControllerClients).Methods("DELETE")
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Fatal("[ERRO RUNN] Erro ao instanciar o servidor")
	}
}
