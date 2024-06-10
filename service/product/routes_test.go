package product

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/leetcode-golang-classroom/golang-ecom-sample/types"
)

func TestProductServiceHandler(t *testing.T) {
	// mock store
	productStore := &mockProductStore{}
	userStore := &mockUserStore{}
	handler := NewHandler(productStore, userStore)
	t.Run("should failed the product payload is invalid", func(t *testing.T) {
		payload := types.CreateProductPayload{
			Name:        "",
			Description: "tset",
			Quantity:    1,
			Price:       2,
			Image:       "",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/products", handler.handleCreateProduct)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
	t.Run("Should correctly create product", func(t *testing.T) {
		payload := types.CreateProductPayload{
			Name:        "product",
			Description: "test product",
			Quantity:    1,
			Price:       2,
			Image:       "prodcut_image",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/products", handler.handleCreateProduct)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockProductStore struct{}

func (m *mockProductStore) GetProducts() ([]types.Product, error) {
	return nil, nil
}

func (m *mockProductStore) CreateProduct(_product types.Product) error {
	return nil
}

func (m *mockProductStore) GetProductsByIDs(_ps []int) ([]types.Product, error) {
	return nil, nil
}

func (m *mockProductStore) UpdateProduct(_tx *sql.Tx, product types.Product) error {
	return nil
}

type mockUserStore struct{}

func (u *mockUserStore) GetUserByEmail(_email string) (*types.User, error) {
	return nil, nil
}

func (u *mockUserStore) GetUserByID(_id int) (*types.User, error) {
	return nil, nil
}

func (u *mockUserStore) CreateUser(_user types.User) error {
	return nil
}
