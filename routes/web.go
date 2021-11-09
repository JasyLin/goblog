package routes

import (
	"github.com/gorilla/mux"
	"jasy/goblog/app/http/controllers"
	"jasy/goblog/app/http/middlewares"
	"net/http"
)

func RegisterWebRoutes(r *mux.Router) {

	pc := new(controllers.PagesController)

	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)

	ac := new(controllers.ArticlesController)

	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")
	r.HandleFunc("/articles/create", ac.Create).Methods("GET").Name("articles.create")
	r.HandleFunc("/articles", ac.Store).Methods("POST").Name("articles.store")
	r.HandleFunc("/articles/{id:[0-9]+}/edit", ac.Edit).Methods("GET").Name("articles.edit")
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Update).Methods("POST").Name("articles.update")
	r.HandleFunc("/articles/{id:[0-9]+}/delete", ac.Delete).Methods("POST").Name("articles.delete")


	// 用户认证
	auc := new(controllers.AuthController)
	r.HandleFunc("/auth/register", auc.Register).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/doregister", auc.DoRegister).Methods("POST").Name("auth.doregister")
	r.HandleFunc("/auth/login", auc.Login).Methods("GET").Name("auth.login")
	r.HandleFunc("/auth/dologin", auc.DoLogin).Methods("POST").Name("auth.dologin")

	// 静态资源
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	r.Use(middlewares.StartSession)

}