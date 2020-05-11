package middlew

import (
	"net/http"

	"github.com/josnunezg/twittor/bd"
)

// CheckingBD middleware to check connection to BD
func CheckingBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckingConnection() == 0 {
			http.Error(w, "Connection lost with the BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
