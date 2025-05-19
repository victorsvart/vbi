package response

import (
	"encoding/json"
	"net/http"
)

func Send(w http.ResponseWriter, status int, data any) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func SendOk(w http.ResponseWriter, data any) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func SendCreated(w http.ResponseWriter, data any) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func SendInternalServerError(w http.ResponseWriter, data any) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(data)
}

func SendBadRequest(w http.ResponseWriter, data any) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(data)
}
