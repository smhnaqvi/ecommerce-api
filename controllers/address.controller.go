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

type AddressController struct {
	// Define any dependencies or services required by the controller
}

// GetAddresses handles GET request to fetch all addresses
func (ac *AddressController) GetAddresses(c echo.Context) error {
	addresses, err := models.GetAllAddresses()
	if err != nil {
		utils.LogError("GetAddresses", "Failed to fetch addresses", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to fetch addresses")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: addresses})
}

// GetAddressByID handles GET request to fetch an address by ID
func (ac *AddressController) GetAddressByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.LogError("GetAddressByID", "Invalid ID format", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid ID format")})
	}

	address, err := models.GetAddressByID(uint(id))
	if err != nil {
		utils.LogError("GetAddressByID", "Failed to fetch address by ID", err)
		return api.Response(c, http.StatusNotFound, api.ResponseType{Error: errors.New("Address not found")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: address})
}

// CreateAddress handles POST request to create a new address
func (ac *AddressController) CreateAddress(c echo.Context) error {
	address := new(models.Address)
	if err := c.Bind(address); err != nil {
		utils.LogError("CreateAddress", "Invalid request payload", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: err})
	}

	if err := models.CreateAddress(address); err != nil {
		utils.LogError("CreateAddress", "Failed to create address", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to create address")})
	}

	return api.Response(c, http.StatusCreated, api.ResponseType{Data: address})
}

// UpdateAddress handles PUT request to update an address by ID
func (ac *AddressController) UpdateAddress(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.LogError("UpdateAddress", "Invalid ID format", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid ID format")})
	}

	updatedAddress := new(models.Address)
	if err := c.Bind(updatedAddress); err != nil {
		utils.LogError("UpdateAddress", "Invalid request payload", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: err})
	}

	if err := models.UpdateAddress(uint(id), updatedAddress); err != nil {
		utils.LogError("UpdateAddress", "Failed to update address", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to update address")})
	}

	return api.Response(c, http.StatusOK, api.ResponseType{Data: updatedAddress})
}

// DeleteAddress handles DELETE request to delete an address by ID
func (ac *AddressController) DeleteAddress(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.LogError("DeleteAddress", "Invalid ID format", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid ID format")})
	}

	if err := models.DeleteAddress(uint(id)); err != nil {
		utils.LogError("DeleteAddress", "Failed to delete address", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to delete address")})
	}

	return api.Response(c, http.StatusOK, api.ResponseType{Message: "Address deleted"})
}
