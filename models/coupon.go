package models

import (
	DB "ecommerce/database"
	"time"
)

type Coupon struct {
	CouponID           uint      `gorm:"primaryKey;autoIncrement"`
	CouponCode         string    `gorm:"unique;not null"`
	DiscountPercentage float64   `gorm:"not null"`
	ExpirationDate     time.Time `gorm:"not null"`
	// Other coupon-related fields as needed
}

func GetAllCoupons() ([]Coupon, error) {
	var coupons []Coupon
	if err := DB.Connection.Find(&coupons).Error; err != nil {
		return nil, err
	}
	return coupons, nil
}

func GetCouponByID(id string) (*Coupon, error) {
	var coupon Coupon
	if err := DB.Connection.First(&coupon, id).Error; err != nil {
		return nil, err
	}
	return &coupon, nil
}

func CreateCoupon(coupon *Coupon) error {
	if err := DB.Connection.Create(&coupon).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCoupon(existingCoupon *Coupon, updatedCoupon *Coupon) error {
	if err := DB.Connection.Model(&existingCoupon).Updates(updatedCoupon).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCoupon(coupon *Coupon) error {
	if err := DB.Connection.Delete(&coupon).Error; err != nil {
		return err
	}
	return nil
}
