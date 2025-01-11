package handlers

import (
	"encoding/json"
	"final_project/store"
	"net/http"
)

type Handler struct {
	Store *store.Stores
}

func NewHandler(store *store.Stores) *Handler {
	return &Handler{
		Store: store,
	}
}

func (handler *Handler) JsonWriteResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
