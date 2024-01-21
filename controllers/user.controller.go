package controllers

import (
	"ecommerce/api"
	"ecommerce/models"
	"ecommerce/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	// Import logrus package
)

type UserController struct {
	// Define any dependencies or services required by the controller
}

// create a instance of user model for use in all controller funcs
var UserModel = &models.User{}

// GetUsers is the handler for fetching all users
func (uc *UserController) GetUsers(c echo.Context) error {
	users, err := UserModel.GetAllUsers()
	if err != nil {
		utils.LogError("GetUsers", "Failed to fetch users", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to fetch users")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: users})
}

// CreateUser handles POST request to create a new user
func (uc *UserController) CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		utils.LogError("CreateUser", "Invalid user ID", err)

		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: err})
	}

	if _, err := user.CreateUser(user); err != nil {
		utils.LogError("CreateUser", "Failed to create user", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{Error: errors.New("Failed to create user")})
	}

	return api.Response(c, http.StatusCreated, api.ResponseType{Data: user})
}

// GetUserByID handles GET request to fetch a user by ID
func (uc *UserController) GetUserByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.LogError("GetUserByID", "Invalid user ID", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{Error: errors.New("Invalid user ID")})
	}

	user, err := UserModel.GetUserByID(uint(id))
	if err != nil {
		utils.LogError("GetUserByID", "Failed to fetch user by ID", err)
		return api.Response(c, http.StatusNotFound, api.ResponseType{Error: errors.New("Invalid user ID")})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: user})
}

// UpdateUser handles PUT request to update a user by ID
func (uc *UserController) UpdateUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.LogError("UpdateUser", "Invalid user ID", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{
			Error: errors.New("Invalid ID"),
		})
	}

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		utils.LogError("UpdateUser", "Invalid user ID", err)
		return err
	}

	if err := UserModel.UpdateUser(uint(id), user); err != nil {
		utils.LogError("UpdateUser", "Failed to update user", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{
			Error: errors.New("Failed to update user"),
		})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{Data: user})
}

// DeleteUser handles DELETE request to delete a user by ID
func (uc *UserController) DeleteUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.LogError("DeleteUser", "Invalid user ID", err)
		return api.Response(c, http.StatusBadRequest, api.ResponseType{
			Error: errors.New("Invalid ID"),
		})
	}

	if err := UserModel.DeleteUser(uint(id)); err != nil {
		utils.LogError("DeleteUser", "Failed to delete user", err)
		return api.Response(c, http.StatusInternalServerError, api.ResponseType{
			Error: errors.New("Failed to delete user"),
		})
	}
	return api.Response(c, http.StatusOK, api.ResponseType{
		Message: "User deleted successfully",
	})
}
