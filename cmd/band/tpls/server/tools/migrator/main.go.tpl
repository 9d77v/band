package main

import (
	"log"

	"github.com/9d77v/band/pkg/stores/orm"
	"github.com/9d77v/band/pkg/stores/orm/impl/postgres"
	"github.com/joho/godotenv"
)

func main() {
	db := initDB()
	migrateEnumTypes(db)
	migrateSchemas(db)
	seedProducts(db)
	seedPrompts(db)
}

// initDB 初始化数据库连接
func initDB() *postgres.PgDB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("加载环境变量失败: %s", err)
	}
	conf := orm.FromEnv()
	postgres.CreateDatabaseIfNotExist(conf)
	db, err := postgres.NewPgDB(conf)
	if err != nil {
		log.Fatalf("数据库连接失败: %s", err)
	}
	return db
}
