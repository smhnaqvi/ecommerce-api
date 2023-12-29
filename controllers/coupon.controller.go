package controllers

import (
	"ecommerce/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CouponController struct {
	// Any dependencies or services required by the controller can be defined here
}

func (cc *CouponController) GetAllCoupons(c echo.Context) error {
	coupons, err := models.GetAllCoupons()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch coupons"})
	}
	return c.JSON(http.StatusOK, coupons)
}

func (cc *CouponController) GetCouponByID(c echo.Context) error {
	id := c.Param("id")
	coupon, err := models.GetCouponByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Coupon not found"})
	}
	return c.JSON(http.StatusOK, coupon)
}

func (cc *CouponController) CreateCoupon(c echo.Context) error {
	coupon := new(models.Coupon)
	if err := c.Bind(coupon); err != nil {
		return err
	}
	if err := models.CreateCoupon(coupon); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create coupon"})
	}
	return c.JSON(http.StatusCreated, coupon)
}

func (cc *CouponController) UpdateCoupon(c echo.Context) error {
	id := c.Param("id")
	existingCoupon, err := models.GetCouponByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Coupon not found"})
	}

	updatedCoupon := new(models.Coupon)
	if err := c.Bind(updatedCoupon); err != nil {
		return err
	}

	if err := models.UpdateCoupon(existingCoupon, updatedCoupon); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update coupon"})
	}
	return c.JSON(http.StatusOK, existingCoupon)
}

func (cc *CouponController) DeleteCoupon(c echo.Context) error {
	id := c.Param("id")
	coupon, err := models.GetCouponByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Coupon not found"})
	}

	if err := models.DeleteCoupon(coupon); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete coupon"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Coupon deleted successfully"})
}

