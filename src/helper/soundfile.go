package helper

import "net/http"

func SoundFileUrl(req *http.Request, orderId string) (requestUrl string) {
	requestUrl = "http://" + req.Host + "/soundfile/" + orderId
	return
}