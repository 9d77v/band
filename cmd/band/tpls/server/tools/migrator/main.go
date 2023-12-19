package main

import (
	"fmt"
	"log"

	"github.com/9d77v/band/pkg/stores/orm"
	"github.com/9d77v/band/pkg/stores/orm/impl/postgres"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	conf := orm.FromEnv()
	fmt.Println(conf)
	postgres.CreateDatabaseIfNotExist(conf)
	db, err := postgres.NewPgDB(conf)
	if err != nil {
		panic("new db connection failed")
	}
	err = db.GetDB().AutoMigrate(
	//TODO
	)
	if err != nil {
		log.Println("auto migrate error:", err)
	}
}
