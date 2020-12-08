# Simple makefile
NAME=kairo
VERSION=v1

## check: run static analysis
check:
	# @sonar-scanner -Dsonar.projectKey="$SONAR_PROJECT_KEY" -Dsonar.sources=. -Dsonar.host.url="$SONAR_HOST" -Dsonar.login="$SONAR_LOGIN"
	@gosec -quiet ./...

## test: Run test
test:
	go test ./... -coverprofile cp.out && go tool cover -func=cp.out

## coverage: Run test with html coverage report
coverage:
	go test ./... -coverprofile cp.out && go tool cover -html=cp.out

## benchmark: Run benchmark test
benchmark:
	go test -bench=.

## watch: development with air
watch:
	air -c .air/development.air.toml

## build: Build binary applications
build:
	@go generate ./...
	@echo building binary to ./dist/${NAME}
	@go build -o ./dist/${NAME} .

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run with parameter options: "
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
