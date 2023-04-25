package middlewares

import (
	"net/http"

	"github.com/agomezguru/twitter/db"
)

/* This middeleware test DB to check if connections is OK */
func DBVerify(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.DBConnectionAlive() == false {
			http.Error(w, "DB connection lost", 500)
			return
		}
		// Middlewares always return the same type of input received
		next.ServeHTTP(w, r)
	}
}