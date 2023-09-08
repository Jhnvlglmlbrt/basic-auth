package middleware

import (
	"basic-auth/internal/entity"
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

func BasicAuth(next http.HandlerFunc, user *entity.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(user.Username))
			expectedPasswordHash := sha256.Sum256([]byte(user.Password))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted", "charset=utf-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
