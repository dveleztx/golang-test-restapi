package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func Welcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, `
		Welcome to the REST API for Sample Users!
		
		/users - Provide list of users
		/users/name - List of users sorted by name
		/users/name/<name> - Return record of name queried`,)
	}
}
