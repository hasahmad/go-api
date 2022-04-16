include .envrc

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	go run ./internal/cmd/api

## db/psql: connect to the database using psql
.PHONY: db/psql
db/psql:
	psql

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new: build/migrate confirm
	@echo 'Creating migration files for ${name}...'
	./bin/migrate -dir=./internal/db/migrations create ${name} go

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: build/migrate confirm
	@echo 'Running up migrations...'
	./bin/migrate -dir=./internal/db/migrations up

## db/migrations/down: rollback previous database migration
.PHONY: db/migrations/down
db/migrations/down: build/migrate confirm
	@echo 'Running down migrations...'
	./bin/migrate -dir=./internal/db/migrations down


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# ==================================================================================== #
# BUILD
# ==================================================================================== #

current_time = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
git_description = $(shell git describe --always --dirty --tags --long)
linker_flags = '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building internal/cmd/api...'
	go build -ldflags=${linker_flags} -o=./bin/api ./internal/cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o=./bin/linux_amd64/api ./internal/cmd/api

## build/migrate: build the cmd/migrate application
.PHONY: build/migrate
build/migrate:
	go build -o=./bin/migrate ./internal/cmd/migrate
	GOOS=linux GOARCH=amd64 go build -o=./bin/linux_amd64/migrate ./internal/cmd/migrate
