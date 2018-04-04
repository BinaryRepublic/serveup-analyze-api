package controller

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"helper"
	"strings"
	"os"
	"library"
)

var Config = helper.ReadConfig()

type controller func(formParams map[string]string) ControllerResult

type ControllerResult struct {
	Success		interface{}
	Error		library.ErrorObj
}

func HandleMessage(required []string, res http.ResponseWriter, req *http.Request, controllerFn controller) {
	// parse query, form and json params to map[string]string
	params := make(map[string]string)
	params = parseForm(params, req)
	params = parseJSON(params, req)
	params = parseMux(params, req)

	// validate params
	missing := validateRequest(params, required, req)

	if len(missing) == 0 {
		controllerRes := controllerFn(params)
		if controllerRes.Error.Msg == "" {
			// 200 SUCCESS
			library.Log(200, req.Method, req.RequestURI, "", "")
			json.NewEncoder(res).Encode(controllerRes.Success)
		} else {
			if res.Header().Get("Content-Type") == "application/json" {
				// 500 INTERNAL
				controllerRes.Error.Type = "INTERNAL_SERVER_ERROR"
				library.LogError(500, req.Method, req.RequestURI, controllerRes.Error)
				res.WriteHeader(500)
				json.NewEncoder(res).Encode(controllerRes.Error.Msg)
			}
		}
	} else {
		missingStr := ""
		for _, missingItem := range missing {
			missingStr += missingItem + ", "
		}
		missingStr = missingStr[:len(missingStr)-2]

		fmt.Println("missing value: ", missingStr)
		// 500 INTERNAL
		errorObj := library.ErrorObj{
			Type: "BAD_REQUEST",
			Msg: "missing value: " + missingStr,
		}
		library.LogError(400, req.Method, req.RequestURI, errorObj)
		res.WriteHeader(400)
		json.NewEncoder(res).Encode(errorObj)
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

func FileResponse(res http.ResponseWriter, req *http.Request, filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		res.Header().Del("Content-Type")
		res.Header().Set("Content-Type", "application/octet-stream")
		filePathSplit := strings.Split(filePath, "/")
		fileName := filePathSplit[len(filePathSplit)-1]

		res.Header().Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
		http.ServeFile(res, req, filePath)

		library.Log(200, req.Method, req.RequestURI, "", "")

		return true
	}
	return false
}