package middleware

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte

type contextKey string

const UserKey contextKey = "user"

func GetUsernameFromContext(ctx context.Context) (string, error) {
	username, ok := ctx.Value(UserKey).(string)
	if !ok {
		return "", fmt.Errorf("no username found in context")
	}
	return username, nil
}

// Initialize secret key either from environment variable or generate a new one
func init() {
	secretKeyStr := os.Getenv("JWT_SECRET_KEY")
	if secretKeyStr == "" {
		secretKey = []byte(generateSecretKey())
		log.Println("Generated new secret key.")
	} else {
		secretKey = []byte(secretKeyStr)
		log.Println("Loaded secret key from environment variable.")
	}
}

// Generate a new secret key if it's not provided in the environment
func generateSecretKey() string {
	secret := make([]byte, 32) // 32 bytes = 256 bits
	_, err := rand.Read(secret)
	if err != nil {
		log.Fatal("Error generating secret key: ", err)
	}
	return base64.StdEncoding.EncodeToString(secret)
}

// IssuesToken creates a new JWT token for the user
func IssuesToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"username": user.UserName,
		"exp":      time.Now().Add(24 * time.Hour).Unix(), // 24 hours expiration
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// VerifyJWT middleware checks the validity of the JWT token in the request
func VerifyJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			log.Println("Error retrieving token from cookie:", err)
			http.Error(w, "Unauthorized: No cookie found", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Println("Error: unexpected signing method")
				http.Error(w, "Unauthorized: Invalid token method", http.StatusUnauthorized)
				return nil, err
			}
			return secretKey, nil
		})

		// If the token is invalid or there's an error parsing it
		if err != nil || !token.Valid {
			log.Println("Error parsing token or invalid token:", err)
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized: Invalid token claims", http.StatusUnauthorized)
			return
		}

		username := claims["username"].(string)
		log.Println("Username from token:", username)

		ctx := context.WithValue(r.Context(), UserKey, username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
