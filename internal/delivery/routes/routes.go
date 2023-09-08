// routes.go
package routes

import (
	handler "basic-auth/internal/delivery/http"
	"basic-auth/internal/delivery/http/middleware"
	"basic-auth/internal/entity"
	"net/http"
)

func SetupRoutes(user *entity.User) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/unprotected", handler.UnprotectedHandler)
	mux.HandleFunc("/protected", middleware.BasicAuth(handler.ProtectedHandler, user))
	return mux
}
