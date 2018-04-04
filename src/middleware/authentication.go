package middleware

import (
	"encoding/json"
	"helper"
	"net/http"
	"strconv"
	"strings"
	"library"
)

var Config = helper.ReadConfig()

type authResponse struct {
	ClientId string `json:"clientId"`
}
type errorResponse struct {
	Error library.ErrorObj `json:"error"`
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
					authError.Error = library.ErrorObj{
						Type: "ACCESS_TOKEN_INVALID",
						Msg:  "accessToken is invalid",
					}
					library.LogError(400, req.Method, req.RequestURI, authError.Error)
					json.NewEncoder(res).Encode(authError)
				}
			} else {
				// no access token given
				res.WriteHeader(400)
				var authError errorResponse
				authError.Error = library.ErrorObj{
					Type: "ACCESS_TOKEN_MISSING",
					Msg:  "Please send a valid access-token in the request header.",
				}
				library.LogError(400, req.Method, req.RequestURI, authError.Error)
				json.NewEncoder(res).Encode(authError)
			}
		}
	})
}
