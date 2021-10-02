CHECK_FILES?=$$(go list ./... | grep -v /vendor/)

.PHONY: image
image: ## Build the Docker image.
	docker build -t go-starter . -f ./d8t/Dockerfile

.PHONY: test-prepare
test-prepare: ## Run migrations.
	go run main.go migrate --config config.test.yaml

.PHONY: test
test: ## Run tests.
	go test -p 1 -v $(CHECK_FILES)

.PHONY: migrate-test
migrate-test: ## Run migrations.
	go run main.go migrate --config config.test.yaml