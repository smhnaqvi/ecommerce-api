package controllers

import (
	"ecommerce/api"
	"ecommerce/models"
	"ecommerce/utils"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// OrderDetailController represents the controller for OrderDetail operations
type OrderDetailController struct {
	// Define any dependencies or services required by the controller
}

// GetAllOrderDetails handles GET request to fetch all order details
func (oc *OrderDetailController) GetAllOrderDetails(c echo.Context) error {
	orderDetails, err := models.GetAllOrderDetails()
	if err != nil {
		utils.LogError("GetAllOrderDetails", "Failed to fetch order details", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to fetch order details")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: orderDetails})
}

// GetOrderDetailByID handles GET request to fetch an order detail by ID
func (oc *OrderDetailController) GetOrderDetailByID(c echo.Context) error {
	id := c.Param("id")
	orderDetail, err := models.GetOrderDetailByID(id)
	if err != nil {
		utils.LogError("GetOrderDetailByID", "Order detail not found", err)
		return api.Response(c, http.StatusNotFound, api.ResponseType{Error: errors.New("Order detail not found")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: orderDetail})
}

// CreateOrderDetail handles POST request to create a new order detail
func (oc *OrderDetailController) CreateOrderDetail(c echo.Context) error {
	orderDetail := new(models.OrderDetail)
	if err := c.Bind(orderDetail); err != nil {
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid request payload")})
	}

	if err := models.CreateOrderDetail(orderDetail); err != nil {
		utils.LogError("CreateOrderDetail", "Failed to create order detail", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to create order detail")})
	}
	return api.Response(c, http.StatusCreated, api.ResponseType{Data: orderDetail})
}

// UpdateOrderDetail handles PUT request to update an order detail by ID
func (oc *OrderDetailController) UpdateOrderDetail(c echo.Context) error {
	id := c.Param("id")
	existingOrderDetail, err := models.GetOrderDetailByID(id)
	if err != nil {
		utils.LogError("UpdateOrderDetail", "Order detail not found", err)
		return api.Response(c, http.StatusNotFound, api.ResponseType{Error: errors.New("Order detail not found")})
	}

	updatedOrderDetail := new(models.OrderDetail)
	if err := c.Bind(updatedOrderDetail); err != nil {
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid request payload")})
	}

	if err := models.UpdateOrderDetail(existingOrderDetail, updatedOrderDetail); err != nil {
		utils.LogError("UpdateOrderDetail", "Failed to update order detail", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to update order detail")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: existingOrderDetail})
}

// DeleteOrderDetail handles DELETE request to delete an order detail by ID
func (oc *OrderDetailController) DeleteOrderDetail(c echo.Context) error {
	id := c.Param("id")
	orderDetail, err := models.GetOrderDetailByID(id)
	if err != nil {
		utils.LogError("DeleteOrderDetail", "Order detail not found", err)
		return api.Response(c, http.StatusNotFound, api.ResponseType{Error: errors.New("Order detail not found")})
	}

	if err := models.DeleteOrderDetail(orderDetail); err != nil {
		utils.LogError("DeleteOrderDetail", "Failed to delete order detail", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to delete order detail")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Message: "Order detail deleted successfully"})
}
