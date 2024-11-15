package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)
var secretKey = []byte("your-secret-key")

func VerifyJWT(next http.Handler) http.Handler{
	return http.HandlerFunc( func(w http.ResponseWriter, r* http.Request){
		cookie, err := r.Cookie("token")
		if err != nil{
			http.Error(w, "Unauthorized: No cookie found", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		
		
		if !token.Valid  || err != nil{
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}
		

		next.ServeHTTP(w, r)
	})
}