package main

import (
	"os"
	"github.com/urfave/cli"
	"github.com/fatih/color"
	"yuc/util"
)

const VERSION = "1.0.1"

func main() {
	app := cli.NewApp()

	app.Name = "yuc"
	app.Usage = "yugo cli tool"
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Init a yugo web project",
			Action:  func(c *cli.Context) error {
				color.Green("Init yugo project now ...")
				CreateNewProject()
				result, _ := util.ExecShell("ls -R |awk '{print i$0}' i=`pwd`'/'")
				color.Blue(result)
				color.Green("Init successful!")
				return nil
			},
		},
	}

	app.Run(os.Args)
}



