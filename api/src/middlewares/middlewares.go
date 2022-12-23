package middlewares

import (
	"log"
	"net/http"

	"go-devbook-api/src/auth"
	"go-devbook-api/src/response"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func DoAuth(next http.HandlerFunc) http.HandlerFunc {
	log.Println("DoAuth.....")
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			response.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}

}
