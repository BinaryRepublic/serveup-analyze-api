package controller

import (
	"net/http"
)

func GetSoundFile(res http.ResponseWriter, req *http.Request) {
	HandleMessage([]string{"order-id"}, res, req, func(params map[string]string) interface{} {
		// authorization > root only
		if req.Header.Get("accountId") == "root" {
			filePath := Config.SoundFiles.Path + params["order-id"] + "." + Config.SoundFiles.Type
			FileResponse(res, req, filePath)
		}
		return false
	})
}