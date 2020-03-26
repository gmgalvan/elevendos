package transport

import (
	"context"
	"encoding/json"
	"lab/productLab/internal/entity"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ProductUC usecase interface
type ProductUC interface {
	Create(ctx context.Context, product *entity.Product) (*entity.Product, error)
	ByID(ctx context.Context, id int) (*entity.Product, error)
	Update(ctx context.Context, id int, product *entity.Product) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, start, count int) ([]*entity.Product, error)
}

// ProductTransport struct for transport
type ProductTransport struct {
	Product ProductUC
}

// NewProductTransport creates new transport product
func NewProductTransport(puc ProductUC) ProductTransport {
	return ProductTransport{
		Product: puc,
	}
}

// GetProduct get a product by ID
func (pt ProductTransport) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
	}

	ctx := context.Background()
	p, err := pt.Product.ByID(ctx, id)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJSON(w, http.StatusOK, p)
}

// GetProducts get multiple products
func (pt ProductTransport) GetProducts(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}
	ctx := context.Background()
	products, err := pt.Product.List(ctx, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

// CreateProduct creates a new product
func (pt ProductTransport) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p entity.Product
	// TODO get dates
	// ..
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	ctx := context.Background()
	rp, err := pt.Product.Create(ctx, &p)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusCreated, rp)
}

// UpdateProduct updates new product
func (pt ProductTransport) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	var p entity.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	p.ID = id
	ctx := context.Background()
	err = pt.Product.Update(ctx, id, &p)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJSON(w, http.StatusOK, p)
}

// DeleteProduct delete a product by id
func (pt ProductTransport) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}
	ctx := context.Background()
	if err := pt.Product.Delete(ctx, id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
