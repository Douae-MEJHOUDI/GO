package handlers

import (
	"context"
	"encoding/json"
	"final_project/store"
	"log"
	"net/http"
	"os"
	"time"
)

type Handler struct {
	Store  *store.Stores
	logger *log.Logger
}

func NewHandler(store *store.Stores) *Handler {
	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return &Handler{
		Store:  store,
		logger: log.New(logFile, "", log.LstdFlags),
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
				h.JsonWriteResponse(w, r, http.StatusGatewayTimeout, "request timed out")
			case context.Canceled:
				h.JsonWriteResponse(w, r, http.StatusBadRequest, "request was canceled")
			}
		case <-done:
			log.Println("Request completed")
		}
	}
}

func (handler *Handler) JsonWriteResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	start := time.Now()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)

	handler.logRequest(r, status, time.Since(start))

}

func (h *Handler) logRequest(r *http.Request, status int, duration time.Duration) {
	h.logger.Printf(
		"Method: %s, Path: %s, Status: %d, Duration: %v",
		r.Method,
		r.URL.Path,
		status,
		duration,
	)
}
