package site

import (
	"net/http"

	"github.com/abhiche/health-check/pkg/logger"
	"github.com/globalsign/mgo"

	"github.com/gorilla/mux"
)

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var staticDir = "/web/"

//NewRouter configures a new router to the API
func NewRouter(s *mgo.Session) *mux.Router {
	var controller = &Controller{Repository: Repository{s}}

	// Routes defines the list of routes of our API
	type Routes []Route

	var routes = Routes{
		Route{
			"Index",
			"GET",
			"/sites",
			controller.Index,
		},
		Route{
			"AddSite",
			"POST",
			"/sites",
			controller.AddSite,
		},
		Route{
			"PatchSite",
			"PATCH",
			"/sites/{id}",
			controller.PatchSite,
		},
		Route{
			"DeleteSite",
			"DELETE",
			"/sites/{id}",
			controller.DeleteSite,
		},
	}
	router := mux.NewRouter().StrictSlash(true)
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
