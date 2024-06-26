package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pass := os.Getenv("TODO_PASSWORD")
		if len(pass) > 0 {
			hash := sha256.Sum256([]byte(os.Getenv("TODO_PASSWORD")))
			hashString := hex.EncodeToString(hash[:])

			var cookieToken string
			cookie, err := r.Cookie("token")
			if err == nil {
				cookieToken = cookie.Value
			}

			jwtToken, err := jwt.Parse(cookieToken, func(t *jwt.Token) (interface{}, error) {
				return []byte(pass), nil
			})
			if err != nil {
				log.Printf("неудалось преобразовать токен: %s\n", err)
				return
			}

			res, ok := jwtToken.Claims.(jwt.MapClaims)
			if !ok {
				log.Println("неудалось привести к типу jwt.MapCalims")
				return
			}

			hashRaw := res["hash"]

			tokenHash, ok := hashRaw.(string)
			if !ok {
				log.Println("неудалось привести к типу string")
				return
			}

			if tokenHash != hashString {
				log.Println("хэши не совпали")
				http.Error(w, "Authentication required", http.StatusUnauthorized)
				return
			}

			log.Println("аутентификация пройдена")
		}
		next(w, r)
	}
}
