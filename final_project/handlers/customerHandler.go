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

type CustomerHandler struct {
	*Handler
}

func NewCustomerHandler(handler *Handler) *CustomerHandler {
	return &CustomerHandler{
		Handler: handler,
	}
}

func (handler *CustomerHandler) GetCustomers(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	customers, err := handler.Store.Customers.GetAllCustomers()

	if err != nil {
		handler.JsonWriteResponse(w, http.StatusInternalServerError, err.Error())
	}

	handler.JsonWriteResponse(w, http.StatusOK, customers)
}

func (handler *CustomerHandler) GetCustomer(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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
	customer, err := handler.Store.Customers.GetCustomer(id)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	handler.JsonWriteResponse(w, http.StatusOK, customer)

}

func (handler *CustomerHandler) CreateCustomer(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var customer mdl.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid URL")
		return
	}

	customer, err = handler.Store.Customers.CreateCustomer(customer)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.JsonWriteResponse(w, http.StatusCreated, customer)
}

func (handler *CustomerHandler) UpdateCustomer(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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

	var customer mdl.Customer
	err = json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid URL")
		return
	}

	customer, err = handler.Store.Customers.UpdateCustomer(id, customer)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.JsonWriteResponse(w, http.StatusOK, customer)
}

func (handler *CustomerHandler) DeleteCustomer(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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

	err = handler.Store.Customers.DeleteCustomer(id)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.JsonWriteResponse(w, http.StatusOK, "Customer deleted")
}

func (handler *CustomerHandler) CustomersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handler.withTimeout(10*time.Second, handler.GetCustomers)(w, r)
	case "POST":
		handler.withTimeout(10*time.Second, handler.CreateCustomer)(w, r)
	default:
		handler.JsonWriteResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (handler *CustomerHandler) CustomerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handler.withTimeout(10*time.Second, handler.GetCustomer)(w, r)
	case "PUT":
		handler.withTimeout(10*time.Second, handler.UpdateCustomer)(w, r)
	case "DELETE":
		handler.withTimeout(10*time.Second, handler.DeleteCustomer)(w, r)
	default:
		handler.JsonWriteResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
