package transport

import "github.com/gorilla/mux"

// NewRouter creates a router
func NewRouter(pt ProductTransport) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/products", pt.GetProducts).Methods("GET")
	router.HandleFunc("/products", pt.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id:[0-9]+}", pt.GetProduct).Methods("GET")
	router.HandleFunc("/products/{id:[0-9]+}", pt.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id:[0-9]+}", pt.DeleteProduct).Methods("DELETE")
	return router
}
