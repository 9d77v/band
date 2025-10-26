package cmd

import (
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
		PKG_DIR:         pkgDir,
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
	execCmd("make", `init-`+service)
	fmt.Println("init service finished")
	execCmd("make", `wire-`+service)
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
	file, err := os.OpenFile("Makefile", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err = file.WriteString(`protoc-` + service + `: api/protobuf/` + service + `pb/*.proto
	protoc -I./api/protobuf/` + service + `pb \
	--go_out=. \
	--go-grpc_out=require_unimplemented_servers=false:. \
	api/protobuf/` + service + `pb/*.proto
init-` + service + `:
	cd apps/` + service + ` && go mod init ` + pkgDir + `
	go work use apps/` + service + `
	cd apps/` + service + ` && go mod tidy
wire-` + service + `:
	cd apps/` + service + `/cmd/server && wire gen && mv wire.go wire.go.back
dev-` + service + `:
	go run apps/` + service + `/cmd/server/*.go`); err != nil {
		panic(err)
	}
	execCmd("make", `protoc-`+service)
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
