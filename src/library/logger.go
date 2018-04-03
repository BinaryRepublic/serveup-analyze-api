package library

import (
	"time"
	"os"
	"fmt"
	"encoding/json"
)

type logEntry struct {
	ResponseStatusCode	int		`json:"responseStatusCode"`
	RequestMethod		string	`json:"requestMethod"`
	RequestPath			string 	`json:"requestPath"`
	ErrorType			string 	`json:"errorType"`
	ErrorMsg			string 	`json:"errorMsg"`
	LogDate				string 	`json:"logDate"`
}
type ErrorObj struct {
	Type	string	`json:"type"`
	Msg		string	`json:"msg"`
}

func Log (responseStatusCode int, requestMethod, requestPath, errorType, errorMsg string) {
	date := time.Now().UTC().Format(time.RFC3339)
	log := logEntry{
		ResponseStatusCode: responseStatusCode,
		RequestMethod: requestMethod,
		RequestPath: requestPath,
		ErrorType: errorType,
		ErrorMsg: errorMsg,
		LogDate: date,
	}
	if os.Getenv("TEST") == "" {
		jsonStr, _ := json.Marshal(log)
		fmt.Println(string(jsonStr))
	}
}

func LogError (responseStatusCode int, requestMethod, requestPath string, error ErrorObj) {
	Log(responseStatusCode, requestMethod, requestPath, error.Type, error.Msg)
}