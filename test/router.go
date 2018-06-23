package main

import (
	"github.com/gorilla/mux"
	. "yugo-template/controller"
	"yugo-template/middleware"
)

func InitRouter(router *mux.Router) {

	router.HandleFunc("/", Index)
	router.HandleFunc("/login", Login)

	home := router.PathPrefix("/home").Subrouter()
	home.HandleFunc("", Home)
	home.Use(middleware.Auth)
}