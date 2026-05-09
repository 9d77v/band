APP={{.APP_NAME}}

IMAGE_TAG=$(shell git log --pretty=format:"%ad_%h" -1 --date=short)
# 初始化项目
init:
	go mod tidy
	go install github.com/google/wire/cmd/wire@latest
fix:
	go fix ./...
migrate:
	go run tools/migration/main.go
gosec:
	gosec -conf gosec.json ./...
docker:
	docker-compose up -d
.PHONY : all 
