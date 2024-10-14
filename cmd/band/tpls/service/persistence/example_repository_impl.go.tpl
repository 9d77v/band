package persistence

import (
	"context"
	"errors"

	"github.com/9d77v/band/pkg/stores/orm"
	"github.com/9d77v/band/pkg/stores/orm/base"
	"github.com/9d77v/band/pkg/stores/orm/impl/postgres"
	"{{.PKG_DIR}}/apps/{{.SERVICE_PACKAGE}}/domain/entity"
	"{{.PKG_DIR}}/apps/{{.SERVICE_PACKAGE}}/domain/repository"
	"{{.PKG_DIR}}/apps/{{.SERVICE_PACKAGE}}/persistence/do"

	"gorm.io/gorm"
)

type {{.SERVICE_UPPER}}RepositoryImpl struct {
	DB orm.DB
}

func New{{.SERVICE_UPPER}}RepositoryImpl(db orm.DB) repository.{{.SERVICE_UPPER}}Repository {
	return &{{.SERVICE_UPPER}}RepositoryImpl{DB: db}
}

// Tx implements repository.{{.SERVICE_UPPER}}Repository
func (r *{{.SERVICE_UPPER}}RepositoryImpl) Tx(ctx context.Context, fun func(ctx context.Context, repo repository.{{.SERVICE_UPPER}}Repository) error) error {
	trx := r.DB.WithContext(ctx)
	return trx.Transaction(func(tx *gorm.DB) error {
		repo := New{{.SERVICE_UPPER}}RepositoryImpl(postgres.FromDB(tx, r.DB.GetConf()))
		return fun(ctx, repo)
	})
}

// List{{.ENTITY_UPPER}} implements repository.{{.SERVICE_UPPER}}Repository.
func (r *{{.SERVICE_UPPER}}RepositoryImpl) List{{.ENTITY_UPPER}}(ctx context.Context, q base.SearchCriteria) ([]*entity.{{.ENTITY_UPPER}}, int64, error) {
	{{.ENTITY_LOWER}}s, total, err := base.Page[do.{{.ENTITY_UPPER}}](ctx, r.DB.GetDB(), &do.{{.ENTITY_UPPER}}{}, q)
	return do.To{{.ENTITY_UPPER}}s({{.ENTITY_LOWER}}s), total, err
}

// Create{{.ENTITY_UPPER}} implements repository.{{.SERVICE_UPPER}}Repository.
func (r *{{.SERVICE_UPPER}}RepositoryImpl) Create{{.ENTITY_UPPER}}(ctx context.Context, in *entity.{{.ENTITY_UPPER}}) (*entity.{{.ENTITY_UPPER}}, error) {
	db := r.DB.WithContext(ctx)
	{{.ENTITY_LOWER}} := do.New{{.ENTITY_UPPER}}FromEntity(in)
	err := db.Create({{.ENTITY_LOWER}}).Error()
	return {{.ENTITY_LOWER}}.To{{.ENTITY_UPPER}}(), err
}

// Update{{.ENTITY_UPPER}} implements repository.{{.SERVICE_UPPER}}Repository.
func (r *{{.SERVICE_UPPER}}RepositoryImpl) Update{{.ENTITY_UPPER}}ByID(ctx context.Context, in *entity.{{.ENTITY_UPPER}}) (*entity.{{.ENTITY_UPPER}}, error) {
	db := r.DB.WithContext(ctx)
	{{.ENTITY_LOWER}} := do.New{{.ENTITY_UPPER}}FromEntity(in)
	err := db.Model({{.ENTITY_LOWER}}).Updates({{.ENTITY_LOWER}}).Error()
	return {{.ENTITY_LOWER}}.To{{.ENTITY_UPPER}}(), err
}

// Get{{.ENTITY_UPPER}}ByID implements repository.{{.SERVICE_UPPER}}Repository.
func (r *{{.SERVICE_UPPER}}RepositoryImpl) Get{{.ENTITY_UPPER}}ByID(ctx context.Context, id int64) (*entity.{{.ENTITY_UPPER}}, error) {
	db := r.DB.WithContext(ctx)
	{{.ENTITY_LOWER}} := new(do.{{.ENTITY_UPPER}})
	err := db.First({{.ENTITY_LOWER}}, id).Error()
	return {{.ENTITY_LOWER}}.To{{.ENTITY_UPPER}}(), err
}

// SoftDelete{{.ENTITY_UPPER}} implements repository.{{.SERVICE_UPPER}}Repository.
func (r *{{.SERVICE_UPPER}}RepositoryImpl) SoftDelete{{.ENTITY_UPPER}}ByIDs(ctx context.Context, ids ...int64) error {
	db := r.DB.WithContext(ctx)
	var err error
	if len(ids) == 0 {
		err = errors.New("id cannot be empty")
	} else if len(ids) == 1 {
		err = db.Where("id = ?", ids[0]).Delete(&do.{{.ENTITY_UPPER}}{}).Error()
	} else {
		err = db.Where("id in ?", ids).Delete(&do.{{.ENTITY_UPPER}}{}).Error()
	}
	return err
}

// Delete{{.ENTITY_UPPER}} implements repository.{{.SERVICE_UPPER}}Repository.
func (r *{{.SERVICE_UPPER}}RepositoryImpl) Delete{{.ENTITY_UPPER}}ByIDs(ctx context.Context, ids ...int64) error {
	db := r.DB.WithContext(ctx)
	var err error
	if len(ids) == 0 {
		err = errors.New("id cannot be empty")
	} else if len(ids) == 1 {
		err = db.Unscoped().Where("id = ?", ids[0]).Delete(&do.{{.ENTITY_UPPER}}{}).Error()
	} else {
		err = db.Unscoped().Where("id in ?", ids).Delete(&do.{{.ENTITY_UPPER}}{}).Error()
	}
	return err
}
