package routers

import (
	"CAP_AdminPortal/controllers"

	"github.com/gorilla/mux"
)

// SetUserRoutes for handling user registration
func SetCategoryrRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/products/list_categories", controllers.GetCategoryData).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/products/add_category", controllers.CreateCategory).Methods("POST", "OPTIONS")
	return router
}
