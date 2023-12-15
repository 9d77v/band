package orm

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DB interface {
	GetConf() Conf
	Session(config *gorm.Session) DB
	WithContext(ctx context.Context) DB
	Debug() DB
	GetDB() *gorm.DB
	Set(key string, value interface{}) DB
	Get(key string) (interface{}, bool)
	InstanceSet(key string, value interface{}) DB
	InstanceGet(key string) (interface{}, bool)
	AddError(err error) error
	DB() (*sql.DB, error)
	SetupJoinTable(model interface{}, field string, joinTable interface{}) error
	Use(plugin gorm.Plugin) error
	ToSQL(queryFn func(tx *gorm.DB) *gorm.DB) string
	Model(value interface{}) DB
	Clauses(conds ...clause.Expression) DB
	Table(name string, args ...interface{}) DB
	Distinct(args ...interface{}) DB
	Select(query interface{}, args ...interface{}) DB
	Omit(columns ...string) DB
	Where(query interface{}, args ...interface{}) DB
	Not(query interface{}, args ...interface{}) DB
	Or(query interface{}, args ...interface{}) DB
	Joins(query string, args ...interface{}) DB
	Group(name string) DB
	Having(query interface{}, args ...interface{}) DB
	Order(value interface{}) DB
	Limit(limit int) DB
	Offset(offset int) DB
	Scopes(funcs ...func(db *gorm.DB) *gorm.DB) DB
	Preload(query string, args ...interface{}) DB
	Attrs(attrs ...interface{}) DB
	Assign(attrs ...interface{}) DB
	Unscoped() DB
	Raw(sql string, values ...interface{}) DB
	Error() error
	Create(value interface{}) DB
	CreateInBatches(value interface{}, batchSize int) DB
	Save(value interface{}) DB
	First(dest interface{}, conds ...interface{}) DB
	Take(dest interface{}, conds ...interface{}) DB
	Last(dest interface{}, conds ...interface{}) DB
	Find(dest interface{}, conds ...interface{}) DB
	FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) DB
	FirstOrInit(dest interface{}, conds ...interface{}) DB
	FirstOrCreate(dest interface{}, conds ...interface{}) DB
	Update(column string, value interface{}) DB
	Updates(values interface{}) DB
	UpdateColumn(column string, value interface{}) DB
	UpdateColumns(values interface{}) DB
	Delete(value interface{}, conds ...interface{}) DB
	Count(count *int64) DB
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	Scan(dest interface{}) DB
	Pluck(column string, dest interface{}) DB
	ScanRows(rows *sql.Rows, dest interface{}) error
	Connection(fc func(db *gorm.DB) error) (err error)
	Transaction(fc func(db *gorm.DB) error, opts ...*sql.TxOptions) (err error)
	Begin(opts ...*sql.TxOptions) DB
	Commit() DB
	Rollback()
	SavePoint(name string) DB
	RollbackTo(name string) DB
	Exec(sql string, values ...interface{}) DB
	Migrator() gorm.Migrator
	AutoMigrate(dst ...interface{}) error
	Association(column string) *gorm.Association
}
