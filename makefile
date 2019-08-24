.PHONY: start
start: stop up dep

.PHONY: stop
stop: ## stop environment
		docker-compose stop

.PHONY: dep
dep: ## spin up environment
		docker-compose exec go dep ensure

.PHONY: up
up: ## up docker
		docker-compose up -d