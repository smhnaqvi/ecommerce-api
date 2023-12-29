package controllers

import (
	"ecommerce/models"
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
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch order details"})
	}
	return c.JSON(http.StatusOK, orderDetails)
}

// GetOrderDetailByID handles GET request to fetch an order detail by ID
func (oc *OrderDetailController) GetOrderDetailByID(c echo.Context) error {
	id := c.Param("id")
	orderDetail, err := models.GetOrderDetailByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Order detail not found"})
	}
	return c.JSON(http.StatusOK, orderDetail)
}

// CreateOrderDetail handles POST request to create a new order detail
func (oc *OrderDetailController) CreateOrderDetail(c echo.Context) error {
	orderDetail := new(models.OrderDetail)
	if err := c.Bind(orderDetail); err != nil {
		return err
	}
	if err := models.CreateOrderDetail(orderDetail); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create order detail"})
	}
	return c.JSON(http.StatusCreated, orderDetail)
}

// UpdateOrderDetail handles PUT request to update an order detail by ID
func (oc *OrderDetailController) UpdateOrderDetail(c echo.Context) error {
	id := c.Param("id")
	existingOrderDetail, err := models.GetOrderDetailByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Order detail not found"})
	}

	updatedOrderDetail := new(models.OrderDetail)
	if err := c.Bind(updatedOrderDetail); err != nil {
		return err
	}

	if err := models.UpdateOrderDetail(existingOrderDetail, updatedOrderDetail); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update order detail"})
	}
	return c.JSON(http.StatusOK, existingOrderDetail)
}

// DeleteOrderDetail handles DELETE request to delete an order detail by ID
func (oc *OrderDetailController) DeleteOrderDetail(c echo.Context) error {
	id := c.Param("id")
	orderDetail, err := models.GetOrderDetailByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Order detail not found"})
	}

	if err := models.DeleteOrderDetail(orderDetail); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete order detail"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Order detail deleted successfully"})
}
