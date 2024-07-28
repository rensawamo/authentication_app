d.DEFAULT_GOAL := help

.PHONY: help test

SQL_MIGRATE_CLI		= sql-migrate

export UID_GID=$(shell id -u):$(shell id -g)

help: ## show this message
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-8s\033[0m %s\n", $$1, $$2}'


buf-gen: ## buf generate code
	buf generate --path proto && go mod tidy


deps: ## install dependencies
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	go install github.com/envoyproxy/protoc-gen-validate@v0.9.0
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@v1.11.1
	go install github.com/bufbuild/buf/cmd/buf@v1.15.0

# mysql targets
up: 		## Up
	docker compose -f database/docker-compose.yml up --force-recreate

.PHONY: chmod
chmod: 		## Up
	sh scripts/mysql_chmod.sh


down: 	## Down
	docker compose -f database/docker-compose.yml down

logs:	## Show logs
	docker compose -f database/docker-compose.yml logs -f

login:	## login db
	docker compose -f database/docker-compose.yml exec db /bin/bash

clean: 		## clean	
clean: down
	docker compose -f database/docker-compose.yml rm -f
	rm -rf database/mysql/data/*

mysql-client:	## connet mysql from mysql cli
	docker compose -f database/docker-compose.yml exec db /bin/bash -c "mysql -u root -p -D db"

