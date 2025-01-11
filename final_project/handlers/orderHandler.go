package handlers

import (
	"encoding/json"
	mdl "final_project/models"
	"net/http"
	"strconv"
	"strings"
)

type OrderHandler struct {
	*Handler
}

func NewOrderHandler(handler *Handler) *OrderHandler {
	return &OrderHandler{
		Handler: handler,
	}
}

func (handler *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := handler.Store.Orders.GetAllOrders()

	if err != nil {
		handler.JsonWriteResponse(w, http.StatusInternalServerError, err.Error())
	}

	handler.JsonWriteResponse(w, http.StatusOK, orders)

}

func (handler *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
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

	order, err := handler.Store.Orders.GetOrder(id)

	if err != nil {
		handler.JsonWriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	handler.JsonWriteResponse(w, http.StatusOK, order)
}

func (handler *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order mdl.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusBadRequest, "invalid URL")
		return
	}
	err = order.Validate()
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	order, err = handler.Store.Orders.CreateOrder(order)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.JsonWriteResponse(w, http.StatusCreated, order)
}

func (handler *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
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

	var order mdl.Order
	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err = handler.Store.Orders.UpdateOrder(id, order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	handler.JsonWriteResponse(w, http.StatusOK, order)
}

func (handler *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
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

	err = handler.Store.Orders.DeleteOrder(id)
	if err != nil {
		handler.JsonWriteResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.JsonWriteResponse(w, http.StatusNoContent, nil)
}

func (handler *OrderHandler) OrdersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handler.GetOrders(w, r)
	case "POST":
		handler.CreateOrder(w, r)
	default:
		handler.JsonWriteResponse(w, http.StatusMethodNotAllowed, nil)
	}
}

func (handler *OrderHandler) OrderHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handler.GetOrder(w, r)
	case "PUT":
		handler.UpdateOrder(w, r)
	case "DELETE":
		handler.DeleteOrder(w, r)
	default:
		handler.JsonWriteResponse(w, http.StatusMethodNotAllowed, nil)
	}
}
