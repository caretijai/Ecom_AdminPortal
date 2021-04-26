package controllers

import (
	"log"
	"net/http"

	clnt "CAP_AdminPortal/service/client"
)

// GetUserDetail Handler for HTTP GET - "/users/register"
func GetCategoryData(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inside GetCategoryData handler...")

	clnt.GetCategoryData(w, r)
}
