package template

import (
	"yuc/util"
)

var index = `package controller

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write( []byte("Hello World"))
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write( []byte("Pass Auth to Home"))
}
`

var login = `package controller

import "net/http"

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write( []byte("Login"))
}
`

func GenerateController()  {
	util.GenerateFile("./controller/index.go", index)
	util.GenerateFile("./controller/login.go", login)
}
