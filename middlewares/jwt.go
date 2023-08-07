package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/oriegenbi27/go-jwt-mux/config"
	"github.com/oriegenbi27/go-jwt-mux/helper"
)

func JWTMidlleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJson(w, http.StatusUnauthorized, response)
				return
			}
		}

		// mengambil tokn value

		tokenString := c.Value

		claims := &config.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {

			return config.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				// token invalid
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJson(w, http.StatusUnauthorized, response)
				return
			case jwt.ValidationErrorExpired:
				// token expired
				response := map[string]string{"message": "Unauthorized, Token expired!"}
				helper.ResponseJson(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJson(w, http.StatusUnauthorized, response)
				return
			}
		}

		if !token.Valid {
			response := map[string]string{"message": "Unauthorized"}
			helper.ResponseJson(w, http.StatusUnauthorized, response)
			return
		}

		next.ServeHttp(w, r)
	})

}
