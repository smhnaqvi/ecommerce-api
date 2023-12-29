// For ShoppingCart model
package models

import (
	"net/http"

	DB "ecommerce/database"

	"github.com/labstack/echo/v4"
)

type ShoppingCart struct {
	CartID    uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	ProductID uint
	Quantity  int `gorm:"not null"`
	// Other shopping cart-related fields as needed
}

// GetAllShoppingCarts returns all shopping carts
func GetAllShoppingCarts(c echo.Context) error {
	var shoppingCarts []ShoppingCart
	DB.Connection.Find(&shoppingCarts)
	return c.JSON(http.StatusOK, shoppingCarts)
}

// GetShoppingCartByID returns a shopping cart by ID
func GetShoppingCartByID(c echo.Context) error {
	id := c.Param("id")
	var shoppingCart ShoppingCart
	if err := DB.Connection.First(&shoppingCart, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Shopping cart not found"})
	}
	return c.JSON(http.StatusOK, shoppingCart)
}

// CreateShoppingCart creates a new shopping cart
func CreateShoppingCart(c echo.Context) error {
	shoppingCart := new(ShoppingCart)
	if err := c.Bind(shoppingCart); err != nil {
		return err
	}
	DB.Connection.Create(&shoppingCart)
	return c.JSON(http.StatusCreated, shoppingCart)
}

// UpdateShoppingCart updates a shopping cart
func UpdateShoppingCart(c echo.Context) error {
	id := c.Param("id")
	var shoppingCart ShoppingCart
	if err := DB.Connection.First(&shoppingCart, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Shopping cart not found"})
	}

	updatedShoppingCart := new(ShoppingCart)
	if err := c.Bind(updatedShoppingCart); err != nil {
		return err
	}

	DB.Connection.Model(&shoppingCart).Updates(updatedShoppingCart)
	return c.JSON(http.StatusOK, shoppingCart)
}

// DeleteShoppingCart deletes a shopping cart
func DeleteShoppingCart(c echo.Context) error {
	id := c.Param("id")
	var shoppingCart ShoppingCart
	if err := DB.Connection.First(&shoppingCart, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Shopping cart not found"})
	}

	DB.Connection.Delete(&shoppingCart)
	return c.JSON(http.StatusOK, map[string]string{"message": "Shopping cart deleted"})
}
// Similarly, adapt this template for Payment, Address, and Coupon models
