package models

import DB "ecommerce/database"

type Address struct {
	AddressID   uint   `gorm:"primaryKey;autoIncrement"`
	UserID      uint
	AddressLine1 string `gorm:"not null"`
	AddressLine2 string
	City        string `gorm:"not null"`
	State       string
	Country     string `gorm:"not null"`
	PostalCode  string `gorm:"not null"`
	// Other address-related fields as needed
}

func GetAllAddresses() ([]Address, error) {
	var addresses []Address
	if err := DB.Connection.Find(&addresses).Error; err != nil {
		return nil, err
	}
	return addresses, nil
}

func GetAddressByID(id uint) (*Address, error) {
	var address Address
	if err := DB.Connection.First(&address, id).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

func CreateAddress(address *Address) error {
	if err := DB.Connection.Create(&address).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAddress(id uint, updatedAddress *Address) error {
	var address Address
	if err := DB.Connection.First(&address, id).Error; err != nil {
		return err
	}
	if err := DB.Connection.Model(&address).Updates(updatedAddress).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAddress(id uint) error {
	var address Address
	if err := DB.Connection.First(&address, id).Error; err != nil {
		return err
	}
	if err := DB.Connection.Delete(&address).Error; err != nil {
		return err
	}
	return nil
}
