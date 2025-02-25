package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/chepaqq/image-service/pkg/logger"
	"github.com/golang-jwt/jwt"
)

type contextKey string

// UserIDKey is a context key to store user id
const UserIDKey contextKey = "user_id"

// AuthMiddleware validates user access
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "missing authorization token", http.StatusUnauthorized)
			logger.Error("missing authorization token")
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			http.Error(w, "invalid authorization token", http.StatusUnauthorized)
			logger.Error("invalid authorization token:", err)
			return
		}

		if !token.Valid {
			http.Error(w, "invalid authorization token", http.StatusUnauthorized)
			logger.Error("invalid authorization token: token not valid")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "invalid authorization token", http.StatusUnauthorized)
			logger.Error("invalid authorization token: invalid claims")
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			http.Error(w, "invalid authorization token", http.StatusUnauthorized)
			logger.Error("invalid authorization token: user_id not found in claims")
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AccessControlMiddleware handles access control and  CORS middleware
func AccessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
