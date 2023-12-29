package models

import (
	"time"

	DB "ecommerce/database"
)

type Order struct {
	OrderID     uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint
	OrderDate   time.Time `gorm:"default:'2006-01-02 15:04:05'"`
	TotalAmount float64   `gorm:"not null"`
	Status      string    `gorm:"not null"`
	// Other order-related fields as needed
}

func GetAllOrders() ([]Order, error) {
	var orders []Order
	if err := DB.Connection.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrderByID(id uint) (*Order, error) {
	var order Order
	if err := DB.Connection.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func CreateOrder(order *Order) error {
	if err := DB.Connection.Create(&order).Error; err != nil {
		return err
	}
	return nil
}

func UpdateOrder(id uint, updatedOrder *Order) error {
	var order Order
	if err := DB.Connection.First(&order, id).Error; err != nil {
		return err
	}

	if err := DB.Connection.Model(&order).Updates(updatedOrder).Error; err != nil {
		return err
	}
	return nil
}

func DeleteOrder(id uint) error {
	var order Order
	if err := DB.Connection.First(&order, id).Error; err != nil {
		return err
	}

	if err := DB.Connection.Delete(&order).Error; err != nil {
		return err
	}
	return nil
}
