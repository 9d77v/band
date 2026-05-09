package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/9d77v/band/cmd/band/tpls"
	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "init a web server",
	Long:  `init a web server`,
	Run: func(cmd *cobra.Command, args []string) {
		if pkgDir == "" {
			log.Fatalln("module path is required, use --pkg flag")
		}
		fmt.Println("start init project")
		initProject()
		fmt.Println("finish init project")
		fmt.Println("start make init")
		execCmd("make", "init")
		fmt.Println("finish make init")
		handleEnvFile()
		fmt.Println("init project success")
	},
}

func handleEnvFile() {
	f, err := os.Open("env.sample")
	if err != nil {
		log.Panicln("err:", err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		log.Panicln("err:", err)
		return
	}
	f.Close()
	file, err := os.OpenFile(".env", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Panicln("打开文件失败", err)
	}
	defer file.Close()
	file.Write(data)
	os.Rename("env.sample", ".env.sample")

}
func execCmd(name string, args ...string) {
	var envVars []string
	var cmdArgs []string
	
	// Extract environment variables from arguments
	for _, arg := range args {
		if strings.Contains(arg, "=") && !strings.HasPrefix(arg, "-") {
			envVars = append(envVars, arg)
		} else {
			cmdArgs = append(cmdArgs, arg)
		}
	}
	
	goCmd := exec.Command(name, cmdArgs...)
	
	// Set environment variables if any were provided
	if len(envVars) > 0 {
		goCmd.Env = append(os.Environ(), envVars...)
	}
	
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
	dirs, err := tpls.ServerFiles.ReadDir(serverDir)
	if err != nil {
		log.Panicln("打开server文件夹失败", err)
	}
	serverTpl := &ServerTpl{
		PKG_DIR:  pkgDir,
		APP_NAME: appName,
	}
	walkServerDir(serverTpl, ".", serverDir, dirs)
}

func walkServerDir(tpl *ServerTpl, localDir, parentServerDir string, dirs []fs.DirEntry) {
	for _, v := range dirs {
		name := v.Name()
		serverPath := filepath.Join(parentServerDir, name)
		localPath := filepath.Join(localDir, name)
		fmt.Println(serverPath)

		if v.IsDir() {
			childDirs, err := tpls.ServerFiles.ReadDir(serverPath)
			if err != nil {
				log.Panicln("打开", serverPath, "文件夹失败", err)
			}
			err = os.MkdirAll(localPath, os.ModePerm)
			if err != nil {
				log.Panicln("创建目录失败", err)
			}
			walkServerDir(tpl, localPath, serverPath, childDirs)
		} else {
			fd, err := tpls.ServerFiles.ReadFile(serverPath)
			if err != nil {
				log.Panicln("打开", serverPath, "文件失败", err)
			}
			fileName := strings.ReplaceAll(localPath, ".tpl", "")
			tmpl, err := template.New("server").Parse(string(fd))
			if err != nil {
				log.Panicln("模版解析失败", err)
			}
			file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0755)
			if err != nil {
				log.Panicln("打开文件失败", err)
			}
			err = tmpl.Execute(file, tpl)
			if err != nil {
				log.Panicln("写文件失败", err)
			}
			file.Close()
		}
	}
}
