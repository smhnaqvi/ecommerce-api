// session.go
package models

import (
	"ecommerce/database"
	"time"
)

type Session struct {
	SessionID    uint      `gorm:"primaryKey;autoIncrement"`
	UserID       uint      `gorm:"not null"`
	AccessToken  string    `gorm:"not null"`
	RefreshToken string    `gorm:"not null"`
	ExpiresAt    time.Time `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

// GetAllSessions returns all sessions
func (m *Session) GetAllSessions() ([]Session, error) {
	var sessions []Session
	if err := database.Connection.Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

// GetSessionByID returns a session by ID
func (m *Session) GetSessionByID(id uint) (*Session, error) {
	var session Session
	if err := database.Connection.First(&session, id).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

// GetSessionByUserID retrieves a session by user ID
func (m *Session) GetSessionByUserID(userID uint) (*Session, error) {
	var session Session
	if err := database.Connection.Where("user_id = ?", userID).First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

// CreateSession creates a new session
func (m *Session) CreateSession(newSession *Session) error {
	if err := database.Connection.Create(newSession).Error; err != nil {
		return err
	}
	return nil
}

// UpdateSession updates a session
func (m *Session) UpdateSession(session *Session) error {
	if err := database.Connection.Save(session).Error; err != nil {
		return err
	}
	return nil
}

// DeleteSessionByID deletes a session by ID
func (m *Session) DeleteSessionByID(id uint) error {
	var session Session
	if err := database.Connection.First(&session, id).Error; err != nil {
		return err
	}
	if err := database.Connection.Delete(&session).Error; err != nil {
		return err
	}
	return nil
}
