package jwt

import (
	"server/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func IssuesToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"username": user.UserName,
		"exp": time.Now().Add(time.Hour* 24).Unix(),
	})
	tokenString, err  := token.SignedString(secretKey)
	if err != nil{
		return "", err
	}
	return tokenString, nil
}

func verifyToken(tokenString string) error {
   token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      return secretKey, nil
   })
  
   if err != nil {
      return err
   }
  
   if !token.Valid {
      return err
   }
  
   return nil
}