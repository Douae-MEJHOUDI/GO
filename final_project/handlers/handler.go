package handlers

import (
	"context"
	"encoding/json"
	"final_project/store"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	Store *store.Stores
}

func NewHandler(store *store.Stores) *Handler {
	return &Handler{
		Store: store,
	}
}

type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func (h *Handler) withTimeout(duration time.Duration, fn HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), duration)
		defer cancel()

		done := make(chan struct{})
		go func() {
			fn(ctx, w, r)
			done <- struct{}{}
		}()

		select {
		case <-ctx.Done():
			switch ctx.Err() {
			case context.DeadlineExceeded:
				h.JsonWriteResponse(w, http.StatusGatewayTimeout, "request timed out")
			case context.Canceled:
				h.JsonWriteResponse(w, http.StatusBadRequest, "request was canceled")
			}
		case <-done:
			log.Println("Request completed")
		}
	}
}

func (handler *Handler) JsonWriteResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
