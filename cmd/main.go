package main

import (
	routes "basic-auth/internal/delivery/routes"
	entity "basic-auth/internal/entity"
	"basic-auth/util"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	app := &entity.User{
		Username: os.Getenv("AUTH_USERNAME"),
		Password: os.Getenv("AUTH_PASSWORD"),
	}

	if app.Username == "" {
		log.Fatal("basic auth username must be provided")
	}

	if app.Password == "" {
		log.Fatal("basic auth password must be provided")
	}

	mux := routes.SetupRoutes(app)

	srv := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("starting server on %s", srv.Addr)
	err := srv.ListenAndServeTLS(util.GetFilePath()+"/certs/localhost.pem", util.GetFilePath()+"/certs/localhost-key.pem")
	log.Fatal(err)
}
