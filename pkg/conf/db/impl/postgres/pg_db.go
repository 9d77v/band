package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/9d77v/band/pkg/conf/db"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	client *gorm.DB
	once   sync.Once
)

type PgDB struct {
	db   *gorm.DB
	conf db.Conf
}

var PgDBIns = wire.NewSet(NewPgDB,
	wire.Bind(new(db.DB), new(*PgDB)))

func NewPgDB(conf db.Conf) (*PgDB, error) {
	var err error
	once.Do(func() {
		client, err = newClient(conf)
		if err != nil {
			log.Panicf("Could not initialize gorm: %s\n", err.Error())
		}
		log.Println("connected to db:", client)
	})
	return &PgDB{
		db:   client,
		conf: conf,
	}, err
}

func newClient(conf db.Conf) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		conf.Host, conf.Port, conf.User, conf.DBName, conf.Password)
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   conf.TablePrefix,
		},
	}
	if conf.Debug {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormConfig.DisableForeignKeyConstraintWhenMigrating = true
	}
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
	return db, err
}

func FromDB(db *gorm.DB, conf db.Conf) db.DB {
	return &PgDB{db: db, conf: conf}
}

func (db *PgDB) GetConf() db.Conf {
	return db.conf
}

// Session create new db session
func (db *PgDB) Session(config *gorm.Session) db.DB {
	return FromDB(db.db.Session(config), db.conf)
}

// WithContext change current instance db's context to ctx
func (db *PgDB) WithContext(ctx context.Context) db.DB {
	return FromDB(db.db.WithContext(ctx), db.conf)
}

// Debug start debug mode
func (db *PgDB) Debug() db.DB {
	return FromDB(db.db.Debug(), db.conf)
}

// Debug start debug mode
func (db *PgDB) GetDB() *gorm.DB {
	return db.db
}

// Set store value with key into current db instance's context
func (db *PgDB) Set(key string, value interface{}) db.DB {
	return FromDB(db.db.Set(key, value), db.conf)
}

// Get get value with key from current db instance's context
func (db *PgDB) Get(key string) (interface{}, bool) {
	return db.db.Get(key)
}

// InstanceSet store value with key into current db instance's context
func (db *PgDB) InstanceSet(key string, value interface{}) db.DB {
	return FromDB(db.db.InstanceSet(key, value), db.conf)
}

// InstanceGet get value with key from current db instance's context
func (db *PgDB) InstanceGet(key string) (interface{}, bool) {
	return db.db.InstanceGet(key)
}

// AddError add error to db
func (db *PgDB) AddError(err error) error {
	return db.db.AddError(err)
}

// DB returns `*sql.DB`
func (db *PgDB) DB() (*sql.DB, error) {
	return db.db.DB()
}

// SetupJoinTable setup join table schema
func (db *PgDB) SetupJoinTable(model interface{}, field string, joinTable interface{}) error {
	return db.db.SetupJoinTable(model, field, joinTable)
}

// Use use plugin
func (db *PgDB) Use(plugin gorm.Plugin) error {
	return db.db.Use(plugin)
}

func (db *PgDB) ToSQL(queryFn func(tx *gorm.DB) *gorm.DB) string {
	return db.db.ToSQL(queryFn)
}

// Model specify the model you would like to run db operations
//
//	// update all users's name to `hello`
//	db.Model(&User{}).Update("name", "hello")
//	// if user's primary key is non-blank, will use it as condition, then will only update the user's name to `hello`
//	db.Model(&user).Update("name", "hello")
func (db *PgDB) Model(value interface{}) db.DB {
	return FromDB(db.db.Model(value), db.conf)
}

// Clauses Add clauses
func (db *PgDB) Clauses(conds ...clause.Expression) db.DB {
	return FromDB(db.db.Clauses(conds...), db.conf)
}

// Table specify the table you would like to run db operations
func (db *PgDB) Table(name string, args ...interface{}) db.DB {
	return FromDB(db.db.Table(name, args...), db.conf)
}

// Distinct specify distinct fields that you want querying
func (db *PgDB) Distinct(args ...interface{}) db.DB {
	return FromDB(db.db.Distinct(args...), db.conf)
}

// Select specify fields that you want when querying, creating, updating
func (db *PgDB) Select(query interface{}, args ...interface{}) db.DB {
	return FromDB(db.db.Select(query, args...), db.conf)
}

// Omit specify fields that you want to ignore when creating, updating and querying
func (db *PgDB) Omit(columns ...string) db.DB {
	return FromDB(db.db.Omit(columns...), db.conf)
}

// Where add conditions
func (db *PgDB) Where(query interface{}, args ...interface{}) db.DB {
	return FromDB(db.db.Where(query, args...), db.conf)
}

// Not add NOT conditions
func (db *PgDB) Not(query interface{}, args ...interface{}) db.DB {
	return FromDB(db.db.Not(query, args...), db.conf)
}

// Or add OR conditions
func (db *PgDB) Or(query interface{}, args ...interface{}) db.DB {
	return FromDB(db.db.Or(query, args...), db.conf)
}

// Joins specify Joins conditions
//
//	db.Joins("Account").Find(&user)
//	db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Find(&user)
//	db.Joins("Account", DB.Select("id").Where("user_id = users.id AND name = ?", "someName").Model(&Account{}))
func (db *PgDB) Joins(query string, args ...interface{}) db.DB {
	return FromDB(db.db.Joins(query, args...), db.conf)
}

// Group specify the group method on the find
func (db *PgDB) Group(name string) db.DB {
	return FromDB(db.db.Group(name), db.conf)
}

// Having specify HAVING conditions for GROUP BY
func (db *PgDB) Having(query interface{}, args ...interface{}) db.DB {
	return FromDB(db.db.Having(query, args...), db.conf)
}

// Order specify order when retrieve records from database
//
//	db.Order("name DESC")
//	db.Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}, Desc: true})
func (db *PgDB) Order(value interface{}) db.DB {
	return FromDB(db.db.Order(value), db.conf)
}

// Limit specify the number of records to be retrieved
func (db *PgDB) Limit(limit int) db.DB {
	return FromDB(db.db.Limit(limit), db.conf)
}

// Offset specify the number of records to skip before starting to return the records
func (db *PgDB) Offset(offset int) db.DB {
	return FromDB(db.db.Offset(offset), db.conf)
}

// Scopes pass current database connection to arguments `func(DB) DB`, which could be used to add conditions dynamically
//
//	func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
//	    return db.Where("amount > ?", 1000)
//	}
//
//	func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
//	    return func (db *gorm.DB) *gorm.DB {
//	        return db.Scopes(AmountGreaterThan1000).Where("status in (?)", status)
//	    }
//	}
//
//	db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
func (db *PgDB) Scopes(funcs ...func(db *gorm.DB) *gorm.DB) db.DB {
	return FromDB(db.db.Scopes(funcs...), db.conf)
}

// Preload preload associations with given conditions
//
//	db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
func (db *PgDB) Preload(query string, args ...interface{}) db.DB {
	return FromDB(db.db.Preload(query, args...), db.conf)
}

func (db *PgDB) Attrs(attrs ...interface{}) db.DB {
	return FromDB(db.db.Attrs(attrs...), db.conf)
}

func (db *PgDB) Assign(attrs ...interface{}) db.DB {
	return FromDB(db.db.Assign(attrs...), db.conf)
}

func (db *PgDB) Unscoped() db.DB {
	return FromDB(db.db.Unscoped(), db.conf)
}

func (db *PgDB) Raw(sql string, values ...interface{}) db.DB {
	return FromDB(db.db.Raw(sql, values...), db.conf)
}

func (db *PgDB) Error() error {
	return db.db.Error
}

// Create insert the value into database
func (db *PgDB) Create(value interface{}) db.DB {
	return FromDB(db.db.Create(value), db.conf)
}

// CreateInBatches insert the value in batches into database
func (db *PgDB) CreateInBatches(value interface{}, batchSize int) db.DB {
	return FromDB(db.db.CreateInBatches(value, batchSize), db.conf)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (db *PgDB) Save(value interface{}) db.DB {
	return FromDB(db.db.Save(value), db.conf)
}

// First find first record that match given conditions, order by primary key
func (db *PgDB) First(dest interface{}, conds ...interface{}) db.DB {
	return FromDB(db.db.First(dest, conds...), db.conf)
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (db *PgDB) Take(dest interface{}, conds ...interface{}) db.DB {
	return FromDB(db.db.Take(dest, conds...), db.conf)
}

// Last find last record that match given conditions, order by primary key
func (db *PgDB) Last(dest interface{}, conds ...interface{}) db.DB {
	return FromDB(db.db.Last(dest, conds...), db.conf)
}

// Find find records that match given conditions
func (db *PgDB) Find(dest interface{}, conds ...interface{}) db.DB {
	return FromDB(db.db.Find(dest, conds...), db.conf)
}

// FindInBatches find records in batches
func (db *PgDB) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) db.DB {
	return FromDB(db.db.FindInBatches(dest, batchSize, fc), db.conf)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (db *PgDB) FirstOrInit(dest interface{}, conds ...interface{}) db.DB {
	return FromDB(db.db.FirstOrInit(dest, conds...), db.conf)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (db *PgDB) FirstOrCreate(dest interface{}, conds ...interface{}) db.DB {
	return FromDB(db.db.FirstOrCreate(dest, conds...), db.conf)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (db *PgDB) Update(column string, value interface{}) db.DB {
	return FromDB(db.db.Update(column, value), db.conf)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (db *PgDB) Updates(values interface{}) db.DB {
	return FromDB(db.db.Updates(values), db.conf)
}

func (db *PgDB) UpdateColumn(column string, value interface{}) db.DB {
	return FromDB(db.db.UpdateColumn(column, value), db.conf)
}

func (db *PgDB) UpdateColumns(values interface{}) db.DB {
	return FromDB(db.db.UpdateColumns(values), db.conf)
}

// Delete delete value match given conditions, if the value has primary key, then will including the primary key as condition
func (db *PgDB) Delete(value interface{}, conds ...interface{}) db.DB {
	return FromDB(db.db.Delete(value, conds...), db.conf)
}

func (db *PgDB) Count(count *int64) db.DB {
	return FromDB(db.db.Count(count), db.conf)
}

func (db *PgDB) Row() *sql.Row {
	return db.db.Row()
}

func (db *PgDB) Rows() (*sql.Rows, error) {
	return db.db.Rows()
}

// Scan scan value to a struct
func (db *PgDB) Scan(dest interface{}) db.DB {
	return FromDB(db.db.Scan(dest), db.conf)
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (db *PgDB) Pluck(column string, dest interface{}) db.DB {
	return FromDB(db.db.Pluck(column, dest), db.conf)
}

func (db *PgDB) ScanRows(rows *sql.Rows, dest interface{}) error {
	return db.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (db *PgDB) Connection(fc func(db *gorm.DB) error) (err error) {
	return db.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (db *PgDB) Transaction(fc func(db *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return db.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (db *PgDB) Begin(opts ...*sql.TxOptions) db.DB {
	return FromDB(db.db.Begin(opts...), db.conf)
}

// Commit commit a transaction
func (db *PgDB) Commit() db.DB {
	return FromDB(db.db.Commit(), db.conf)
}

// Rollback rollback a transaction
func (db *PgDB) Rollback() {
	db.db.Rollback()
}

func (db *PgDB) SavePoint(name string) db.DB {
	return FromDB(db.db.SavePoint(name), db.conf)
}

func (db *PgDB) RollbackTo(name string) db.DB {
	return FromDB(db.db.RollbackTo(name), db.conf)
}

// Exec execute raw sql
func (db *PgDB) Exec(sql string, values ...interface{}) db.DB {
	return FromDB(db.db.Exec(sql, values...), db.conf)
}

func (db *PgDB) Migrator() gorm.Migrator {
	return db.db.Migrator()
}

func (db *PgDB) AutoMigrate(dst ...interface{}) error {
	return db.db.AutoMigrate(dst...)
}

func (db *PgDB) Association(column string) *gorm.Association {
	return db.db.Association(column)
}
