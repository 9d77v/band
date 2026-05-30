package cmd

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/9d77v/band/cmd/band/tpls"
	"github.com/9d77v/band/cmd/band/util"
	"github.com/9d77v/band/pkg/utils"
	"github.com/spf13/cobra"
)

var serviceCommand = &cobra.Command{
	Use:   "service",
	Short: "init a service module",
	Long:  `init a service module`,
	Run: func(cmd *cobra.Command, args []string) {
		addNewService()
	},
}

const (
	serviceDir = "service"
	cmdDir     = "cmd"
)

type ServiceTpl struct {
	PKG_DIR         string
	SERVICE_UPPER   string
	SERVICE_LOWER   string
	SERVICE_PACKAGE string
	ENTITY_UPPER    string
	ENTITY_LOWER    string
	ENTITY_PACKAGE  string
	ID_TYPE         string
}

// addNewService
func addNewService() {
	if service == "" {
		log.Fatalln("service name is required")
	}
	path := "./apps/" + service
	if utils.FileExist(path) {
		log.Println("service already exists")
		return
	}
	serviceDir := serviceDir
	dirs, err := tpls.ServiceFiles.ReadDir(serviceDir)
	if err != nil {
		log.Println("open service directory failed:", err)
	}
	if entity == "" {
		entity = service
	}
	serviceTpl := &ServiceTpl{
		PKG_DIR:         getPkgDirFromGoMod(),
		SERVICE_UPPER:   util.UnderscoreToCamelCase(service),
		SERVICE_LOWER:   util.FirstLower(util.UnderscoreToCamelCase(service)),
		SERVICE_PACKAGE: service,
		ENTITY_UPPER:    util.UnderscoreToCamelCase(entity),
		ENTITY_LOWER:    util.FirstLower(util.UnderscoreToCamelCase(entity)),
		ENTITY_PACKAGE:  entity,
		ID_TYPE:         idType,
	}
	servicelocalDir := "./apps/" + service
	if utils.FileExist(servicelocalDir) {
		fmt.Println("service created")
		return
	}
	os.MkdirAll(servicelocalDir, os.ModePerm)
	walkServiceDir(serviceTpl, servicelocalDir, serviceDir, dirs)
	fmt.Println("copy service files finished")
	handleServiceProto(serviceTpl)
	fmt.Println("generate service proto finished")
	handleServiceCmd(serviceTpl)
	fmt.Println("generate service cmd finished")
	fmt.Println("start make init")
	execCmd("make", "init")
	fmt.Println("finish make init")
	execCmd("make", "GOWORK=off", `wire-`+service)
	fmt.Println("wire service finished")
}

func walkServiceDir(tpl *ServiceTpl, localDir, serverDir string, dirs []fs.DirEntry) {
	for _, v := range dirs {
		name := v.Name()
		serverPath := serverDir + "/" + name
		name = strings.ReplaceAll(name, "example", service)
		name = strings.ReplaceAll(name, "model", entity)
		localPath := localDir + "/" + name
		if v.IsDir() {
			dirs, err := tpls.ServiceFiles.ReadDir(serverPath)
			if err != nil {
				log.Println("open ", serverPath, "directory failed:", err)
			}
			os.MkdirAll(localPath, os.ModePerm)
			walkServiceDir(tpl, localPath, serverPath, dirs)
		} else {
			fd, err := tpls.ServiceFiles.ReadFile(serverPath)
			if err != nil {
				log.Println("open", serverPath, "file failed:", err)
			}
			fileName := strings.ReplaceAll(localPath, ".tpl", "")
			tmpl, err := template.New("service").Parse(string(fd))
			if err != nil {
				log.Println("template parse failed:", err)
			}
			file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0755)
			if err != nil {
				log.Println("open file failed:", err)
			}
			defer file.Close()
			err = tmpl.Execute(file, tpl)
			if err != nil {
				log.Println("write file failed:", err)
			}
		}
	}
}

func handleServiceProto(tpl *ServiceTpl) {
	protoDir := "./api/protobuf/" + service + "pb"
	os.MkdirAll(protoDir, os.ModePerm)
	fd, err := tpls.ProtoFiles.ReadFile("proto/example.proto.tpl")
	if err != nil {
		log.Println("open", "example.proto.tpl", "file failed:", err)
	}
	fileName := protoDir + "/" + service + ".proto"
	writeTplFile(tpl, fd, fileName)
	updateMakefileServices(service)
	execCmd("make", `protoc-`+service)
}

func updateMakefileServices(serviceName string) {
	data, err := os.ReadFile("Makefile")
	if err != nil {
		log.Println("read Makefile failed:", err)
		return
	}

	content := string(data)
	servicesLine := "SERVICES := "

	if !strings.Contains(content, servicesLine) {
		log.Println("SERVICES configuration not found in Makefile")
		return
	}

	// 在 "SERVICES := " 行后面添加 service 名称
	content = strings.Replace(content, servicesLine, servicesLine+serviceName+" ", 1)

	err = os.WriteFile("Makefile", []byte(content), 0600)
	if err != nil {
		log.Println("write Makefile failed:", err)
		return
	}
}

func writeTplFile(tpl *ServiceTpl, fd []byte, fileName string) {
	tmpl, err := template.New("service").Parse(string(fd))
	if err != nil {
		log.Println("template parse failed:", err)
	}
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Println("open file failed:", err)
	}
	defer file.Close()
	err = tmpl.Execute(file, tpl)
	if err != nil {
		log.Println("write file failed:", err)
	}
}

func handleServiceCmd(tpl *ServiceTpl) {
	cmdLocalDir := "./cmd/apps/" + service + "-service"
	os.MkdirAll(cmdLocalDir, os.ModePerm)
	dirs, err := tpls.CmdFiles.ReadDir(cmdDir + "/server")
	if err != nil {
		log.Println("open cmd directory failed:", err)
	}
	for _, v := range dirs {
		name := v.Name()
		serverPath := cmdDir + "/server/" + name
		fd, err := tpls.CmdFiles.ReadFile(serverPath)
		if err != nil {
			log.Println("open", serverPath, "file failed:", err)
		}
		fileName := cmdLocalDir + "/" + strings.ReplaceAll(name, ".tpl", "")
		tmpl, err := template.New("cmd").Parse(string(fd))
		if err != nil {
			log.Println("template parse failed:", err)
		}
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			log.Println("open file failed:", err)
		}
		defer file.Close()
		err = tmpl.Execute(file, tpl)
		if err != nil {
			log.Println("write file failed:", err)
		}
	}
}

// getPkgDirFromGoMod 从项目根目录的 go.mod 文件中读取 module path
func getPkgDirFromGoMod() string {
	f, err := os.Open("go.mod")
	if err != nil {
		log.Fatalln("open go.mod failed:", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if after, ok := strings.CutPrefix(line, "module "); ok {
			return strings.TrimSpace(after)
		}
	}
	log.Fatalln("module path not found in go.mod")
	return ""
}
