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

// func EditCategory Handler for editing category
func EditCategory(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inside edit category data handler...")

	clnt.EditCategory(w, r)
}

// func EditCategory Handler for editing category
func RemoveCategory(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inside remove category data handler...")

	clnt.RemoveCategory(w, r)
}

// func EditSubCategory for editing sub-categories
func EditSubCategory(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inside edit sub-category for data handler...")

	clnt.EditSubCategory(w, r)
}

// func RemoveSubCategory for removing sub-category
func RemoveSubCategory(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inside RemoveSubCategory sub-category for data handler...")

	clnt.RemoveSubCategory(w, r)
}
