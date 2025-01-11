package handlers

import (
	"encoding/json"
	mdl "final_project/models"
	"net/http"
	"strconv"
	"strings"
)

type BookHandler struct {
	*Handler
}

func NewBookHandler(handler *Handler) *BookHandler {
	return &BookHandler{
		Handler: handler,
	}
}

func (handler *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book mdl.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err = handler.Store.Books.CreateBook(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handler.JsonWriteResponse(w, http.StatusCreated, book)
}

func (handler *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 3 {
		handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid URL")
		//http.Error(w, "invalid URL", )
		return
	}

	arg_id := paths[2]
	id, err := strconv.Atoi(arg_id)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid ID")
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book, err := handler.Store.Books.GetBook(id)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusNotFound, err.Error())
		//http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	handler.JsonWriteResponse(w, http.StatusOK, book)
}

func (handler *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 3 {
		handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid URL")
		//http.Error(w, "invalid URL", )
		return
	}

	arg_id := paths[2]
	id, err := strconv.Atoi(arg_id)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid ID")
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

	handler.JsonWriteResponse(w, http.StatusOK, book)
}

func (handler *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 3 {
		handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid URL")
		//http.Error(w, "invalid URL", )
		return
	}

	arg_id := paths[2]
	id, err := strconv.Atoi(arg_id)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid ID")
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.Store.Books.DeleteBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handler.JsonWriteResponse(w, http.StatusOK, "Book deleted")
}

func (handler *BookHandler) SearchBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	criteria := mdl.SearchCriteria{
		Title:  query.Get("title"),
		Author: query.Get("author"),
	}
	minPrice := query.Get("min_price")
	if minPrice != "" {
		price, err := strconv.ParseFloat(minPrice, 64)

		if err != nil {
			handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid min_price")
			return
		}
		criteria.MinPrice = &price
	}

	maxPrice := query.Get("max_price")
	if maxPrice != "" {
		price, err := strconv.ParseFloat(maxPrice, 64)

		if err != nil {
			handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid max_price")
			return
		}
		criteria.MaxPrice = &price
	}

	inStock := query.Get("in_stock")

	if inStock != "" {
		exists, err := strconv.ParseBool(inStock)
		if err != nil {
			handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid in_stock")
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
		handler.JsonWriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.JsonWriteResponse(w, http.StatusOK, books)
}

func (handler *BookHandler) BooksRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler.CreateBook(w, r)
	case http.MethodGet:
		handler.SearchBooks(w, r)
	default:
		handler.JsonWriteResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}

}

func (handler *BookHandler) BookRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handler.GetBook(w, r)
	case http.MethodPut:
		handler.UpdateBook(w, r)
	case http.MethodDelete:
		handler.DeleteBook(w, r)
	default:
		handler.JsonWriteResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
