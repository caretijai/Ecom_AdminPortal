package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"CAP_AdminPortal/common"
	clnt "CAP_AdminPortal/service/client"
)

// GetUserDetail Handler for HTTP GET - "/users/register"
func GetUserDetail(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inside GetUserDetail handler...")
	// var dataResource UserResource
	// // Decode the incoming User json
	// err := json.NewDecoder(r.Body).Decode(&dataResource)
	// if err != nil {
	// 	common.DisplayAppError(
	// 		w,
	// 		err,
	// 		"Invalid User data",
	// 		500,
	// 	)
	// 	return
	// }
	// user := &dataResource.Data

	authorizationHeader := r.Header.Get("authorization")
	var bearerToken []string

	if authorizationHeader != "" {
		bearerToken = strings.Split(authorizationHeader, " ")
	} else {
		log.Println("token missing")
		result, err := json.Marshal([]byte("marshalling error: could not marshall the server response"))
		log.Printf("result: %v, err: %v", result, err)
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"Could Not Marshal the response",
				500,
			)
			return
		} else {
			common.DisplayAppError(
				w,
				errors.New("missing token"),
				"Authorization Token Missing",
				403,
			)
			return
		}
		// w.WriteHeader(http.StatusForbidden)
		// w.Header().Set("Content-Type", "application/json")
		// w.Write(result)
		// return
	}

	clnt.GetUserData(w, r, bearerToken[1])
}
