package routes

import (
	"github.com/gorilla/mux"
	"jasy/goblog/app/http/controllers"
	"net/http"
)

func RegisterWebRoutes(r *mux.Router) {

	pc := controllers.PagesController{}

	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
}