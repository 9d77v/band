APP={{.APP_NAME}}

IMAGE_TAG=$(shell git log --pretty=format:"%ad_%h" -1 --date=short)
# Services configuration
SERVICES := 
# 初始化项目
init:
	go mod tidy
	go install github.com/google/wire/cmd/wire@latest
fix:
	go fix ./...
migrate:
	go run tools/migrator/*.go
gosec:
	gosec -conf gosec.json ./...
docker:
	docker-compose up -d

# ==============================================
# Template Functions for Services
# ==============================================

# Protoc template for generating Go code from .proto files
define protoc_template
protoc-$(1): api/protobuf/$(1)pb/*.proto
	protoc -I./api/protobuf/$(1)pb \
	--go_out=. \
	--go-grpc_out=require_unimplemented_servers=false:. \
	api/protobuf/$(1)pb/*.proto
endef

# Wire template
define wire_template
wire-$(1):
	cd cmd/apps/$(1)-service && wire gen && mv wire.go wire.go.back

wire-$(1)-clean:
	cd cmd/apps/$(1)-service && rm wire_gen.go && mv wire.go.back wire.go
endef

# Service run template
define run_service_template
$(1)-service:
	go run cmd/apps/$(1)-service/*.go
endef

# Generate all protoc rules
$(foreach service,$(SERVICES),$(eval $(call protoc_template,$(service))))

# Generate wire rules
$(foreach service,$(SERVICES),$(eval $(call wire_template,$(service))))

# Generate service run rules
$(foreach service,$(SERVICES),$(eval $(call run_service_template,$(service))))

.PHONY : all

