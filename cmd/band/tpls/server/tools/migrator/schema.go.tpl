package main

import (
	"log"

	"github.com/9d77v/band/pkg/stores/orm/impl/postgres"
)

// migrateSchemas 自动迁移数据库表结构
func migrateSchemas(db *postgres.PgDB) {
	err := db.GetDB().AutoMigrate(
	)
	if err != nil {
		log.Println("auto migrate error:", err)
	}
}
