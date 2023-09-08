package handler

import (
	"basic-auth/internal/entity"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	entity.User
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "This is the protected handler")
	if err != nil {
		log.Println("Error writing response:", err)
	}
}

func UnprotectedHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "This is the unprotected handler")
	if err != nil {
		log.Println("Error writing response:", err)
	}
}
