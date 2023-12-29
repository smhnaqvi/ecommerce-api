package controllers

import (
	"ecommerce/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllPayments(c echo.Context) error {
	payments, err := models.GetAllPayments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch payments"})
	}
	return c.JSON(http.StatusOK, payments)
}

func GetPaymentByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payment ID"})
	}

	payment, err := models.GetPaymentByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Payment not found"})
	}
	return c.JSON(http.StatusOK, payment)
}

func CreatePayment(c echo.Context) error {
	payment := new(models.Payment)
	if err := c.Bind(payment); err != nil {
		return err
	}

	if err := models.CreatePayment(payment); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create payment"})
	}
	return c.JSON(http.StatusCreated, payment)
}

func UpdatePayment(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payment ID"})
	}

	updatedPayment := new(models.Payment)
	if err := c.Bind(updatedPayment); err != nil {
		return err
	}

	if err := models.UpdatePayment(uint(id), updatedPayment); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update payment"})
	}
	return c.JSON(http.StatusOK, updatedPayment)
}

func DeletePayment(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payment ID"})
	}

	if err := models.DeletePayment(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete payment"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Payment deleted"})
}
