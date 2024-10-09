package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/frhnfrnk/go-ecommerce/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandler(t *testing.T){
	userStore := &mockUserStore{}
	handler := Newhandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T){
		payload := types.RegisterUserPayload{
			FirstName: "John",
			LastName: "Doe",
			Email: "invalid",
			Password: "password",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should correctly register the user", func(t *testing.T){
		payload := types.RegisterUserPayload{
			FirstName: "John",
			LastName: "Doe",
			Email: "valid@gmail.com",
			Password: "password",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})


}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error){
	return  &types.User{}, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error){
	return &types.User{}, nil
}

func (m *mockUserStore) CreateUser(u types.User) error{
	return nil
}
