package controllers

import (
	"ecommerce/api"
	"ecommerce/models"
	"ecommerce/utils"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
}

// LoginRequest struct with email and password
type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

// Login handles user authentication
func (ac *AuthController) Login(c echo.Context) error {
	loginRequest := new(LoginRequest)
	if err := c.Bind(loginRequest); err != nil {
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid request payload")})
	}

	userModel := &models.User{}
	user, err := userModel.GetUserByEmail(loginRequest.Email)
	if err != nil {
		utils.LogError("AuthController.Login", "Invalid email or password", err)
		return api.Response(c, http.StatusUnauthorized, api.ResponseType{Error: errors.New("Invalid email or password")})
	}

	// Validate the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginRequest.Password)); err != nil {
		utils.LogError("AuthController.Login", "Invalid email or password", err)
		return api.Response(c, http.StatusUnauthorized, api.ResponseType{Error: errors.New("Invalid email or password")})
	}

	// Password is correct, generate tokens and save session
	tokenPair, err := utils.GenerateTokenPair(user.UserID)
	if err != nil {
		utils.LogError("AuthController.Login", "Failed to generate token pair", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to generate token pair")})
	}

	// Save session in the database
	sessionModel := &models.Session{
		UserID:       user.UserID,
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		ExpiresAt:    time.Now().Add(15 * time.Minute), // Set expiration time as needed
	}
	if err := sessionModel.CreateSession(sessionModel); err != nil {
		utils.LogError("AuthController.Login", "Failed to save session", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to save session")})
	}

	// Return access token in the response
	return api.Response(c, http.StatusCreated, api.ResponseType{
		Data: map[string]interface{}{
			"user":          user,
			"access_token":  tokenPair.AccessToken,
			"refresh_token": tokenPair.RefreshToken,
		},
	})
}

type RegisterRequest struct {
	FirstName string `json:"first_name" form:"first_name" validate:"required"`
	LastName  string `json:"last_name" form:"last_name" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required,email"`
	Address   string `json:"address" form:"address"`
	// Other user-related fields as needed
}

// Register creates a new user account
func (ac *AuthController) Register(c echo.Context) error {
	registerRequest := new(RegisterRequest)
	if err := c.Bind(registerRequest); err != nil {
		utils.LogError("AuthController.Register", "Invalid request payload", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid request payload")})
	}

	// Check if a user with the same email already exists
	userModel := &models.User{}
	existingUser, err := userModel.GetUserByEmail(registerRequest.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.LogError("AuthController.Register", "Failed to check existing user", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to check existing user")})
	}

	if existingUser != nil {
		// User with the same email already exists
		return api.Response(c, http.StatusConflict, api.ResponseType{Error: errors.New("User with the same email already exists")})
	}

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.LogError("AuthController.Register", "Failed to hash password", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to hash password")})
	}

	// Create a new user
	newUser := &models.User{
		FirstName:    registerRequest.FirstName,
		LastName:     registerRequest.LastName,
		PasswordHash: string(hashedPassword),
		Email:        registerRequest.Email,
		Address:      registerRequest.Address,
		// Other fields as needed
	}

	if _, err := userModel.CreateUser(newUser); err != nil {
		utils.LogError("AuthController.Register", "Failed to create user", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to create user")})
	}

	// Generate tokens for the newly registered user
	tokenPair, err := utils.GenerateTokenPair(newUser.UserID)
	if err != nil {
		utils.LogError("AuthController.Register", "Failed to generate token pair", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to generate token pair")})
	}

	// Save session in the database
	sessionModel := &models.Session{
		UserID:       newUser.UserID,
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		ExpiresAt:    time.Now().Add(15 * time.Minute), // Set expiration time as needed
	}
	if err := sessionModel.CreateSession(sessionModel); err != nil {
		utils.LogError("AuthController.Login", "Failed to save session", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to save session")})
	}

	// User successfully registered
	return api.Response(c, http.StatusCreated, api.ResponseType{
		Data: map[string]string{
			"access_token":  tokenPair.AccessToken,
			"refresh_token": tokenPair.RefreshToken,
		},
	})
}
