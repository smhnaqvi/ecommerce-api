package controllers

import (
	"ecommerce/api"
	"ecommerce/models"
	"ecommerce/utils"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CouponController represents the controller for Coupon operations
type CouponController struct {
	// Any dependencies or services required by the controller can be defined here
}

// GetAllCoupons handles GET request to fetch all coupons
func (cc *CouponController) GetAllCoupons(c echo.Context) error {
	coupons, err := models.GetAllCoupons()
	if err != nil {
		utils.LogError("GetAllCoupons", "Failed to fetch coupons", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to fetch coupons")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: coupons})
}

// GetCouponByID handles GET request to fetch a coupon by ID
func (cc *CouponController) GetCouponByID(c echo.Context) error {
	id := c.Param("id")
	coupon, err := models.GetCouponByID(id)
	if err != nil {
		utils.LogError("GetCouponByID", "Coupon not found", err)
		return api.Response(c, http.StatusNotFound, api.ResponseType{Error: errors.New("Coupon not found")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: coupon})
}

// CreateCoupon handles POST request to create a new coupon
func (cc *CouponController) CreateCoupon(c echo.Context) error {
	coupon := new(models.Coupon)
	if err := c.Bind(coupon); err != nil {
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid request payload")})
	}

	if err := models.CreateCoupon(coupon); err != nil {
		utils.LogError("CreateCoupon", "Failed to create coupon", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to create coupon")})
	}
	return api.Response(c, http.StatusCreated, api.ResponseType{Data: coupon})
}

// UpdateCoupon handles PUT request to update a coupon by ID
func (cc *CouponController) UpdateCoupon(c echo.Context) error {
	id := c.Param("id")
	existingCoupon, err := models.GetCouponByID(id)
	if err != nil {
		utils.LogError("UpdateCoupon", "Coupon not found", err)
		return api.Response(c, http.StatusNotFound, api.ResponseType{Error: errors.New("Coupon not found")})
	}

	updatedCoupon := new(models.Coupon)
	if err := c.Bind(updatedCoupon); err != nil {
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid request payload")})
	}

	if err := models.UpdateCoupon(existingCoupon, updatedCoupon); err != nil {
		utils.LogError("UpdateCoupon", "Failed to update coupon", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to update coupon")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: existingCoupon})
}

// DeleteCoupon handles DELETE request to delete a coupon by ID
func (cc *CouponController) DeleteCoupon(c echo.Context) error {
	id := c.Param("id")
	coupon, err := models.GetCouponByID(id)
	if err != nil {
		utils.LogError("DeleteCoupon", "Coupon not found", err)
		return api.Response(c, http.StatusNotFound, api.ResponseType{Error: errors.New("Coupon not found")})
	}

	if err := models.DeleteCoupon(coupon); err != nil {
		utils.LogError("DeleteCoupon", "Failed to delete coupon", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to delete coupon")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Message: "Coupon deleted successfully"})
}
