package repository

import (
	"context"

	"github.com/9d77v/band/pkg/stores/orm/base"
	"{{.PKG_DIR}}/domain/entity"
)

type {{.SERVICE_UPPER}}RepoTxFunc = func(ctx context.Context, repo {{.SERVICE_UPPER}}Repository) error

type {{.SERVICE_UPPER}}Repository interface {
	Tx(ctx context.Context, f {{.SERVICE_UPPER}}RepoTxFunc) error
	List{{.ENTITY_UPPER}}(ctx context.Context, q base.SearchCriteria) ([]*entity.{{.ENTITY_UPPER}}, int64, error)
	Get{{.ENTITY_UPPER}}ByID(ctx context.Context, id int64) (*entity.{{.ENTITY_UPPER}}, error)
	Create{{.ENTITY_UPPER}}(ctx context.Context, in *entity.{{.ENTITY_UPPER}}) (*entity.{{.ENTITY_UPPER}}, error)
	Update{{.ENTITY_UPPER}}ByID(ctx context.Context, in *entity.{{.ENTITY_UPPER}}) (*entity.{{.ENTITY_UPPER}}, error)
	SoftDelete{{.ENTITY_UPPER}}ByIDs(ctx context.Context, ids ...int64) error
}
