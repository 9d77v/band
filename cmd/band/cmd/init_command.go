package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/9d77v/band/cmd/band/tpls"
	"github.com/9d77v/band/pkg/utils"
	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "init a web server",
	Long:  `init a web server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("begin init project")
		if utils.FileExist("go.mod") {
			fmt.Println("project created")
			return
		}
		initProject()
		execCmd("go", "mod", "init", pkgDir)
		execCmd("make", "init")
		handleEnvFile()
		fmt.Println("end init project")
	},
}

func handleEnvFile() {
	f, err := os.Open("env.sample")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	f.Close()
	file, err := os.OpenFile(".env", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Println("打开文件失败", err)
	}
	defer file.Close()
	file.Write(data)
	os.Rename("env.sample", ".env.sample")

}
func execCmd(name string, args ...string) {
	goCmd := exec.Command(name, args...)
	var stdout, stderr bytes.Buffer
	goCmd.Stdout = &stdout // 标准输出
	goCmd.Stderr = &stderr // 标准错误
	err := goCmd.Run()
	if err != nil {
		log.Fatalf("goCmd.Run() failed with %s\n", err)
	}
	fmt.Println(stdout.String())
	fmt.Println(stderr.String())
}

const (
	serverDir = "server"
)

type ServerTpl struct {
	PKG_DIR  string
	APP_NAME string
}

// initProject 项目初始化
func initProject() {
	if len(appName) == 0 {
		if len(pkgDir) != 0 {
			arr := strings.Split(pkgDir, "/")
			appName = arr[len(arr)-1]
		} else {
			panic("项目名称不能为空")
		}
	}
	dirs, err := tpls.ServerFiles.ReadDir(serverDir)
	if err != nil {
		log.Println("打开server文件夹失败", err)
	}
	serverTpl := &ServerTpl{
		PKG_DIR:  pkgDir,
		APP_NAME: appName,
	}
	walkServerDir(serverTpl, "./", serverDir, dirs)
}

func walkServerDir(tpl *ServerTpl, localDir, serverDir string, dirs []fs.DirEntry) {
	for _, v := range dirs {
		name := v.Name()
		serverPath := serverDir + "/" + name
		fmt.Println(serverPath)
		localPath := localDir + "/" + strings.Replace(name, "example", service, -1)
		if v.IsDir() {
			dirs, err := tpls.ServerFiles.ReadDir(serverPath)
			if err != nil {
				log.Println("打开", serverPath, "文件夹失败", err)
			}
			os.MkdirAll(localPath, os.ModePerm)
			walkServerDir(tpl, localPath, serverPath, dirs)
		} else {
			fd, err := tpls.ServerFiles.ReadFile(serverPath)
			if err != nil {
				log.Println("打开", serverPath, "文件失败", err)
			}
			fileName := strings.ReplaceAll(localPath, ".tpl", "")
			tmpl, err := template.New("server").Parse(string(fd))
			if err != nil {
				log.Println("模版解析失败", err)
			}
			file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0755)
			if err != nil {
				log.Println("打开文件失败", err)
			}
			defer file.Close()
			err = tmpl.Execute(file, tpl)
			if err != nil {
				log.Println("写文件失败", err)
			}
		}
	}
}
