package responses

import "github.com/labstack/echo/v4"



type UserResponse struct {
	status int 			`json:"status"`
	message string 		`json:"message"`
	data *echo.Map 			`json:"data"`
}