// + build wireinject
package main

import (
	"github.com/9d77v/band/pkg/app"
	"github.com/9d77v/band/pkg/stores/orm"
	"github.com/9d77v/band/pkg/stores/orm/orm_factory"
	"{{.PKG_DIR}}/internal/apps/{{.SERVICE_LOWER}}"
	"{{.PKG_DIR}}/internal/apps/{{.SERVICE_LOWER}}/application/service"
	"{{.PKG_DIR}}/internal/apps/{{.SERVICE_LOWER}}/infrastructure/persistence"
	"github.com/google/wire"
)

func initApp(serviceName string) (*{{.SERVICE_LOWER}}.App, error) {
	wire.Build(
		app.RPCFromEnv,
		orm.FromEnv,orm_factory.OrmSigleton,
		persistence.New{{.SERVICE_UPPER}}RepositoryImpl,
		service.New{{.SERVICE_UPPER}}AppService, {{.SERVICE_LOWER}}.NewApp)
	return &{{.SERVICE_LOWER}}.App{}, nil
}
