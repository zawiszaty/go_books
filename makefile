.PHONY: start
start: stop up dep server

.PHONY: start_test
start_test: stop up dep server_test

.PHONY: stop
stop: ## stop environment
		docker-compose stop

.PHONY: dep
dep: ## spin up environment
		docker-compose exec go dep ensure

.PHONY: up
up: ## up docker
		docker-compose up -d

.PHONY: server
server: ## up server
		docker-compose exec go go run src/*.go

.PHONY: server_test
server_test: ## up server test
		docker-compose exec go go run src/*.go &

.PHONY: test
test: ## run test
		docker-compose exec go go test ./src/Controller