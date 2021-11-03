package main

import (
	"jasy/goblog/app/http/middlewares"
	"jasy/goblog/bootstrap"
	"net/http"
)

func main() {

	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()


	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))

	//type test struct {
	//	a string
	//	b string
	//}
	//var asd *test
	//asd = &test{a: "111", b: "222"}
	//
	//fmt.Println(asd.a)
}
