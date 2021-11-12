package main

import (
	"jasy/goblog/app/http/middlewares"
	"jasy/goblog/bootstrap"
	"jasy/goblog/config"
	c "jasy/goblog/pkg/config"
	"net/http"
)

func init() {
	config.Initialize()
}

func main() {

	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()


	http.ListenAndServe(":" + c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))

}
