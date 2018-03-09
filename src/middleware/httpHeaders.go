package middleware

import "net/http"

func HttpHeaders(res http.ResponseWriter, req *http.Request, responseType string) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,content-type")
	res.Header().Set("Access-Control-Allow-Credentials", "true")

	res.Header().Set("Cache-Control", "public, max-age=0")
	res.Header().Set("Connection", "keep-alive")

	HttpHeadersContentType(res, req, responseType)
}

func HttpHeadersContentType(res http.ResponseWriter, req *http.Request, contentType string) {
	switch contentType {
	case "json":
		res.Header().Set("Content-Type", "application/json")
	case "file":
		res.Header().Set("Content-Type", "application/octet-stream")
	}
}