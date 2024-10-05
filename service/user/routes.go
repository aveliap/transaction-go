package user

import (
	"fmt"
	"net/http"

	"github.com/aveliap/transaction-go/config"
	"github.com/aveliap/transaction-go/service/auth"
	"github.com/aveliap/transaction-go/types"
	"github.com/aveliap/transaction-go/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct{
	Repo types.UserRepo
}

func NewHandler(repo types.UserRepo) *Handler  {
	return &Handler{Repo:repo}
}

func (handler *Handler) RegisterRoutes(router *mux.Router)  {
	router.HandleFunc("/login", handler.handleLogin).Methods("POST") 
	router.HandleFunc("/register", handler.handleRegister).Methods("POST") 
}

func (handler *Handler) handleLogin( writer http.ResponseWriter, req *http.Request){
	// get json payload
	var payload types.LoginUserPayload
	if err:= utils.ParseJSON(req,&payload)
	err!=nil{
		utils.WriteError(writer, http.StatusBadRequest, err)
		return
	}

	// validate
	if err := utils.Validate.Struct(payload)
	err!=nil{
		errors := err.(validator.ValidationErrors)
		utils.WriteError(writer, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return 
	}

	user, err := handler.Repo.GetUserByEmail(payload.Email)
	if err != nil{
		utils.WriteError(writer, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if !auth.ComparePassword(user.Password, []byte(payload.Password)) {
		utils.WriteError(writer, http.StatusBadRequest, fmt.Errorf("not found, invalid password"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, int(user.ID))
	if err != nil{
		utils.WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(writer, http.StatusOK, map[string]string{"token":token})
}

func (handler *Handler) handleRegister( writer http.ResponseWriter, req *http.Request){
	// get json payload
	var payload types.RegisterUserPayload
	if err:= utils.ParseJSON(req,&payload)
	err!=nil{
		utils.WriteError(writer, http.StatusBadRequest, err)
		return
	}

	// validate
	if err := utils.Validate.Struct(payload)
	err!=nil{
		errors := err.(validator.ValidationErrors)
		utils.WriteError(writer, http.StatusBadRequest, fmt.Errorf("invalid payoad %v", errors))
		return 
	}
	// check if user exists
	_,err := handler.Repo.GetUserByEmail(payload.Email)
	if err == nil{
		utils.WriteError(writer, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPass, err := auth.HashPass(payload.Password)
	if err != nil {
		utils.WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	err = handler.Repo.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Email: payload.Email,
		Password: hashedPass,
	})

	if err != nil {
		utils.WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(writer, http.StatusCreated,nil)
}