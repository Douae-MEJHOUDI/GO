package handlers

import (
	"encoding/json"
	mdl "final_project/models"
	"net/http"
	"strconv"
	"strings"
)

type CustomerHandler struct {
	*Handler
}

func NewCustomerHandler(handler *Handler) *CustomerHandler {
	return &CustomerHandler{
		Handler: handler,
	}
}

func (handler *CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := handler.Store.Customers.GetAllCustomers()

	if err != nil {
		handler.JsonWriteResponse(w, http.StatusInternalServerError, err.Error())
	}

	handler.JsonWriteResponse(w, http.StatusOK, customers)
}

func (handler *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
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

func (handler *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
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

func (handler *CustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
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

func (handler *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
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
		handler.GetCustomers(w, r)
	case "POST":
		handler.CreateCustomer(w, r)
	default:
		handler.JsonWriteResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (handler *CustomerHandler) CustomerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handler.GetCustomer(w, r)
	case "PUT":
		handler.UpdateCustomer(w, r)
	case "DELETE":
		handler.DeleteCustomer(w, r)
	default:
		handler.JsonWriteResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
