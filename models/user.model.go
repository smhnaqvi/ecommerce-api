package models

import DB "ecommerce/database"

type User struct {
	UserID       uint   `gorm:"primaryKey;autoIncrement"`
	Username     string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	Address      string
	// Other user-related fields as needed
}

// GetAllUsers returns all users
func (m *User) GetAllUsers() ([]User, error) {
	var users []User
	if err := DB.Connection.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID returns a user by ID
func (m *User) GetUserByID(id uint) (*User, error) {
	var user User
	if err := DB.Connection.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user
func (m *User) CreateUser(newUser *User) error {
	if err := DB.Connection.Create(newUser).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser updates a user
func (m *User) UpdateUser(id uint, updatedUser *User) error {
	var user User
	if err := DB.Connection.First(&user, id).Error; err != nil {
		return err
	}
	if err := DB.Connection.Model(&user).Updates(updatedUser).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user
func (m *User) DeleteUser(id uint) error {
	var user User
	if err := DB.Connection.First(&user, id).Error; err != nil {
		return err
	}
	if err := DB.Connection.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
