# Note: @ symbol will not print out command to screen

build:
	@go build -o bin/vitals-server-api .

run-build: build
	@./bin/vitals-server-api

run: swagger
	@go run main.go

lint:
    # Install golangci-lint from here: https://golangci-lint.run/usage/install/#macos
	@golangci-lint run

clean:
	@rm -rf ./bin/* && rm -f vitals.db && rm -f vitals-server-api

clean-cache:
	@go clean -modcache && make build


swagger:
    # Install gin-swagger from here: https://github.com/swaggo/gin-swagger
    # Run: go install github.com/swaggo/swag/cmd/swag@latest
    # Ensure that GOPATH/bin is in the system PATH.
    # To find GOPATH, run `go env GOPATH`
	@swag init


# Docker commands

build-docker: swagger
    # Build Swagger documentation before building Docker image
    # so we don't need to install the gin-swagger CLI in the Docker image
	@docker build -t vitals-server-api .

run-docker: build-docker
    # Run the Docker image with the environment variables loaded from .env
    # and override the PORT to use :8080 instead of localhost:8080
	@docker run -p 8080:8080 --env-file .env -e PORT=':8080' vitals-server-api

stop-docker:
	@docker stop $(shell docker ps -q --filter ancestor=vitals-server-api)
