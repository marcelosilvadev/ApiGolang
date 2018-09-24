package api

import (
	"database/sql"
	"erncliente/api/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// App struct ...
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//StartServer ...
func (a *App) StartServer() {
	a.Router = mux.NewRouter()
	s := a.Router.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/health", handler.HealthCheck).Methods(http.MethodGet)
	s.HandleFunc("/clientes", handler.InsertClient).Methods(http.MethodPost)
	s.HandleFunc("/clientes/{id:[0-9]+}", handler.UpdateClient).Methods(http.MethodPut)
	s.HandleFunc("/clientes/{id:[0-9]+}", handler.DeleteClient).Methods(http.MethodDelete)
	s.HandleFunc("/clientes/{id:[0-9]+}", handler.GetClient).Methods(http.MethodGet)
	s.HandleFunc("/clientes", handler.GetClients).Methods(http.MethodGet)
	a.Router.Handle("/api/v1/{_:.*}", a.Router)
	port := 10001
	log.Printf("Starting Server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router))
}
