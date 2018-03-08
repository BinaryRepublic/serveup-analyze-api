package controller

import (
	"net/http"
	"fmt"
	"middleware"
	"encoding/json"
)

type controller func(formParams map[string]string) interface{}

func HandleMessage(required []string, res http.ResponseWriter, req *http.Request, controllerFn controller) {

	// go through middleware
	middleware.AccessControl(res, req)
	middleware.ContentType(res, req)

	// parse query, form and json params to map[string]string
	params := make(map[string]string)
	params = parseForm(params, req)
	params = parseJSON(params, req)

	// validate params
	missing := validateRequest(params, required)

	if len(missing) == 0 {
		json.NewEncoder(res).Encode(controllerFn(params))
	} else {
		for _, missingItem := range missing {
			fmt.Println("missing value: ", missingItem)
		}
	}
}
func parseForm(params map[string]string, request *http.Request) (newParams map[string]string) {
	newParams = params
	// parse query and body params
	request.ParseForm()
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

func validateRequest(params map[string]string, required []string) []string {
	var missing []string
	for _, requiredItem := range required {
		if params[requiredItem] == "" {
			missing = append(missing, requiredItem)
		}
	}
	return missing
}