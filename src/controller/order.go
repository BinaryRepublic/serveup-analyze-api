package controller

import (
	"net/http"
	"helper"
	"strconv"
	"encoding/json"
	"os"
)

var Config = helper.ReadConfig()
var orderApiUrl = Config.OrderApi.Host + ":" + strconv.Itoa(Config.OrderApi.Port)

type orderItem struct {
	Id		string			`json:"id"`
	Name	string			`json:"name"`
	Size	int				`json:"size"`
	Nb		int				`json:"nb"`
}
type order struct {
	Id            string      `json:"id"`
	Timestamp     string      `json:"timestamp"`
	Items         []orderItem `json:"items"`
	VoiceDeviceId string      `json:"voicedevice-id"`
	RestaurantId  string      `json:"restaurant-id"`
	Status        int         `json:"status"`
	SoundFilePath string      `json:"soundfile-path"`
}

func GetOrder(res http.ResponseWriter, req *http.Request) {

	HandleMessage([]string{"restaurant-id"}, res, req, func(params map[string]string) interface{} {
		res := helper.HttpGet(orderApiUrl + "/order/restaurant", helper.HttpQueryParams(req))
		var orders []order
		json.Unmarshal(res, &orders)

		for index := range orders {
			soundFilePath := Config.SoundFiles.Path + orders[index].Id + "." + Config.SoundFiles.Type
			order := &orders[index]
			if _, err := os.Stat(soundFilePath); err == nil {
				order.SoundFilePath = soundFilePath
			}
		}

		return orders
	})
}

func PostOrder(res http.ResponseWriter, req *http.Request) {

	HandleMessage([]string{"order-id", "soundfile"}, res, req, func(params map[string]string) interface{} {

		getQuery := make(map[string]string)
		getQuery["id"] = params["order-id"]
		resp := helper.HttpGet(orderApiUrl + "/order", getQuery)

		filename := helper.HttpSaveFile(res, req, Config.SoundFiles.Path + params["order-id"] + "." + Config.SoundFiles.Type)
		if filename != "" {
			var orderItem order
			json.Unmarshal(resp, &orderItem)
			orderItem.SoundFilePath = filename
			return orderItem
		}
		return nil
	})
}
