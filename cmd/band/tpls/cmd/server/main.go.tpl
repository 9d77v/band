package main

import (
	"github.com/9d77v/band/pkg/log"
	"github.com/9d77v/band/pkg/utils"

	"github.com/joho/godotenv"
)

func main() {
	if utils.FileExist(".env") {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}
	log.Init()
	app, err := initApp("{{.SERVICE_PACKAGE}}-service")
	if err != nil {
		panic(err)
	}
	app.Run()
}
