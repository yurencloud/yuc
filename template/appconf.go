package template

import (
	"github.com/yurencloud/yuc/util"
)

var appconf = `app.name = yugo
port = 8080
`

func GenerateAppconf()  {
	util.GenerateFile("./config/app.conf", appconf)
}
