package Auth

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
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


func WraperLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Info(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
