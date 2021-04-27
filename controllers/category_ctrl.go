package controllers

import (
	"log"
	"net/http"

	clnt "CAP_AdminPortal/service/client"
)

// func GetCategoryData Handler to get category data
func GetCategoryData(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inside GetCategoryData handler...")

	clnt.GetCategoryData(w, r)
}

// func CreateCategory Handler for creating category
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inside CreateCategoryData handler...")

	clnt.CreateCategory(w, r)
}
