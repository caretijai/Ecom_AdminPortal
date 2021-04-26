package main

import (
	"log"
	"net/http"
	"os"

	"CAP_AdminPortal/common"
	"CAP_AdminPortal/routers"

	"github.com/codegangsta/negroni"
	"github.com/joho/godotenv"

	srv "CAP_AdminPortal/service/server"
)

//Entry point of the program
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("cap_admin_portal db host: ", os.Getenv("POSTGRES_HOST"))

	// Running RPC server
	s := srv.Server{}
	go s.InitServer()

	// Get the mux router object
	router := routers.InitRoutes()
	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("AppServer Listening...")
	server.ListenAndServe()
}
