package products

import (
	"log"
	"net/http"
	"strconv"

	repo "github.com/MostafaSensei106/E-Commerce/internal/adapters/postgresql/sqlc"
	"github.com/MostafaSensei106/E-Commerce/internal/json"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAllProducts(r.Context())
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, products)
}

func (h *handler) GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProductByID(r.Context(), id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	json.Write(w, http.StatusOK, product)
}

func (h *handler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var body repo.Product
	if err := json.Read(r, &body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := h.service.CreateNewProduct(r.Context(), repo.CreateProductParams{
		Name:         body.Name,
		PriceInCents: body.PriceInCents,
		Quantity:     body.Quantity,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, product)

}

func (h *handler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}
	var body repo.Product
	if err := json.Read(r, &body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := h.service.UpdateProduct(r.Context(), repo.UpdateProductWhereIDParams{
		ID:           id,
		Name:         body.Name,
		PriceInCents: body.PriceInCents,
		Quantity:     body.Quantity,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, product)

}

func (h *handler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteProduct(r.Context(), id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
