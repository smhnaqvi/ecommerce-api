package controllers

import (
	"ecommerce/api"
	"ecommerce/models"
	"ecommerce/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	// Any dependencies or services required by the controller can be defined here
}

// GetAllOrders handles GET request to fetch all orders
func (oc *OrderController) GetAllOrders(c echo.Context) error {
	orders, err := models.GetAllOrders()
	if err != nil {
		utils.LogError("GetAllOrders", "Failed to fetch orders", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to fetch orders")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: orders})
}

// GetOrderByID handles GET request to fetch an order by ID
func (oc *OrderController) GetOrderByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.LogError("GetOrderByID", "Invalid order ID", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid order ID")})
	}

	order, err := models.GetOrderByID(uint(id))
	if err != nil {
		utils.LogError("GetOrderByID", "Order not found", err)
		return api.Response(c, http.StatusNotFound, api.ResponseType{Error: errors.New("Order not found")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: order})
}

// CreateOrder handles POST request to create a new order
func (oc *OrderController) CreateOrder(c echo.Context) error {
	var order models.Order
	if err := c.Bind(&order); err != nil {
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid request payload")})
	}

	if err := models.CreateOrder(&order); err != nil {
		utils.LogError("CreateOrder", "Failed to create order", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to create order")})
	}
	return api.Response(c, http.StatusCreated, api.ResponseType{Data: order})
}

// UpdateOrder handles PUT request to update an order by ID
func (oc *OrderController) UpdateOrder(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.LogError("UpdateOrder", "Invalid order ID", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid order ID")})
	}

	var updatedOrder models.Order
	if err := c.Bind(&updatedOrder); err != nil {
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid request payload")})
	}

	if err := models.UpdateOrder(uint(id), &updatedOrder); err != nil {
		utils.LogError("UpdateOrder", "Failed to update order", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to update order")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: updatedOrder})
}

// DeleteOrder handles DELETE request to delete an order by ID
func (oc *OrderController) DeleteOrder(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.LogError("DeleteOrder", "Invalid order ID", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid order ID")})
	}

	if err := models.DeleteOrder(uint(id)); err != nil {
		utils.LogError("DeleteOrder", "Failed to delete order", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to delete order")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Message: "Order deleted"})
}
