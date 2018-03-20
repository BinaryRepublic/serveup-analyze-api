package middleware

import "net/http"

func HttpHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		res.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,content-type")
		res.Header().Set("Access-Control-Allow-Credentials", "true")

		res.Header().Set("Cache-Control", "public, max-age=0")
		res.Header().Set("Connection", "keep-alive")

		res.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(res, req)
	})
}