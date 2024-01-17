package models

import (
	DB "ecommerce/database"
)

type Product struct {
	ProductID     uint   `gorm:"primaryKey;autoIncrement"`
	Name          string `gorm:"not null"`
	Description   string
	Price         float64 `gorm:"not null"`
	StockQuantity int     `gorm:"not null"`
	CategoryID    uint
	// Other product-related fields as needed
}

// GetAllProducts returns all products
func (m *Product) GetAllProducts() ([]Product, error) {
	var products []Product
	if err := DB.Connection.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// GetProductByID returns a product by ID
func (m *Product) GetProductByID(id uint) (*Product, error) {
	var product Product
	if err := DB.Connection.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// CreateProduct creates a new product
func (m *Product) CreateProduct(product *Product) error {
	if err := DB.Connection.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

// UpdateProduct updates a product
func (m *Product) UpdateProduct(id uint, updatedProduct *Product) error {
	var product Product
	if err := DB.Connection.First(&product, id).Error; err != nil {
		return err
	}
	if err := DB.Connection.Model(&product).Updates(updatedProduct).Error; err != nil {
		return err
	}
	return nil
}

// DeleteProduct deletes a product
func (m *Product) DeleteProduct(id uint) error {
	var product Product
	if err := DB.Connection.First(&product, id).Error; err != nil {
		return err
	}
	if err := DB.Connection.Delete(&product).Error; err != nil {
		return err
	}
	return nil
}
