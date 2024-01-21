package controllers

import (
	"ecommerce/api"
	"ecommerce/models"
	"ecommerce/utils"
	"errors"
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
		utils.LogError("GetAllReviews", "Failed to fetch reviews", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to fetch reviews")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: reviews})
}

// GetReviewByID handles GET request to fetch a review by ID
func (rc *ReviewController) GetReviewByID(c echo.Context) error {
	id := c.Param("id")
	review, err := models.GetReviewByID(id)
	if err != nil {
		utils.LogError("GetReviewByID", "Review not found", err)
		return api.Response(c, http.StatusNotFound, api.ResponseType{Error: errors.New("Review not found")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: review})
}

// CreateReview handles POST request to create a new review
func (rc *ReviewController) CreateReview(c echo.Context) error {
	review := new(models.Review)
	if err := c.Bind(review); err != nil {
		return err
	}
	if err := models.CreateReview(review); err != nil {
		utils.LogError("CreateReview", "Failed to create review", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to create review")})
	}
	return api.Response(c, http.StatusCreated, api.ResponseType{Data: review})
}

// UpdateReview handles PUT request to update a review by ID
func (rc *ReviewController) UpdateReview(c echo.Context) error {
	id := c.Param("id")
	existingReview, err := models.GetReviewByID(id)
	if err != nil {
		utils.LogError("UpdateReview", "Review not found", err)
		return api.Response(c, http.StatusNotFound, api.ResponseType{Error: errors.New("Review not found")})
	}

	updatedReview := new(models.Review)
	if err := c.Bind(updatedReview); err != nil {
		return err
	}

	if err := models.UpdateReview(existingReview, updatedReview); err != nil {
		utils.LogError("UpdateReview", "Failed to update review", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to update review")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: existingReview})
}

// DeleteReview handles DELETE request to delete a review by ID
func (rc *ReviewController) DeleteReview(c echo.Context) error {
	id := c.Param("id")
	review, err := models.GetReviewByID(id)
	if err != nil {
		utils.LogError("DeleteReview", "Review not found", err)
		return api.Response(c, http.StatusNotFound, api.ResponseType{Error: errors.New("Review not found")})
	}

	if err := models.DeleteReview(review); err != nil {
		utils.LogError("DeleteReview", "Failed to delete review", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to delete review")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Message: "Review deleted successfully"})
}
