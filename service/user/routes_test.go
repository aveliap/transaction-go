package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aveliap/transaction-go/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(test *testing.T)  {
	userRepo := &mockUserRepo{}

	handler:= NewHandler(userRepo)

	test.Run("should fail if the user payload is invalid", func (test *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName: "username",
			Email: "test",
			Password: "123",
		}

		marshalled,_ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))		

		if err != nil {
			test.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr,req)

		if rr.Code != http.StatusBadRequest{
			test.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	test.Run("should correctly registered the user", func (test *testing.T)  {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName: "username",
			Email: "test@mail.com",
			Password: "123",
		}

		marshalled,_ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))		

		if err != nil {
			test.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr,req)

		if rr.Code != http.StatusCreated{
			test.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockUserRepo struct{}

func (mock *mockUserRepo) GetUserByEmail(email string) (*types.User, error)  {
	return nil, fmt.Errorf("user not found")
}

func (mock *mockUserRepo) GetUserByID(id uint) (*types.User, error)  {
	return nil,nil
}

func (mock *mockUserRepo) CreateUser(types.User) error{
	return nil
}
	