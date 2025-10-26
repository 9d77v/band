APP={{.APP_NAME}}

IMAGE_TAG=$(shell git log --pretty=format:"%ad_%h" -1 --date=short)
# 初始化项目
init: 
	go install github.com/google/wire/cmd/wire@latest
# 同步数据库结构
migrate:
	go run tools/migrator/main.go
# 代码安全检查
gosec:
	gosec -conf gosec.json ./...
docker:
	docker-compose up -d
dev:
	go run cmd/server/*.go
.PHONY : all 
