package products

import (
	"log"
	"net/http"

	"github.com/MostafaSensei106/E-Commerce/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	err := h.service.GatAll(r.Context())
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, []string{})
}
