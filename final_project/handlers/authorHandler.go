package handlers

import (
	"encoding/json"
	mdl "final_project/models"
	"net/http"
	"strconv"
	"strings"
)

type AuthorHandler struct {
	*Handler
}

func NewAuthorHandler(handler *Handler) *AuthorHandler {
	return &AuthorHandler{
		Handler: handler,
	}
}

func (handler *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author mdl.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid URL")
		return
	}

	author, err = handler.Store.Authors.CreateAuthor(author)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		handler.JsonWriteResponse(w, http.StatusInternalServerError, err.Error())

		return
	}

	handler.JsonWriteResponse(w, http.StatusCreated, author)
}

func (handler *AuthorHandler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 3 {
		handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid URL")
		return
	}

	arg_id := paths[2]
	id, err := strconv.Atoi(arg_id)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid ID")
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	author, err := handler.Store.Authors.GetAuthor(id)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusNotFound, err.Error())
		//http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	handler.JsonWriteResponse(w, http.StatusOK, author)
}

func (handler *AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 3 {
		handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid URL")
		return
	}

	arg_id := paths[2]
	id, err := strconv.Atoi(arg_id)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid ID")
		return
	}

	var author mdl.Author
	err = json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	author, err = handler.Store.Authors.UpdateAuthor(id, author)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	handler.JsonWriteResponse(w, http.StatusOK, author)
}

func (handler *AuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 3 {
		handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid URL")
		return
	}

	arg_id := paths[2]
	id, err := strconv.Atoi(arg_id)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid ID")
		return
	}

	err = handler.Store.Authors.DeleteAuthor(id)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	handler.JsonWriteResponse(w, http.StatusOK, "Author deleted")
}

func (handler *AuthorHandler) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := handler.Store.Authors.GetAllAuthors()
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.JsonWriteResponse(w, http.StatusOK, authors)
}

func (handler *AuthorHandler) AuthorsRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handler.GetAllAuthors(w, r)
	case http.MethodPost:
		handler.CreateAuthor(w, r)
	default:
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}

func (handler *AuthorHandler) AuthorRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handler.GetAuthor(w, r)
	case http.MethodPut:
		handler.UpdateAuthor(w, r)
	case http.MethodDelete:
		handler.DeleteAuthor(w, r)
	default:
		handler.JsonWriteResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
