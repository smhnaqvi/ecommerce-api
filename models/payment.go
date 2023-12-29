package models

import DB "ecommerce/database"

type Payment struct {
	PaymentID     uint    `gorm:"primaryKey;autoIncrement"`
	OrderID       uint
	PaymentMethod string  `gorm:"not null"`
	TransactionID string  `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
	PaymentStatus string  `gorm:"not null"`
	// Other payment-related fields as needed
}

func GetAllPayments() ([]Payment, error) {
	var payments []Payment
	if err := DB.Connection.Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func GetPaymentByID(id uint) (*Payment, error) {
	var payment Payment
	if err := DB.Connection.First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func CreatePayment(payment *Payment) error {
	if err := DB.Connection.Create(&payment).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePayment(id uint, updatedPayment *Payment) error {
	var payment Payment
	if err := DB.Connection.First(&payment, id).Error; err != nil {
		return err
	}

	if err := DB.Connection.Model(&payment).Updates(updatedPayment).Error; err != nil {
		return err
	}
	return nil
}

func DeletePayment(id uint) error {
	var payment Payment
	if err := DB.Connection.First(&payment, id).Error; err != nil {
		return err
	}

	if err := DB.Connection.Delete(&payment).Error; err != nil {
		return err
	}
	return nil
}
