package helper

import (
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

func HttpQueryParams(req *http.Request) (getParams map[string]string) {
	getParams = make(map[string]string)
	for key, value := range req.URL.Query() {
		getParams[key] = value[0]
	}
	return
}

func HttpGet(url string, params map[string]string) []byte {
	// create query
	queryStr := "?"
	for key, value := range params {
		queryStr += key + "=" + value + "&"
	}
	if queryStr != "?" {
		queryStr = strings.TrimRight(queryStr, "&")
	} else {
		queryStr = ""
	}
	response, err := http.Get(url + queryStr)
	if  err != nil {
		fmt.Println(err)
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err.Error())
		}
		return []byte(body)
	}
	return nil
}

func HttpPost(url string, params map[string]interface{}) []byte {
	jsonStr, _ := json.Marshal(params)
	jsonStr = []byte(jsonStr)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	return []byte(body)
}
