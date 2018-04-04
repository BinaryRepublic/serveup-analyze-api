package helper

import (
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"os"
	"io"
	"strconv"
)

func HttpQueryParams(req *http.Request) (getParams map[string]string) {
	getParams = make(map[string]string)
	for key, value := range req.URL.Query() {
		getParams[key] = value[0]
	}
	return
}

func HttpSaveFile(res http.ResponseWriter, req *http.Request, path string) (filename string) {
	file, _, err := req.FormFile("soundfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	return f.Name()
}

func HttpGet(url string, headers map[string]string, params map[string]string) []byte {
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
	url = url + queryStr

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	// set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if  err != nil {
		fmt.Println(err)
	} else {
		if response.StatusCode == 200 {
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err.Error())
			}
			return []byte(body)
		} else {
			fmt.Println("GET " + url + " - Statuscode " + strconv.Itoa(response.StatusCode))
		}
	}
	return nil
}

func HttpPost(url string, headers map[string]string, params map[string]interface{}) []byte {
	jsonStr, _ := json.Marshal(params)
	jsonStr = []byte(jsonStr)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")
	// set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err.Error())
		}
		return []byte(body)
	} else {
		fmt.Println("POST " + url + " - Statuscode " + strconv.Itoa(response.StatusCode))
	}
	return nil
}
