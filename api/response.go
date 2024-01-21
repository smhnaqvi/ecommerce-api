package api

import "github.com/labstack/echo/v4"

type ResponseType struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   error       `json:"error,omitempty"`
}

// NewResponse creates a new Response object with the given status, message, and data
func Response(c echo.Context, status int, params ResponseType) error {
	return c.JSON(status, params)
}
