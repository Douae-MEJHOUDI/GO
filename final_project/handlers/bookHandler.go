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
	var criteria mdl.SearchCriteria
	err := json.NewDecoder(r.Body).Decode(&criteria)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		handler.JsonWriteResponse(w, http.StatusBadRequest, err.Error()+"here za3ma?")
		return
	}

	books, err := handler.Store.Books.SearchBooks(criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
