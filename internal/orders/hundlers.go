package orders

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

func (h *handler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.service.GetAllOrders(r.Context())
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, orders)
}

func (h *handler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "invalid order ID", http.StatusBadRequest)
		return
	}

	order, err := h.service.GetOrderByID(r.Context(), id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}

	json.Write(w, http.StatusOK, order)

}

func (h *handler) PlaceNewOrderHandler(w http.ResponseWriter, r *http.Request) {
	var body repo.CreateOrderParams
	if err := json.Read(r, &body); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := body.Validate(); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := h.service.PlaceNewOrder(r.Context(), body)

	if err != nil {
		log.Println(err.Error())

		if err == ErrProductNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if err == ErrProductHasNoStock {
			http.Error(w, err.Error(), http.StatusNoContent)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, order)
}
