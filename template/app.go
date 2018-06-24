package template

import (
	"strings"
	"github.com/yurencloud/yuc/util"
)

var main = `package main

func main()  {
	Run()
}`

var router = `package main

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
`

var run = `package main

import (
	"github.com/gorilla/mux"
	"yugo/config"
	"strings"
	"net/http"
	"github.com/gorilla/csrf"
	_ "yugo/log"
	"github.com/sirupsen/logrus"
	"strconv"
)

func staticServer(router *mux.Router) {
	staticConfig := config.Get("static")
	staticArray := strings.Split(staticConfig, ",")
	// 生成一个或多个静态目录，默认static,可自行修改，或添加，以英文逗号分隔
	for index := range staticArray {
		static := staticArray[index]
		router.PathPrefix("/").Handler(http.StripPrefix("/"+static, http.FileServer(http.Dir(static))))
	}
}

func Run() {

	router := mux.NewRouter()

	InitRouter(router)

	staticServer(router)

	configMap := config.GetConfigMap()

	appName := configMap["app.name"]

	logrus.Info("app: " + appName + ", started at port " + configMap["port"])

	if configMap["csrf.enabled"] == "true" {
		maxAge, _ := strconv.Atoi(configMap["csrf.max.age"])
		CSRF := csrf.Protect(
			[]byte(configMap["csrf.key"]),
			csrf.RequestHeader(configMap["csrf.request.header"]),
			csrf.FieldName(configMap["csrf.field.name"]),
			csrf.MaxAge(maxAge),
			csrf.Secure(false), // 本地开发要加false,生产环境要加true
		)
		http.ListenAndServe(":"+configMap["port"], CSRF(router))
	}else{
		http.ListenAndServe(":"+configMap["port"], router)
	}
}
`

func GenerateApp() {
	util.GenerateFile("./main.go", main)


	router := strings.Replace(router, "yugo-template", util.GetCurrentDirectoryName(), -1)
	util.GenerateFile("./router.go", router)

	util.GenerateFile("./run.go", run)
}




