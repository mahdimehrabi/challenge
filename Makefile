include .env
MIGRATE=migrate -path=internal/migration -database "postgres://${DBUsername}:${DBPassword}@${DBHost}:${DBPort}/${DBName}?sslmode=disable" -verbose

db-migrate-up:
		$(MIGRATE) up
db-migrate-down:
		$(MIGRATE) down
db-force:
		@read -p  "Which version do you want to force?" VERSION; \
		$(MIGRATE) force $$VERSION

db-goto:
		@read -p  "Which version do you want to migrate?" VERSION; \
		$(MIGRATE) goto $$VERSION

db-drop:
		$(MIGRATE) drop

db-create-migration:
		@read -p  "What is the name of migration?" NAME; \
		${MIGRATE} create -ext sql -seq -dir migration  $$NAME

fmt:
	gofumpt -l -w .

devtools:
	@echo "Installing devtools"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install mvdan.cc/gofumpt@latest

check:
	golangci-lint run \
		--build-tags "${BUILD_TAG}" \
		--timeout=20m0s \
		--enable=gofmt \
		--enable=unconvert \
		--enable=unparam \
		--enable=asciicheck \
		--enable=misspell \
		--enable=revive \
		--enable=decorder \
		--enable=reassign \
		--enable=usestdlibvars \
		--enable=nilerr \
		--enable=gosec \
		--enable=exportloopref \
		--enable=whitespace \
		--enable=goimports \
		--enable=gocyclo \
		--enable=nestif \
		--enable=gochecknoinits \
		--enable=gocognit \
		--enable=funlen \
		--enable=forbidigo \
		--enable=godox \
		--enable=gocritic \
		--enable=gci \
		--enable=lll

unit_test:
	go test ./... -v
