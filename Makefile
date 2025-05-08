##@ 🎯 The commands are:

GO_ENV := GO111MODULE=on GOPROXY=https://goproxy.cn,direct
IMG ?= devflow:latest

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<command>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: all
all: clean build

.PHONY: run
run: ## Running devflow
	@go run main.go

.PHONY: build
build: ## Build devflow binary
	@echo "👉 Compile devflow..."
	$(GO_ENV) go build -v -o devflow
	@echo "✅ Compile complete."

.PHONY: build-linux
build-linux: ## ## Build devflow Linux binary
	GOOS=linux GOARCH=amd64 $(GO_ENV) go build -v -o devflow

.PHONY: docker-build
docker-build: ## Build docker image with the devflow.
	@echo "📝 Generating temporary Dockerfile..."
	@echo "\
FROM golang:1.23.4-alpine as Builder \n\
WORKDIR /app \n\
COPY . . \n\
RUN $(GO_ENV) go build -o devflow \n\
\n\
FROM alpine \n\
COPY --from=Builder /app/devflow /usr/local/bin/devflow \n\
ENTRYPOINT [\"devflow\"] \n\
" > Dockerfile
	@echo "✅ Dockerfile created."
	@echo "Builder image"
	@docker build -t ${IMG} .

.PHONY: clean
clean: ## Clean generated files
	@rm -rf ./Dockerfile
	@rm -rf ./devflow