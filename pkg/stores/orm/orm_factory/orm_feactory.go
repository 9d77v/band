package orm_factory

import (
	"sync"

	"github.com/9d77v/band/pkg/stores/orm"
	"github.com/9d77v/band/pkg/stores/orm/impl/postgres"
)

var (
	client orm.DB
	once   sync.Once
)

const (
	DriverPostgres = "postgres"
)

func NewOrm(conf orm.Conf) (orm.DB, error) {
	var client orm.DB
	var err error
	switch conf.Driver {
	default:
		client, err = postgres.NewPgDB(conf)
	}
	return client, err
}

func OrmSigleton(conf orm.Conf) (orm.DB, error) {
	var err error
	once.Do(func() {
		client, err = NewOrm(conf)
	})
	return client, err
}
