package middleware

import "net/http"

func ContentType(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
}