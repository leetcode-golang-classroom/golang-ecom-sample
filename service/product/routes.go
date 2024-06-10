package product

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/leetcode-golang-classroom/golang-ecom-sample/service/auth"
	"github.com/leetcode-golang-classroom/golang-ecom-sample/types"
	"github.com/leetcode-golang-classroom/golang-ecom-sample/utils"
)

type Handler struct {
	store     types.ProductStore
	userStore types.UserStore
}

func NewHandler(store types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", auth.WithJWTAuth(h.handleGetProducts, h.userStore)).Methods(http.MethodGet)
	router.HandleFunc("/products", auth.WithJWTAuth(h.handleCreateProduct, h.userStore)).Methods(http.MethodPost)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.FailOnError(utils.WriteJSON(w, http.StatusOK, ps), "utilsWriteJSON err:")
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	// json payload
	var payload types.CreateProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		var valErrs validator.ValidationErrors
		if errors.As(err, &valErrs) {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", valErrs))
		}
		return
	}
	// write into db
	err := h.store.CreateProduct(types.Product{
		Name:        payload.Name,
		Image:       payload.Image,
		Description: payload.Description,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.FailOnError(utils.WriteJSON(w, http.StatusCreated, nil), "utils.WriteJSON err:")
}
