package service

import (
	"{{.PKG_DIR}}/apps/{{.SERVICE_PACKAGE}}/domain/repository"
	{{.SERVICE_LOWER}}pb "{{.PKG_DIR}}/proto/{{.SERVICE_PACKAGE}}pb"
)

type {{.SERVICE_UPPER}}AppService struct {
	{{.SERVICE_UPPER}}Repository repository.{{.SERVICE_UPPER}}Repository
}

func New{{.SERVICE_UPPER}}AppService(repo repository.{{.SERVICE_UPPER}}Repository) {{.SERVICE_LOWER}}pb.{{.SERVICE_UPPER}}ServiceServer {
	return &{{.SERVICE_UPPER}}AppService{
		{{.SERVICE_UPPER}}Repository: repo}
}
