package controllers

import (
	"ecommerce/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct{}

// GetProducts handles GET request to fetch all products
func (pc *ProductController) GetProducts(c echo.Context) error {
	products, err := models.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch products"})
	}
	return c.JSON(http.StatusOK, products)
}

// GetProductByID handles GET request to fetch a product by ID
func (pc *ProductController) GetProductByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	product, err := models.GetProductByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}
	return c.JSON(http.StatusOK, product)
}

// CreateProduct handles POST request to create a new product
func (pc *ProductController) CreateProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := models.CreateProduct(&product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create product"})
	}
	return c.JSON(http.StatusCreated, product)
}

// UpdateProduct handles PUT request to update a product by ID
func (pc *ProductController) UpdateProduct(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	var updatedProduct models.Product
	if err := c.Bind(&updatedProduct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := models.UpdateProduct(uint(id), &updatedProduct); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update product"})
	}
	return c.JSON(http.StatusOK, updatedProduct)
}

// DeleteProduct handles DELETE request to delete a product by ID
func (pc *ProductController) DeleteProduct(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	if err := models.DeleteProduct(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete product"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Product deleted successfully"})
}
