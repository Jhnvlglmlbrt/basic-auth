package main

import (
	routes "basic-auth/internal/delivery/routes"
	entity "basic-auth/internal/entity"
	path "basic-auth/path"
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

	log.Print(path.GetFilePath())

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
	err := srv.ListenAndServeTLS(path.GetFilePath()+"/certs/localhost.pem", path.GetFilePath()+"/certs/localhost-key.pem")
	log.Fatal(err)
}
