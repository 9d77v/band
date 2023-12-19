package {{.SERVICE_LOWER}}

import (
	"github.com/9d77v/band/pkg/app"
	"{{.PKG_DIR}}/internal/proto/{{.SERVICE_LOWER}}pb"
	"google.golang.org/grpc"
)

type App struct {
	app.App
	{{.SERVICE_UPPER}}Service {{.SERVICE_LOWER}}pb.{{.SERVICE_UPPER}}ServiceServer
}

func NewApp(conf app.Conf, svc {{.SERVICE_LOWER}}pb.{{.SERVICE_UPPER}}ServiceServer) (*App, error) {
	return &App{
		App:         app.NewApp(conf),
		{{.SERVICE_UPPER}}Service: svc,
	}, nil
}

func (a *App) Run() {
	a.StartGrpcServer(func(srv *grpc.Server) {
		{{.SERVICE_LOWER}}pb.Register{{.SERVICE_UPPER}}ServiceServer(srv, a.{{.SERVICE_UPPER}}Service)
	})
}
