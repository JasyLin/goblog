package controllers

import (
	"jasy/goblog/pkg/view"
	"net/http"
)

type AuthController struct {

}

func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {

}