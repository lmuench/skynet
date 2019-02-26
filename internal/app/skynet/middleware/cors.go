package middleware

import "net/http"

// CORS middleware sets "Access-Control-Allow-Origin *" in header
func CORS(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	next(w, r)
}
