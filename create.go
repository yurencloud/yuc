package main

import (
	"github.com/yurencloud/yuc/util"
	"github.com/yurencloud/yuc/template"
)

func CreateNewProject()  {
	folder := []string{ "config", "controller", "log", "middleware", "model", "static"}
	for i := range folder {
		util.CreateFolder(folder[i])
	}
	template.GenerateAppconf()
	template.GenerateController()
	template.GenerateMiddleware()
	template.GenerateGitignore()
	template.GenerateApp()
}