IS_IN_PROGRESS = "is in progress ..."

## setup: Set up database temporary for integration testing
.PHONY: setup
setup:
	@docker-compose -f ./infrastructure/docker-compose.yml up -d
	@sleep 10

## down: Set down database temporary for integration testing
.PHONY: down
down: 
	@docker-compose -f ./infrastructure/docker-compose.yml down -t 1

## mod: tidy up golang module
.PHONY: mod
mod: 
	@go mod tidy && go mod vendor

## gen: will generate mock for usecases & repositories interfaces
.PHONY: gen
gen:
	@echo "make gen ${IS_IN_PROGRESS}"
	@mockgen -source internal/${DOMAIN}/usecase.go -destination internal/${DOMAIN}/mocks/usecase_mock.go -package=mocks
	@mockgen -source internal/${DOMAIN}/repository.go -destination internal/${DOMAIN}/mocks/repository_mock.go -package=mocks