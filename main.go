package main

import (
	"os/exec"
	"bytes"
	"log"
	"os"
	"github.com/urfave/cli"
	"github.com/fatih/color"
	"fmt"
	"bufio"
	"io"
)

const VERSION = "1.0.1"

func main() {
	//isVersion := flag.Bool("v", false, "version")
	//
	//init := flag.Bool("init", false, "create a new web project")


	// 命令类型
	//isRegister := flag.Bool("register", false, "register command")
	//isLogin := flag.Bool("login", false, "login command")
	//isReset := flag.Bool("reset", false, "reset password command")
	//
	//// 注册/登录用户
	//username := flag.String("u", "", "register by username")
	//password := flag.String("p", "", "your password")
	//
	//// 创建命令
	//cmd := flag.String("c", "", "save your command prompt to web")
	//comment := flag.String("m", "", "comment your command prompt")
	//
	//// 删除指定id的命令
	//id := flag.Int("d", 0, "delete cmd by id")

	//flag.Parse()
	//
	//
	//
	//if *isVersion {
	//	color.Green("yuc version "+VERSION)
	//}
	//
	//if *init {
	//	color.Green("Init yugo project now ...")
	//	_, err := exec_shell("git clone https://github.com/yurencloud/yugo-template.git")
	//	if err!=nil {
	//		color.Red("clone project fail")
	//		panic(err)
	//	}
	//	exec_shell("cp -rf yugo-template/* ./")
	//	exec_shell("rm -rf yugo-template")
	//	exec_shell("rm -rf ./.git")
	//	result, _ := exec_shell("ls -R |awk '{print i$0}' i=`pwd`'/'")
	//	color.Blue(result)
	//	color.Green("Init successful!")
	//}

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
				_, err := exec_shell("git clone https://github.com/yurencloud/yugo-template.git")
				if err!=nil {
					color.Red("clone project fail")
					panic(err)
				}
				exec_shell("cp -rf yugo-template/* ./")
				exec_shell("rm -rf yugo-template")
				exec_shell("rm -rf ./.git")
				result, _ := exec_shell("ls -R |awk '{print i$0}' i=`pwd`'/'")
				color.Blue(result)
				color.Green("Init successful!")
				return nil
			},
		},
	}

	app.Run(os.Args)
}

func exec_shell(s string) (string, error){
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()

	if err != nil {
		log.Println(err)
	}

	return out.String(), err
}

func execCommand(commandName string, params []string) bool {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command(commandName, params...)

	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Print(line)
	}

	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	cmd.Wait()
	return true
}



