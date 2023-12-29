package controllers

import (
	"ecommerce/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AddressController struct {
	// Define any dependencies or services required by the controller
}

// GetAddresses handles GET request to fetch all addresses
func (ac *AddressController) GetAddresses(c echo.Context) error {
	addresses, err := models.GetAllAddresses()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch addresses"})
	}
	return c.JSON(http.StatusOK, addresses)
}

// GetAddressByID handles GET request to fetch an address by ID
func (ac *AddressController) GetAddressByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	address, err := models.GetAddressByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Address not found"})
	}
	return c.JSON(http.StatusOK, address)
}

// CreateAddress handles POST request to create a new address
func (ac *AddressController) CreateAddress(c echo.Context) error {
	address := new(models.Address)
	if err := c.Bind(address); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := models.CreateAddress(address); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create address"})
	}
	return c.JSON(http.StatusCreated, address)
}

// UpdateAddress handles PUT request to update an address by ID
func (ac *AddressController) UpdateAddress(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	updatedAddress := new(models.Address)
	if err := c.Bind(updatedAddress); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := models.UpdateAddress(uint(id), updatedAddress); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update address"})
	}
	return c.JSON(http.StatusOK, updatedAddress)
}

// DeleteAddress handles DELETE request to delete an address by ID
func (ac *AddressController) DeleteAddress(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	if err := models.DeleteAddress(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete address"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Address deleted"})
}
