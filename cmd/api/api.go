package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leetcode-golang-classroom/golang-ecom-sample/service/cart"
	"github.com/leetcode-golang-classroom/golang-ecom-sample/service/order"
	"github.com/leetcode-golang-classroom/golang-ecom-sample/service/product"
	"github.com/leetcode-golang-classroom/golang-ecom-sample/service/user"
)

type HttpApiServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *HttpApiServer {
	return &HttpApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *HttpApiServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	// setup user route
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)
	// setup product route
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subrouter)
	// setup cart route
	orderStore := order.NewStore(s.db)
	cartHandler := cart.NewHandler(s.db, orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(subrouter)
	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
