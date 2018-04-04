package controller

import (
	"net/http"
)

func GetSoundFile(res http.ResponseWriter, req *http.Request) {
	HandleMessage([]string{"order-id"}, res, req, func(params map[string]string) ControllerResult {
		var result ControllerResult
		// authorization > root only
		if req.Header.Get("accountId") == "root" {
			filePath := Config.SoundFiles.Path + params["order-id"] + "." + Config.SoundFiles.Type
			FileResponse(res, req, filePath)
		}
		result.Error.Msg = "not authorized to access soundfiles: root only"
		return result
	})
}