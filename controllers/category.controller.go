package controllers

import (
	"ecommerce/models"
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
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch categories"})
	}
	return c.JSON(http.StatusOK, categories)
}

// GetCategoryByID handles GET request to fetch a category by ID
func (cc *CategoryController) GetCategoryByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid category ID"})
	}

	category, err := models.GetCategoryByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Category not found"})
	}
	return c.JSON(http.StatusOK, category)
}

// CreateCategory handles POST request to create a new category
func (cc *CategoryController) CreateCategory(c echo.Context) error {
	newCategory := new(models.Category)
	if err := c.Bind(newCategory); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if err := models.CreateCategory(newCategory); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create category"})
	}
	return c.JSON(http.StatusCreated, newCategory)
}

// UpdateCategory handles PUT request to update a category by ID
func (cc *CategoryController) UpdateCategory(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid category ID"})
	}

	updatedCategory := new(models.Category)
	if err := c.Bind(updatedCategory); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if err := models.UpdateCategory(uint(id), updatedCategory); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update category"})
	}
	return c.JSON(http.StatusOK, updatedCategory)
}

// DeleteCategory handles DELETE request to delete a category by ID
func (cc *CategoryController) DeleteCategory(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid category ID"})
	}

	if err := models.DeleteCategory(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete category"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Category deleted successfully"})
}
