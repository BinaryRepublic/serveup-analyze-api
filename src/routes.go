package main

import (
	"net/http"
	"controller"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"GetOrder",
		"GET",
		"/order/{id}",
		controller.GetOrderById,
	},
	Route{
		"GetOrder",
		"GET",
		"/order/restaurant",
		controller.GetOrderByRestaurant,
	},
	Route{
		"PostOrder",
		"POST",
		"/order",
		controller.PostOrder,
	},
	Route{
		"GetSoundFile",
		"GET",
		"/soundfile/{order-id}",
		controller.GetSoundFile,
	},
}