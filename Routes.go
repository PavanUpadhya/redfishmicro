package main

import(
	"github.com/gorilla/mux"
	"net/http"
)

//Route struct defines the structure for all Routes
type Route struct{
	name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

// Routes - Slice of Routes
type Routes []Route

var routes = Routes {
	Route{
		"discoverBasicInfo",
		"POST",
		"/discoverBasicInfo",
		discoverBasicInfoHandler,
	},
	Route{
		"getPowerState",
		"POST",
		"/getPowerState",
		discoverPowerState,
	},
}

/*GetAndRegisterRouter...
* Method to register all routes and return
* the router
*/
func GetAndRegisterRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		router.Methods(route.Method).
		Path(route.Pattern).
		Name(route.name).
		Handler(route.HandlerFunc)
	}
	return router
}
