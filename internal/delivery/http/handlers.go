package handler

import (
	"fmt"
	"log"
	"net/http"
)

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
