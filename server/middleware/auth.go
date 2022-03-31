package middleware

import (
	"context"
	models "go_mysql/db/models"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func UserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, e := FetchUserFromCookie(r)
		if e != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
		ctx := context.WithValue(r.Context(), "user", u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func FetchUserFromCookie(r *http.Request) (*models.User, error) {
	c, err := r.Cookie("uid")
	if err != nil {
		return nil, err
	}

	type Claims struct {
		Uid int64 `json:"uid"`
		jwt.StandardClaims
	}

	claims := &Claims{}
	secret := os.Getenv("JWT_SECRET")
	t, e := jwt.ParseWithClaims(c.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if t == nil || e != nil {
		return nil, e
	}
	var u models.User
	u.FindById(claims.Uid)
	return &u, nil
}
