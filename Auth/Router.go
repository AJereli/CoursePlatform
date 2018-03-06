package Auth

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct{
	Name, Method, Pattern string
	HandlerFunc http.HandlerFunc

}

type Routes []Route

type HandlFunc func(w http.ResponseWriter, r *http.Request)

var routes = Routes{
	Route{
		"LogIn",
		"GET",
		"/login",
		Login,

	},
	Route{
		"Registration",
		"GET",
		"/registration",
		Registration,
	},
}


func InitRouter () *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes{
		var handler http.Handler
		handler = route.HandlerFunc
		handler = WraperLogger(handler, route.Name)

		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}

