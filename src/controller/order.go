package controller

import (
	"net/http"
	"helper"
	"strconv"
	"encoding/json"
	"os"
)

var orderApiUrl = Config.OrderApi.Host + ":" + strconv.Itoa(Config.OrderApi.Port)

type drinkItem struct {
	Id		string			`json:"id"`
	Name	string			`json:"name"`
	Size	int				`json:"size"`
	Nb		int				`json:"nb"`
}
type serviceItem struct {
	Id		string			`json:"id"`
	Name	string			`json:"name"`
}
type order struct {
	Id            string      	`json:"id"`
	Timestamp     string      	`json:"timestamp"`
	Drinks        []drinkItem 	`json:"drinks"`
	Services      []serviceItem	`json:"services"`
	VoiceDeviceId string      	`json:"voicedeviceId"`
	RestaurantId  string      	`json:"restaurantId"`
	Status        int         	`json:"status"`
	SoundFilePath string      	`json:"soundfile-path"`
}

func GetOrderById(res http.ResponseWriter, req *http.Request) {
	HandleMessage([]string{"id"}, res, req, func(params map[string]string) ControllerResult {
		var result ControllerResult
		// prepare headers
		headers := make(map[string]string)
		headers["Access-Token"] = req.Header.Get("Access-Token")
		// send request
		res := helper.HttpGet(orderApiUrl + "/order", headers, params)
		var orderItem order
		json.Unmarshal(res, &orderItem)
		if orderItem.Id != "" {
			soundFilePath := Config.SoundFiles.Path + orderItem.Id + "." + Config.SoundFiles.Type
			if _, err := os.Stat(soundFilePath); err == nil {
				orderItem.SoundFilePath = helper.SoundFileUrl(req, orderItem.Id)
			}
			result.Success = orderItem
			return result
		}
		result.Error.Msg = "Invalid orderId: " + params["id"]
		return result
	})
}
func GetOrderByRestaurant(res http.ResponseWriter, req *http.Request) {
	HandleMessage([]string{"restaurant-id"}, res, req, func(params map[string]string) ControllerResult {
		var result ControllerResult
		// prepare headers
		headers := make(map[string]string)
		headers["Access-Token"] = req.Header.Get("Access-Token")
		// send request
		res := helper.HttpGet(orderApiUrl + "/order/restaurant", headers, params)
		if res != nil {
			var orders []order
			json.Unmarshal(res, &orders)
			for index := range orders {
				soundFilePath := Config.SoundFiles.Path + orders[index].Id + "." + Config.SoundFiles.Type
				order := &orders[index]
				if _, err := os.Stat(soundFilePath); err == nil {
					order.SoundFilePath = helper.SoundFileUrl(req, orders[index].Id)
				}
			}
			result.Success = orders
			return result
		} else {
			result.Error.Msg = "Invalid restaurantId: " + params["restaurant-id"]
			return result
		}
	})
}

func PostOrder(res http.ResponseWriter, req *http.Request) {
	HandleMessage([]string{"order-id", "soundfile"}, res, req, func(params map[string]string) ControllerResult {
		var result ControllerResult
		// prepare query
		getQuery := make(map[string]string)
		getQuery["id"] = params["order-id"]
		// prepare headers
		headers := make(map[string]string)
		headers["Access-Token"] = req.Header.Get("Access-Token")
		resp := helper.HttpGet(orderApiUrl + "/order", headers, getQuery)
		var orderItem order
		json.Unmarshal(resp, &orderItem)

		fullPath := Config.SoundFiles.Path + params["order-id"] + "." + Config.SoundFiles.Type

		if orderItem.Id != "" {
			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				filename := helper.HttpSaveFile(res, req, fullPath)
				if filename != "" {
					orderItem.SoundFilePath = helper.SoundFileUrl(req, params["order-id"])
					result.Success = orderItem
					return result
				}
			}
		}
		result.Error.Msg = "Invalid orderId: " + params["order-id"]
		return result
	})
}
