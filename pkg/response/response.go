package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message any
	Status  int
}

func Send(w http.ResponseWriter, status int, data any) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
