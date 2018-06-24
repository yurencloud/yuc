package template

import (
	"yuc/util"
)

var middleware = `package middleware

import (
	"net/http"
	"yugo/session"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := session.GetInstance(w, r)
		user := sess.Values["user"]
		if user != nil {
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	})
}`

func GenerateMiddleware()  {
	//以读写方式打开文件，如果不存在，则创建
	util.GenerateFile("./middleware/auth.conf", middleware)
}


