package controllers

import (
	"ecommerce/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ReviewController represents the controller for Review operations
type ReviewController struct {
	// Define any dependencies or services required by the controller
}

// GetAllReviews handles GET request to fetch all reviews
func (rc *ReviewController) GetAllReviews(c echo.Context) error {
	reviews, err := models.GetAllReviews()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch reviews"})
	}
	return c.JSON(http.StatusOK, reviews)
}

// GetReviewByID handles GET request to fetch a review by ID
func (rc *ReviewController) GetReviewByID(c echo.Context) error {
	id := c.Param("id")
	review, err := models.GetReviewByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Review not found"})
	}
	return c.JSON(http.StatusOK, review)
}

// CreateReview handles POST request to create a new review
func (rc *ReviewController) CreateReview(c echo.Context) error {
	review := new(models.Review)
	if err := c.Bind(review); err != nil {
		return err
	}
	if err := models.CreateReview(review); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create review"})
	}
	return c.JSON(http.StatusCreated, review)
}

// UpdateReview handles PUT request to update a review by ID
func (rc *ReviewController) UpdateReview(c echo.Context) error {
	id := c.Param("id")
	existingReview, err := models.GetReviewByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Review not found"})
	}

	updatedReview := new(models.Review)
	if err := c.Bind(updatedReview); err != nil {
		return err
	}

	if err := models.UpdateReview(existingReview, updatedReview); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update review"})
	}
	return c.JSON(http.StatusOK, existingReview)
}

// DeleteReview handles DELETE request to delete a review by ID
func (rc *ReviewController) DeleteReview(c echo.Context) error {
	id := c.Param("id")
	review, err := models.GetReviewByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Review not found"})
	}

	if err := models.DeleteReview(review); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete review"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Review deleted successfully"})
}
