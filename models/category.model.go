package models

import (
	DB "ecommerce/database"
	"time"
)

type Category struct {
	ID          uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string     `json:"name" gorm:"not null"`
	Description string     `json:"description"`
	ParentID    *uint      `json:"parent_id"`
	Slug        string     `json:"slug" gorm:"unique"`
	Parent      *Category  `json:"parent" gorm:"foreignKey:ParentID"`
	Children    []Category `json:"children,omitempty" gorm:"foreignKey:ParentID"`

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

// GetAllCategories returns all categories with nested structure
func GetAllCategories() ([]Category, error) {
	var categories []Category
	if err := DB.Connection.Preload("Parent").Find(&categories).Error; err != nil {
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

// GetAllChildrenCategories returns all children categories for a given parent ID
func GetAllChildrenCategories(parentID uint) ([]Category, error) {
	var childrenCategories []Category
	if err := DB.Connection.Where("parent_id = ?", parentID).Preload("Parent").Find(&childrenCategories).Error; err != nil {
		return nil, err
	}

	// Recursively fetch children for each child category
	for i, child := range childrenCategories {
		grandchildren, err := GetAllChildrenCategories(child.ID)
		if err != nil {
			return nil, err
		}
		childrenCategories[i].Parent = &child
		childrenCategories = append(childrenCategories, grandchildren...)
	}

	return childrenCategories, nil
}

// CreateCategory creates a new category
func CreateCategory(newCategory *Category) error {
	if err := DB.Connection.Create(newCategory).Error; err != nil {
		return err
	}

	// If ParentID is set, update the Parent field after creation
	if newCategory.ParentID != nil && *newCategory.ParentID != 0 {
		parentCategory, err := GetCategoryByID(*newCategory.ParentID)
		if err != nil {
			return err
		}
		newCategory.Parent = parentCategory
		if err := UpdateCategory(newCategory.ID, newCategory); err != nil {
			return err
		}
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

// GetAllTopLevelCategories returns all top-level categories with nested structure
func GetAllTopLevelCategories() ([]Category, error) {
	var categories []Category
	if err := DB.Connection.Preload("Parent").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// ----------------------------------------------------------------------------------------------------------------

// GetAllCategoriesTree returns all categories with nested structure
func GetAllCategoriesTree() ([]Category, error) {
	var categories []Category
	if err := DB.Connection.Find(&categories).Error; err != nil {
		return nil, err
	}

	// Map to store categories by ID for quick access
	categoryMap := make(map[uint]*Category)
	for i := range categories {
		categoryMap[categories[i].ID] = &categories[i]
	}

	// Build the tree structure
	var resultCategories []Category
	for _, category := range categories {
		if category.ParentID == nil {
			resultCategories = append(resultCategories, buildCategoryTree(&category, categoryMap))
		}
	}

	return resultCategories, nil
}

// Recursive function to build the category tree
func buildCategoryTree(category *Category, categoryMap map[uint]*Category) Category {
	children := []Category{}

	for _, child := range categoryMap {
		if child.ParentID != nil && *child.ParentID == category.ID {
			children = append(children, buildCategoryTree(child, categoryMap))
		}
	}

	category.Children = children
	return *category
}
