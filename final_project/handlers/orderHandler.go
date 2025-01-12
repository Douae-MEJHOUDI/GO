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

type OrderHandler struct {
	*Handler
}

func NewOrderHandler(handler *Handler) *OrderHandler {
	return &OrderHandler{
		Handler: handler,
	}
}

func (handler *OrderHandler) GetOrders(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	orders, err := handler.Store.Orders.GetAllOrders()

	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusInternalServerError, err.Error())
	}

	handler.JsonWriteResponse(w, r, http.StatusOK, orders)

}

func (handler *OrderHandler) GetOrder(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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

	order, err := handler.Store.Orders.GetOrder(id)

	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusNotFound, err.Error())
		return
	}

	handler.JsonWriteResponse(w, r, http.StatusOK, order)
}

func (handler *OrderHandler) CreateOrder(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var order mdl.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusBadRequest, "invalid URL")
		return
	}
	err = order.Validate()
	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	order, err = handler.Store.Orders.CreateOrder(order)
	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	handler.JsonWriteResponse(w, r, http.StatusCreated, order)
}

func (handler *OrderHandler) UpdateOrder(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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

	handler.JsonWriteResponse(w, r, http.StatusOK, order)
}

func (handler *OrderHandler) DeleteOrder(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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

	err = handler.Store.Orders.DeleteOrder(id)
	if err != nil {
		handler.JsonWriteResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	handler.JsonWriteResponse(w, r, http.StatusNoContent, nil)
}

func (handler *OrderHandler) OrdersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handler.withTimeout(10*time.Second, handler.GetOrders)(w, r)
	case "POST":
		handler.withTimeout(10*time.Second, handler.CreateOrder)(w, r)
	default:
		handler.JsonWriteResponse(w, r, http.StatusMethodNotAllowed, nil)
	}
}

func (handler *OrderHandler) OrderHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handler.withTimeout(10*time.Second, handler.GetOrder)(w, r)
	case "PUT":
		handler.withTimeout(10*time.Second, handler.UpdateOrder)(w, r)
	case "DELETE":
		handler.withTimeout(10*time.Second, handler.DeleteOrder)(w, r)
	default:
		handler.JsonWriteResponse(w, r, http.StatusMethodNotAllowed, nil)
	}
}
