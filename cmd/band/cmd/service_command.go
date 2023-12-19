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
		fmt.Println("Start add service")
		initService()
		fmt.Println("Ends add service")
	},
}

const (
	serviceDir = "service"
)

type ServiceTpl struct {
	PKG_DIR       string
	SERVICE_UPPER string
	SERVICE_LOWER string
	ENTITY_UPPER  string
	ENTITY_LOWER  string
	ID_TYPE       string
}

// initService
func initService() {
	moduleName := util.GetGoModule("")
	if len(pkgDir) == 0 {
		if len(moduleName) != 0 {
			pkgDir = moduleName
		} else {
			panic("module name can not be empty")
		}
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
		PKG_DIR:       pkgDir,
		SERVICE_UPPER: strings.ToUpper(service[:1]) + service[1:],
		SERVICE_LOWER: service,
		ENTITY_UPPER:  strings.ToUpper(entity[:1]) + entity[1:],
		ENTITY_LOWER:  entity,
		ID_TYPE:       idType,
	}
	servicelocalDir := "./internal/apps/" + service
	if utils.FileExist(servicelocalDir) {
		fmt.Println("service created")
		return
	}
	os.MkdirAll(servicelocalDir, os.ModePerm)
	walkServiceDir(serviceTpl, servicelocalDir, serviceDir, dirs)
	handleServiceProto(serviceTpl)
	handleServiceCmd(serviceTpl)
	fmt.Println("add service finished")
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
	protoDir := "./internal/proto/" + service + "pb"
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

	if _, err = file.WriteString(`protoc-` + service + `: internal/proto/` + service + `pb/*.proto
	protoc -I./internal/proto/` + service + `pb \
	-I./internal/proto/include \
	--go_out=. \
	--go-grpc_out=require_unimplemented_servers=false:. \
	internal/proto/` + service + `pb/*.proto
wire-` + service + `:
	cd cmd/apps/project-service && wire gen && mv wire.go wire.go.back && cd ../../../
` + service + `-service:
	go run cmd/apps/` + service + `-service/*.go`); err != nil {
		panic(err)
	}
	execCmd("make", `protoc-`+service)
}

func handleServiceCmd(tpl *ServiceTpl) {
	cmdDir := "./cmd/apps/" + service + "-service"
	os.MkdirAll(cmdDir, os.ModePerm)
	files := []string{"main.go", "wire.go"}
	for _, v := range files {
		fd, err := tpls.CmdFiles.ReadFile("cmd/" + v + ".tpl")
		if err != nil {
			log.Println("open", v+".tpl", "file failed:", err)
		}
		fileName := cmdDir + "/" + v
		writeTplFile(tpl, fd, fileName)
	}
	execCmd("make", `wire-`+service)
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
