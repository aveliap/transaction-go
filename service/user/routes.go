package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{

}

func NewHandler() *Handler  {
	return &Handler{}
}

func (handler *Handler) RegisterRoutes(router *mux.Router)  {
	router.HandleFunc("/login", handler.handleLogin).Methods("POST") 
	router.HandleFunc("/register", handler.handleRegister).Methods("POST") 
}

func (handler *Handler) handleLogin( writer http.ResponseWriter, req *http.Request){

}

func (handler *Handler) handleRegister( writer http.ResponseWriter, req *http.Request){

}