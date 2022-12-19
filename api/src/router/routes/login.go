package routes

import (
	"net/http"

	"go-devbook-api/src/controllers"
)

var loginRoute = Route{
	URI:            "/login",
	Method:         http.MethodPost,
	Function:       controllers.Login,
	Authentication: false,
}
