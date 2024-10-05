package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON (req *http.Request, payload any) error{
	if req.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(req.Body).Decode(payload)
}

func WriteJSON(write http.ResponseWriter, status int, val any) error {
	write.Header().Add("Content-Type","application/json")
	write.WriteHeader(status)

	return json.NewEncoder(write).Encode(val)
}

func WriteError(writer http.ResponseWriter, status int, err error){
	WriteJSON(writer,status, map[string]string{"error":  err.Error()})
}
