package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aveliap/transaction-go/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct{
	addr string
	db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer{
	return &APIServer{
		addr:addr,
		db: db,
	}
}



func (server *APIServer) Run() error  {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userService :=user.NewHandler()
	userService.RegisterRoutes(subrouter)
	log.Println("Listening on", server.addr)
	return http.ListenAndServe(server.addr, router)
}