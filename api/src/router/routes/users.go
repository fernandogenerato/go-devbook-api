package routes

import (
	"net/http"

	"go-devbook-api/src/controllers"
)

var userRoutes = []Route{

	{
		URI:            "/users",
		Method:         http.MethodPost,
		Function:       controllers.CreateUser,
		Authentication: false,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodGet,
		Function:       controllers.FindUserById,
		Authentication: true,
	},
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Function:       controllers.FindUsers,
		Authentication: true,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodPut,
		Function:       controllers.UpdateUser,
		Authentication: true,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodDelete,
		Function:       controllers.DeleteUser,
		Authentication: true,
	},
}
