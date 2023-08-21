package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/chepaqq/jungle-task/internal/service"
	"github.com/golang-jwt/jwt"
)

type contextKey string

const userCtx contextKey = "user_id"

// UserMiddleware represents middlewares for user-related operations
type UserMiddleware struct {
	userMiddleware service.UserService
}

// NewUserMiddleware creates and returns a new UserMiddleware object
func NewUserMiddleware(userService service.UserService) *UserMiddleware {
	return &UserMiddleware{userMiddleware: userService}
}

// AccessMiddleware validates user access
func (m *UserMiddleware) AccessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			// TODO: get secret key from env file
			return []byte("secret"), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
