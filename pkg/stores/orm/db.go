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
	Set(key string, value any) DB
	Get(key string) (any, bool)
	InstanceSet(key string, value any) DB
	InstanceGet(key string) (any, bool)
	AddError(err error) error
	DB() (*sql.DB, error)
	SetupJoinTable(model any, field string, joinTable any) error
	Use(plugin gorm.Plugin) error
	ToSQL(queryFn func(tx *gorm.DB) *gorm.DB) string
	Model(value any) DB
	Clauses(conds ...clause.Expression) DB
	Table(name string, args ...any) DB
	Distinct(args ...any) DB
	Select(query any, args ...any) DB
	Omit(columns ...string) DB
	Where(query any, args ...any) DB
	Not(query any, args ...any) DB
	Or(query any, args ...any) DB
	Joins(query string, args ...any) DB
	Group(name string) DB
	Having(query any, args ...any) DB
	Order(value any) DB
	Limit(limit int) DB
	Offset(offset int) DB
	Scopes(funcs ...func(db *gorm.DB) *gorm.DB) DB
	Preload(query string, args ...any) DB
	Attrs(attrs ...any) DB
	Assign(attrs ...any) DB
	Unscoped() DB
	Raw(sql string, values ...any) DB
	Error() error
	Create(value any) DB
	CreateInBatches(value any, batchSize int) DB
	Save(value any) DB
	First(dest any, conds ...any) DB
	Take(dest any, conds ...any) DB
	Last(dest any, conds ...any) DB
	Find(dest any, conds ...any) DB
	FindInBatches(dest any, batchSize int, fc func(tx *gorm.DB, batch int) error) DB
	FirstOrInit(dest any, conds ...any) DB
	FirstOrCreate(dest any, conds ...any) DB
	Update(column string, value any) DB
	Updates(values any) DB
	UpdateColumn(column string, value any) DB
	UpdateColumns(values any) DB
	Delete(value any, conds ...any) DB
	Count(count *int64) DB
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	Scan(dest any) DB
	Pluck(column string, dest any) DB
	ScanRows(rows *sql.Rows, dest any) error
	Connection(fc func(db *gorm.DB) error) (err error)
	Transaction(fc func(db *gorm.DB) error, opts ...*sql.TxOptions) (err error)
	Begin(opts ...*sql.TxOptions) DB
	Commit() DB
	Rollback()
	SavePoint(name string) DB
	RollbackTo(name string) DB
	Exec(sql string, values ...any) DB
	Migrator() gorm.Migrator
	AutoMigrate(dst ...any) error
	Association(column string) *gorm.Association
}
