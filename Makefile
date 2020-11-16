.PHONY: deps migrate test
CHECK_FILES?=$$(go list ./... | grep -v /vendor/)

deps: ## Install dependencies.
	go mod download
	go mod verify

migrate_test: ## Run migrations.
	go run main.go migrate --config config/config.test.yaml

test: ## Run tests.
	go test -p 1 -v $(CHECK_FILES)
