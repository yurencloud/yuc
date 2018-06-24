package main

import (
	"github.com/gorilla/mux"
	"yugo-template/controller"
	"yugo-template/middleware"
)

func InitRouter(router *mux.Router) {

	router.HandleFunc("/", controller.Index)
	router.HandleFunc("/login", controller.Login)

	home := router.PathPrefix("/home").Subrouter()
	home.HandleFunc("", controller.Home)
	home.Use(middleware.Auth)
}