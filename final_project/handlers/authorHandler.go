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

type AuthorHandler struct {
	*Handler
}

func NewAuthorHandler(handler *Handler) *AuthorHandler {
	return &AuthorHandler{
		Handler: handler,
	}
}

func (handler *AuthorHandler) CreateAuthor(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var author mdl.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		handler.JsonWriteResponse(w, r, http.StatusBadRequest, "invalid URL")
		return
	}

	author, err = handler.Store.Authors.CreateAuthor(author)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		handler.JsonWriteResponse(w, r, http.StatusInternalServerError, err.Error())

		return
	}

	handler.JsonWriteResponse(w, r, http.StatusCreated, author)
}

func (handler *AuthorHandler) GetAuthor(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 3 {
		handler.JsonWriteResponse(w, r, http.StatusBadRequest, "invalid URL")
		return
	}

	arg_id := paths[2]
	id, err := strconv.Atoi(arg_id)
	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusBadRequest, "invalid ID")
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	author, err := handler.Store.Authors.GetAuthor(id)
	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusNotFound, err.Error())
		//http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	handler.JsonWriteResponse(w, r, http.StatusOK, author)
}

func (handler *AuthorHandler) UpdateAuthor(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 3 {
		handler.JsonWriteResponse(w, r, http.StatusBadRequest, "invalid URL")
		return
	}

	arg_id := paths[2]
	id, err := strconv.Atoi(arg_id)
	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusBadRequest, "invalid ID")
		return
	}

	var author mdl.Author
	err = json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	author, err = handler.Store.Authors.UpdateAuthor(id, author)
	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusNotFound, err.Error())
		return
	}

	handler.JsonWriteResponse(w, r, http.StatusOK, author)
}

func (handler *AuthorHandler) DeleteAuthor(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 3 {
		handler.JsonWriteResponse(w, r, http.StatusBadRequest, "invalid URL")
		return
	}

	arg_id := paths[2]
	id, err := strconv.Atoi(arg_id)
	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusBadRequest, "invalid ID")
		return
	}

	err = handler.Store.Authors.DeleteAuthor(id)
	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusNotFound, err.Error())
		return
	}

	handler.JsonWriteResponse(w, r, http.StatusOK, "Author deleted")
}

func (handler *AuthorHandler) GetAllAuthors(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	authors, err := handler.Store.Authors.GetAllAuthors()
	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	handler.JsonWriteResponse(w, r, http.StatusOK, authors)
}

func (handler *AuthorHandler) AuthorsRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handler.withTimeout(10*time.Second, handler.GetAllAuthors)(w, r)
	case http.MethodPost:
		handler.withTimeout(10*time.Second, handler.CreateAuthor)(w, r)
	default:
		handler.JsonWriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (handler *AuthorHandler) AuthorRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handler.withTimeout(10*time.Second, handler.GetAuthor)(w, r)
	case http.MethodPut:
		handler.withTimeout(10*time.Second, handler.UpdateAuthor)(w, r)
	case http.MethodDelete:
		handler.withTimeout(10*time.Second, handler.DeleteAuthor)(w, r)
	default:
		handler.JsonWriteResponse(w, r, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
