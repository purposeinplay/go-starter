CHECK_FILES?=$$(go list ./... | grep -v /vendor/)

.PHONY: image
image: ## Build the Docker image.
	docker build -t go-starter . -f ./d8t/Dockerfile

.PHONY: test
test: ## Run tests.
	go test -p 1 -v $(CHECK_FILES)