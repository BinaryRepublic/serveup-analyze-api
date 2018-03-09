package controller

import (
	"net/http"
	"fmt"
	"middleware"
	"encoding/json"
	"github.com/gorilla/mux"
	"helper"
	"strings"
)

var Config = helper.ReadConfig()

type controller func(formParams map[string]string) interface{}

func HandleMessage(required []string, res http.ResponseWriter, req *http.Request, controllerFn controller) {

	// go through middleware
	middleware.HttpHeaders(res, req, "json")

	// parse query, form and json params to map[string]string
	params := make(map[string]string)
	params = parseForm(params, req)
	params = parseJSON(params, req)
	params = parseMux(params, req)

	// validate params
	missing := validateRequest(params, required, req)

	if len(missing) == 0 {
		jsonResult := controllerFn(params)
		if jsonResult != false {
			json.NewEncoder(res).Encode(jsonResult)
		}
	} else {
		for _, missingItem := range missing {
			fmt.Println("missing value: ", missingItem)
		}
	}
}
func parseForm(params map[string]string, request *http.Request) (newParams map[string]string) {
	newParams = params
	// parse query and body params
	request.ParseMultipartForm(32 << 20)
	// convert map[string][]string to map[string]string
	formParams := request.Form
	for key, value := range formParams {
		newParams[key] = value[0]
	}
	return
}
func parseJSON(params map[string]string, request *http.Request) (newParams map[string]string) {
	newParams = params
	var jsonMap map[string]string
	json.NewDecoder(request.Body).Decode(&jsonMap)
	for key, value := range jsonMap {
		newParams[key] = value
	}
	return
}
func parseMux(params map[string]string, request *http.Request) (newParams map[string]string) {
	newParams = params
	for key, value := range mux.Vars(request) {
		newParams[key] = value
	}
	return
}

func validateRequest(params map[string]string, required []string, request *http.Request) []string {
	var missing []string
	for _, requiredItem := range required {
		if params[requiredItem] == "" {
			// check if file
			_, _, err := request.FormFile(requiredItem)
			if err != nil {
				missing = append(missing, requiredItem)
			}
		}
	}
	return missing
}

func FileResponse(res http.ResponseWriter, req *http.Request, filePath string) {
	middleware.HttpHeadersContentType(res, req, "file")

	filePathSplit := strings.Split(filePath, "/")
	fileName := filePathSplit[len(filePathSplit)-1]

	res.Header().Set("Content-Disposition", "attachment; filename=\"" + fileName + "\"")
	http.ServeFile(res, req, filePath)
}