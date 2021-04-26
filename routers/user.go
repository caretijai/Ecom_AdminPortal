package routers

import (
	"CAP_AdminPortal/controllers"

	"github.com/gorilla/mux"
)

// SetUserRoutes for handling user registration
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/users/list_data", controllers.GetUserDetail).Methods("GET", "OPTIONS")
	return router
}
