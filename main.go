package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"jasy/goblog/app/http/middlewares"
	"jasy/goblog/bootstrap"
	"jasy/goblog/pkg/database"
	"net/http"
)

var router *mux.Router
var db *sql.DB


func main() {

	database.Initialize()
	db = database.DB

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()


	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
