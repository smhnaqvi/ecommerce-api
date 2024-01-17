package controllers

import (
	"ecommerce/api"
	"ecommerce/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus" // Import logrus package
)

type UserController struct {
	// Define any dependencies or services required by the controller
}

var UserModel = &models.User{}

// GetUsers is the handler for fetching all users
func (uc *UserController) GetUsers(c echo.Context) error {
	users, err := UserModel.GetAllUsers()
	if err != nil {
		log.Error("Failed to fetch users: ", err)
		response := api.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to fetch users",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := api.Response{
		Status: http.StatusOK,
		Data:   users,
	}
	return c.JSON(http.StatusOK, response)
}

// CreateUser handles POST request to create a new user
func (uc *UserController) CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		log.Error("Invalid request payload: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := user.CreateUser(user); err != nil {
		log.Error("Invalid request payload: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}
	return c.JSON(http.StatusCreated, user)
}

// GetUserByID handles GET request to fetch a user by ID
func (uc *UserController) GetUserByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Error("Invalid request payload: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	user, err := UserModel.GetUserByID(uint(id))
	if err != nil {
		log.Error("Invalid request payload: ", err)
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

// UpdateUser handles PUT request to update a user by ID
func (uc *UserController) UpdateUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Error("Invalid request payload: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		log.Error("Invalid request payload: ", err)
		return err
	}

	if err := UserModel.UpdateUser(uint(id), user); err != nil {
		log.Error("Invalid request payload: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
	}
	return c.JSON(http.StatusOK, user)
}

// DeleteUser handles DELETE request to delete a user by ID
func (uc *UserController) DeleteUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Error("Invalid request payload: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	if err := UserModel.DeleteUser(uint(id)); err != nil {
		log.Error("Invalid request payload: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}
