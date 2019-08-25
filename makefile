.PHONY: start
start: stop up dep server

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

.PHONY: test
test: ## run test
		docker-compose exec go go test ./src/Controller