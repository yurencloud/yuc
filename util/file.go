package util

import (
	"os"
	"fmt"
	"strings"
	"path/filepath"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateFolder(folderName string)  {
	var path string
	if os.IsPathSeparator('\\') {  //前边的判断是否是系统的分隔符
		path = "\\"
	} else {
		path = "/"
	}

	dir, _ := os.Getwd()  //当前的目录

	target := dir+path+folderName

	isExist, err := PathExists(target)

	if err != nil {
		fmt.Println(err)
	}

	if isExist {
		return
	}

	err2 := os.Mkdir(target, os.ModePerm)  //在当前目录下生成md目录
	if err2 != nil {
		fmt.Println(err2)
	}
}

func GenerateFile(filePath string, content string)  {
	//以读写方式打开文件，如果不存在，则创建
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0766);
	if err != nil {
		fmt.Println(err);
	}
	file.Write([]byte(content))
	file.Close();
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func GetCurrentDirectoryName() string {
	path := GetCurrentDirectory()
	return path[strings.LastIndex(path, "/")+1:]
}


