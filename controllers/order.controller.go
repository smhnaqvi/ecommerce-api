package controllers

import (
	"ecommerce/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
    // Any dependencies or services required by the controller can be defined here
}

func (oc *OrderController) GetAllOrders(c echo.Context) error {
    orders, err := models.GetAllOrders()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch orders"})
    }
    return c.JSON(http.StatusOK, orders)
}

func (oc *OrderController) GetOrderByID(c echo.Context) error {
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid order ID"})
    }

    order, err := models.GetOrderByID(uint(id))
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Order not found"})
    }
    return c.JSON(http.StatusOK, order)
}

func (oc *OrderController) CreateOrder(c echo.Context) error {
    var order models.Order
    if err := c.Bind(&order); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
    }

    if err := models.CreateOrder(&order); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create order"})
    }
    return c.JSON(http.StatusCreated, order)
}

func (oc *OrderController) UpdateOrder(c echo.Context) error {
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid order ID"})
    }

    var updatedOrder models.Order
    if err := c.Bind(&updatedOrder); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
    }

    if err := models.UpdateOrder(uint(id), &updatedOrder); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update order"})
    }
    return c.JSON(http.StatusOK, updatedOrder)
}

func (oc *OrderController) DeleteOrder(c echo.Context) error {
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid order ID"})
    }

    if err := models.DeleteOrder(uint(id)); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete order"})
    }
    return c.JSON(http.StatusOK, map[string]string{"message": "Order deleted"})
}
