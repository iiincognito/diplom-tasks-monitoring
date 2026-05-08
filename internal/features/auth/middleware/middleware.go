package auth_middleware

import (
	"net/http"

	auth_service "github.com/iiincognito/diplom-tasks-monitoring/internal/features/auth/service"
)

func NewAuthMiddleware(authService *auth_service.AuthService) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// If no password is set, skip authentication
			if !authService.IsAuthRequired() {
				next(w, r)
				return
			}

			// Get token from cookie
			var token string
			cookie, err := r.Cookie("token")
			if err == nil {
				token = cookie.Value
			}

			// Validate token
			if !authService.ValidateToken(token) {
				http.Error(w, "Authentication required", http.StatusUnauthorized)
				return
			}

			next(w, r)
		}
	}
}
