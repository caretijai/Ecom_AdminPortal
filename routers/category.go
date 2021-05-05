package routers

import (
	"CAP_AdminPortal/controllers"

	"github.com/gorilla/mux"
)

// SetUserRoutes for handling user registration
func SetCategoryrRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/products/list_categories", controllers.GetCategoryData).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/products/add_category", controllers.CreateCategory).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/products/edit_category", controllers.EditCategory).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/products/remove_category", controllers.RemoveCategory).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/products/edit_sub_category", controllers.EditSubCategory).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/products/remove_sub_category", controllers.RemoveSubCategory).Methods("POST", "OPTIONS")
	return router
}
