package json

import (
	"encoding/json"
	"net/http"

	"github.com/MostafaSensei106/E-Commerce/internal/constants"
)

func Write(w http.ResponseWriter, status int, data any) {
	w.Header().Set(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}
