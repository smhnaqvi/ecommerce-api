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

// CategoryController represents the controller for Category operations
type CategoryController struct {
	// Define any dependencies or services required by the controller
}

// GetAllCategories handles GET request to fetch all categories
func (cc *CategoryController) GetAllCategories(c echo.Context) error {
	categories, err := models.GetAllCategories()
	if err != nil {
		utils.LogError("GetAllCategories", "Failed to fetch categories", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to fetch categories")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: categories})
}

// GetCategoryByID handles GET request to fetch a category by ID
func (cc *CategoryController) GetCategoryByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.LogError("GetCategoryByID", "Invalid category ID", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid category ID")})
	}

	category, err := models.GetCategoryByID(uint(id))
	if err != nil {
		utils.LogError("GetCategoryByID", "Failed to fetch category by ID", err)
		return api.Response(c, http.StatusNotFound, api.ResponseType{Error: errors.New("Category not found")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: category})
}

// CreateCategory handles POST request to create a new category
func (cc *CategoryController) CreateCategory(c echo.Context) error {
	newCategory := new(models.Category)
	if err := c.Bind(newCategory); err != nil {
		utils.LogError("CreateCategory", "Invalid request payload", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: err})
	}

	if err := models.CreateCategory(newCategory); err != nil {
		utils.LogError("CreateCategory", "Failed to create category", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to create category")})
	}

	return api.Response(c, http.StatusCreated, api.ResponseType{Data: newCategory})
}

// UpdateCategory handles PUT request to update a category by ID
func (cc *CategoryController) UpdateCategory(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.LogError("UpdateCategory", "Invalid category ID", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid category ID")})
	}

	updatedCategory := new(models.Category)
	if err := c.Bind(updatedCategory); err != nil {
		utils.LogError("UpdateCategory", "Invalid request payload", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: err})
	}

	if err := models.UpdateCategory(uint(id), updatedCategory); err != nil {
		utils.LogError("UpdateCategory", "Failed to update category", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to update category")})
	}

	return api.Response(c, http.StatusOK, api.ResponseType{Data: updatedCategory})
}

// DeleteCategory handles DELETE request to delete a category by ID
func (cc *CategoryController) DeleteCategory(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.LogError("DeleteCategory", "Invalid category ID", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid category ID")})
	}

	if err := models.DeleteCategory(uint(id)); err != nil {
		utils.LogError("DeleteCategory", "Failed to delete category", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to delete category")})
	}

	return api.Response(c, http.StatusOK, api.ResponseType{Message: "Category deleted successfully"})
}
