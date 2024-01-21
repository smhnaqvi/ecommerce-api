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

var ProductModel = &models.Product{}

type ProductController struct{}

// GetProducts handles GET request to fetch all products
func (pc *ProductController) GetProducts(c echo.Context) error {
	products, err := ProductModel.GetAllProducts()
	if err != nil {
		utils.LogError("GetProducts", "Failed to fetch products", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to fetch products")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: products})
}

// GetProductByID handles GET request to fetch a product by ID
func (pc *ProductController) GetProductByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.LogError("GetProductByID", "Invalid product ID", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid product ID")})
	}

	product, err := ProductModel.GetProductByID(uint(id))
	if err != nil {
		utils.LogError("GetProductByID", "Failed to fetch product by ID", err)
		return api.Response(c, http.StatusNotFound, api.ResponseType{Error: errors.New("Product not found")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: product})
}

// CreateProduct handles POST request to create a new product
func (pc *ProductController) CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(&product); err != nil {
		utils.LogError("CreateProduct", "Invalid request payload", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: err})
	}

	if err := ProductModel.CreateProduct(product); err != nil {
		utils.LogError("CreateProduct", "Failed to create product", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to create product")})
	}

	return api.Response(c, http.StatusCreated, api.ResponseType{Data: product})
}

// UpdateProduct handles PUT request to update a product by ID
func (pc *ProductController) UpdateProduct(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.LogError("UpdateProduct", "Invalid product ID", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid product ID")})
	}

	product := new(models.Product)
	if err := c.Bind(&product); err != nil {
		utils.LogError("UpdateProduct", "Invalid request payload", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: err})
	}

	if err := ProductModel.UpdateProduct(uint(id), product); err != nil {
		utils.LogError("UpdateProduct", "Failed to update product", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to update product")})
	}

	return api.Response(c, http.StatusOK, api.ResponseType{Data: product})
}

// DeleteProduct handles DELETE request to delete a product by ID
func (pc *ProductController) DeleteProduct(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.LogError("DeleteProduct", "Invalid product ID", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid product ID")})
	}

	if err := ProductModel.DeleteProduct(uint(id)); err != nil {
		utils.LogError("DeleteProduct", "Failed to delete product", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to delete product")})
	}

	return api.Response(c, http.StatusOK, api.ResponseType{Message: "Product deleted successfully"})
}
