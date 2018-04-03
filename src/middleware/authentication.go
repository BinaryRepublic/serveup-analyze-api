package middleware

import (
	"encoding/json"
	"helper"
	"net/http"
	"strconv"
	"strings"
)

var Config = helper.ReadConfig()

type authResponse struct {
	ClientId string `json:"clientId"`
}
type errorResponse struct {
	Error map[string]string `json:"error"`
}

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		url := req.URL.Path
		noAuth := false
		// check noAuthRoutes
		for _, route := range Config.NoAuthRoutes {
			if strings.HasPrefix(url, route) {
				noAuth = true
			}
		}
		if noAuth {
			next.ServeHTTP(res, req)
		} else {
			if req.Header.Get("Access-Token") != "" {
				// prepare body
				accessParams := make(map[string]interface{})
				accessParams["accessToken"] = req.Header.Get("Access-Token")
				// send request
				response := helper.HttpPost(Config.AuthApi.Host+":"+strconv.Itoa(Config.AuthApi.Port)+"/access", nil, accessParams)
				if response != nil {
					// status 200
					// parse JSON
					var authResponseItem authResponse
					json.Unmarshal(response, &authResponseItem)
					// add accountId for authorization middleware
					req.Header.Set("accountId", authResponseItem.ClientId)
					next.ServeHTTP(res, req)
				} else {
					res.WriteHeader(400)
					var authError errorResponse
					authError.Error = map[string]string{
						"type": "ACCESS_TOKEN_INVALID",
						"msg":  "accessToken is invalid",
					}
					json.NewEncoder(res).Encode(authError)
				}
			} else {
				// no access token given
				res.WriteHeader(400)
				var authError errorResponse
				authError.Error = map[string]string{
					"type": "ACCESS_TOKEN_MISSING",
					"msg":  "Please send a valid access-token in the request header.",
				}
				json.NewEncoder(res).Encode(authError)
			}
		}
	})
}