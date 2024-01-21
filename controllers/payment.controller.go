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

// PaymentController represents the controller for Payment operations
type PaymentController struct {
	// Define any dependencies or services required by the controller
}

// GetAllPayments handles GET request to fetch all payments
func (pc *PaymentController) GetAllPayments(c echo.Context) error {
	payments, err := models.GetAllPayments()
	if err != nil {
		utils.LogError("GetAllPayments", "Failed to fetch payments", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to fetch payments")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: payments})
}

// GetPaymentByID handles GET request to fetch a payment by ID
func (pc *PaymentController) GetPaymentByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.LogError("GetPaymentByID", "Invalid payment ID", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid payment ID")})
	}

	payment, err := models.GetPaymentByID(uint(id))
	if err != nil {
		utils.LogError("GetPaymentByID", "Payment not found", err)
		return api.Response(c, http.StatusNotFound, api.ResponseType{Error: errors.New("Payment not found")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: payment})
}

// CreatePayment handles POST request to create a new payment
func (pc *PaymentController) CreatePayment(c echo.Context) error {
	payment := new(models.Payment)
	if err := c.Bind(payment); err != nil {
		return err
	}

	if err := models.CreatePayment(payment); err != nil {
		utils.LogError("CreatePayment", "Failed to create payment", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to create payment")})
	}
	return api.Response(c, http.StatusCreated, api.ResponseType{Data: payment})
}

// UpdatePayment handles PUT request to update a payment by ID
func (pc *PaymentController) UpdatePayment(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.LogError("UpdatePayment", "Invalid payment ID", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid payment ID")})
	}

	updatedPayment := new(models.Payment)
	if err := c.Bind(updatedPayment); err != nil {
		return err
	}

	if err := models.UpdatePayment(uint(id), updatedPayment); err != nil {
		utils.LogError("UpdatePayment", "Failed to update payment", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to update payment")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: updatedPayment})
}

// DeletePayment handles DELETE request to delete a payment by ID
func (pc *PaymentController) DeletePayment(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.LogError("DeletePayment", "Invalid payment ID", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid payment ID")})
	}

	if err := models.DeletePayment(uint(id)); err != nil {
		utils.LogError("DeletePayment", "Failed to delete payment", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to delete payment")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Message: "Payment deleted"})
}
