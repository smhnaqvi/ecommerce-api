package models

import DB "ecommerce/database"

type Category struct {
	CategoryID  uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"not null"`
	Description string
	// Other category-related fields as needed
}

// GetAllCategories returns all categories
func GetAllCategories() ([]Category, error) {
	var categories []Category
	if err := DB.Connection.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// GetCategoryByID returns a category by ID
func GetCategoryByID(id uint) (*Category, error) {
	var category Category
	if err := DB.Connection.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// CreateCategory creates a new category
func CreateCategory(newCategory *Category) error {
	if err := DB.Connection.Create(&newCategory).Error; err != nil {
		return err
	}
	return nil
}

// UpdateCategory updates a category
func UpdateCategory(id uint, updatedCategory *Category) error {
	var category Category
	if err := DB.Connection.First(&category, id).Error; err != nil {
		return err
	}
	if err := DB.Connection.Model(&category).Updates(updatedCategory).Error; err != nil {
		return err
	}
	return nil
}

// DeleteCategory deletes a category
func DeleteCategory(id uint) error {
	var category Category
	if err := DB.Connection.First(&category, id).Error; err != nil {
		return err
	}
	if err := DB.Connection.Delete(&category).Error; err != nil {
		return err
	}
	return nil
}
