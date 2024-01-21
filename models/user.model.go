package models

import DB "ecommerce/database"

type User struct {
	UserID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName    string `json:"first_name" gorm:"not null"`
	LastName     string `json:"last_name" gorm:"not null"`
	PasswordHash string `json:"password" gorm:"not null"`
	Email        string `json:"email" gorm:"unique;not null"`
	Address      string `json:"address"`
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

// GetUserByEmail returns a user by email
func (m *User) GetUserByEmail(email string) (*User, error) {
	var user User
	if err := DB.Connection.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user and returns the created user ID
func (m *User) CreateUser(newUser *User) (uint, error) {
	if err := DB.Connection.Create(newUser).Error; err != nil {
		return 0, err
	}
	return newUser.UserID, nil
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
