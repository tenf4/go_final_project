package api

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		pass := os.Getenv("TODO_PASSWORD")
		if len(pass) > 0 {
			var jwtToken string

			cookie, err := req.Cookie("token")
			if err == nil {
				jwtToken = cookie.Value
			}

			var valid bool
			if jwtToken != "" {
				token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method")
					}
					return []byte(pass), nil
				})

				if err == nil && token.Valid {
					claims, ok := token.Claims.(jwt.MapClaims)
					if ok {
						storedHash := claims["pass_hash"].(string)
						currentHash := fmt.Sprintf("%x", sha256.Sum256([]byte(pass)))
						if storedHash == currentHash {
							valid = true
						}
					}
				}
			}

			if !valid {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}
		}
		next(w, req)
	})
}

func signinHandler(w http.ResponseWriter, req *http.Request) {
	pass := os.Getenv("TODO_PASSWORD")
	if pass == "" {
		http.Error(w, `{"error": "Authentication not required"}`, http.StatusBadRequest)
		return
	}

	var auth struct {
		Password string `json:"password"`
	}

	data, _ := io.ReadAll(req.Body)
	json.Unmarshal(data, &auth)

	if auth.Password != pass {
		http.Error(w, `{"error": "Invalid password"}`, http.StatusUnauthorized)
		return
	}
	hashPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(pass)))
	claims := jwt.MapClaims{"pass_hash": hashPassword}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token_string, err := jwtToken.SignedString([]byte(pass))
	if err != nil {
		http.Error(w, `{"error": "token generation error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, token_string)))
}
