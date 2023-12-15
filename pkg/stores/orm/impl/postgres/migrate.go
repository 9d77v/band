package postgres

import (
	"fmt"
	"log"

	"github.com/9d77v/band/pkg/stores/orm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDatabaseIfNotExist(conf orm.Conf) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s sslmode=disable password=%s",
		conf.Host, conf.Port, conf.User, conf.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("connect to postgres failed:", err)
	}
	if databaseNotExist(db, conf) {
		createDatabase(db, conf)
	}
	sqlDBInit, _ := db.DB()
	sqlDBInit.Close()
}

func databaseNotExist(db *gorm.DB, conf orm.Conf) bool {
	var total int64
	err := db.Raw("SELECT 1 FROM pg_database WHERE datname = ?", conf.DBName).Scan(&total).Error
	if err != nil {
		log.Println("check db failed", err)
	}
	return total == 0
}

func createDatabase(db *gorm.DB, conf orm.Conf) {
	initSQL := fmt.Sprintf("CREATE DATABASE \"%s\" WITH  OWNER =%s ENCODING = 'UTF8' CONNECTION LIMIT=-1;",
		conf.DBName, conf.User)
	err := db.Exec(initSQL).Error
	if err != nil {
		log.Println("create db failed:", err)
	} else {
		log.Printf("create db '%s' succeed\n", conf.DBName)
	}
}
