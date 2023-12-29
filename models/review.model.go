package models

import (
	DB "ecommerce/database"
	"time"
)

type Review struct {
	ReviewID   uint      `gorm:"primaryKey;autoIncrement"`
	ProductID  uint
	UserID     uint
	ReviewText string    `gorm:"type:text"`
	Rating     float64   `gorm:"not null"`
	ReviewDate time.Time `gorm:"default:'2006-01-02 15:04:05'"`
	// Other review-related fields as needed
}

// GetAllReviews returns all reviews
func GetAllReviews() ([]Review, error) {
	var reviews []Review
	if err := DB.Connection.Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

// GetReviewByID returns a review by ID
func GetReviewByID(id string) (*Review, error) {
	var review Review
	if err := DB.Connection.First(&review, id).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

// CreateReview creates a new review
func CreateReview(review *Review) error {
	if err := DB.Connection.Create(&review).Error; err != nil {
		return err
	}
	return nil
}

// UpdateReview updates a review
func UpdateReview(existingReview *Review, updatedReview *Review) error {
	if err := DB.Connection.Model(&existingReview).Updates(updatedReview).Error; err != nil {
		return err
	}
	return nil
}

// DeleteReview deletes a review
func DeleteReview(review *Review) error {
	if err := DB.Connection.Delete(&review).Error; err != nil {
		return err
	}
	return nil
}
