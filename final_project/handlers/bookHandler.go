package handlers

import (
	"context"
	"encoding/json"
	mdl "final_project/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type BookHandler struct {
	*Handler
}

func NewBookHandler(handler *Handler) *BookHandler {
	return &BookHandler{
		Handler: handler,
	}
}

func (handler *BookHandler) CreateBook(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var book mdl.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		handler.JsonWriteResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	book, err = handler.Store.Books.CreateBook(book)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		handler.JsonWriteResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	handler.JsonWriteResponse(w, r, http.StatusCreated, book)
}

func (handler *BookHandler) GetBook(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 3 {
		handler.JsonWriteResponse(w, r, http.StatusBadRequest, "invalid URL")
		//http.Error(w, "invalid URL", )
		return
	}

	arg_id := paths[2]
	id, err := strconv.Atoi(arg_id)
	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusBadRequest, "invalid ID")
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book, err := handler.Store.Books.GetBook(id)
	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusNotFound, err.Error())
		//http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	handler.JsonWriteResponse(w, r, http.StatusOK, book)
}

func (handler *BookHandler) UpdateBook(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 3 {
		handler.JsonWriteResponse(w, r, http.StatusBadRequest, "invalid URL")
		//http.Error(w, "invalid URL", )
		return
	}

	arg_id := paths[2]
	id, err := strconv.Atoi(arg_id)
	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusBadRequest, "invalid ID")
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var book mdl.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err = handler.Store.Books.UpdateBook(id, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handler.JsonWriteResponse(w, r, http.StatusOK, book)
}

func (handler *BookHandler) DeleteBook(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 3 {
		handler.JsonWriteResponse(w, r, http.StatusBadRequest, "invalid URL")
		//http.Error(w, "invalid URL", )
		return
	}

	arg_id := paths[2]
	id, err := strconv.Atoi(arg_id)
	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusBadRequest, "invalid ID")
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.Store.Books.DeleteBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handler.JsonWriteResponse(w, r, http.StatusOK, "Book deleted")
}

func (handler *BookHandler) SearchBooks(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	criteria := mdl.SearchCriteria{
		Title:  query.Get("title"),
		Author: query.Get("author"),
	}
	minPrice := query.Get("min_price")
	if minPrice != "" {
		price, err := strconv.ParseFloat(minPrice, 64)

		if err != nil {
			handler.JsonWriteResponse(w, r, http.StatusBadRequest, "invalid min_price")
			return
		}
		criteria.MinPrice = &price
	}

	maxPrice := query.Get("max_price")
	if maxPrice != "" {
		price, err := strconv.ParseFloat(maxPrice, 64)

		if err != nil {
			handler.JsonWriteResponse(w, r, http.StatusBadRequest, "invalid max_price")
			return
		}
		criteria.MaxPrice = &price
	}

	inStock := query.Get("in_stock")

	if inStock != "" {
		exists, err := strconv.ParseBool(inStock)
		if err != nil {
			handler.JsonWriteResponse(w, r, http.StatusBadRequest, "invalid in_stock")
			return
		}
		criteria.InStock = &exists
	}

	genres := query.Get("genres")
	if genres != "" {
		criteria.Genres = strings.Split(genres, ",")
	}
	books, err := handler.Store.Books.SearchBooks(criteria)
	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	handler.JsonWriteResponse(w, r, http.StatusOK, books)
}

func (handler *BookHandler) BooksRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler.withTimeout(10*time.Second, handler.CreateBook)(w, r)
	case http.MethodGet:
		handler.withTimeout(10*time.Second, handler.SearchBooks)(w, r)
	default:
		handler.JsonWriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed")
	}

}

func (handler *BookHandler) BookRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handler.withTimeout(10*time.Second, handler.GetBook)(w, r)
	case http.MethodPut:
		handler.withTimeout(10*time.Second, handler.UpdateBook)(w, r)
	case http.MethodDelete:
		handler.withTimeout(10*time.Second, handler.DeleteBook)(w, r)
	default:
		handler.JsonWriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
