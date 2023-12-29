package models

import (
	DB "ecommerce/database"
)

type OrderDetail struct {
	OrderDetailID uint    `gorm:"primaryKey;autoIncrement"`
	OrderID       uint
	ProductID     uint
	Quantity      int     `gorm:"not null"`
	Price         float64 `gorm:"not null"`
	// Other order detail-related fields as needed
}

// GetAllOrderDetails returns all order details
func GetAllOrderDetails() ([]OrderDetail, error) {
	var orderDetails []OrderDetail
	if err := DB.Connection.Find(&orderDetails).Error; err != nil {
		return nil, err
	}
	return orderDetails, nil
}

// GetOrderDetailByID returns an order detail by ID
func GetOrderDetailByID(id string) (*OrderDetail, error) {
	var orderDetail OrderDetail
	if err := DB.Connection.First(&orderDetail, id).Error; err != nil {
		return nil, err
	}
	return &orderDetail, nil
}

// CreateOrderDetail creates a new order detail
func CreateOrderDetail(orderDetail *OrderDetail) error {
	if err := DB.Connection.Create(&orderDetail).Error; err != nil {
		return err
	}
	return nil
}

// UpdateOrderDetail updates an order detail
func UpdateOrderDetail(existingOrderDetail *OrderDetail, updatedOrderDetail *OrderDetail) error {
	if err := DB.Connection.Model(&existingOrderDetail).Updates(updatedOrderDetail).Error; err != nil {
		return err
	}
	return nil
}

// DeleteOrderDetail deletes an order detail
func DeleteOrderDetail(orderDetail *OrderDetail) error {
	if err := DB.Connection.Delete(&orderDetail).Error; err != nil {
		return err
	}
	return nil
}
