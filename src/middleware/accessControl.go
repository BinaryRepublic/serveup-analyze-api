package middleware

import "net/http"

func AccessControl(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,content-type")
	res.Header().Set("Access-Control-Allow-Credentials", "true")
}