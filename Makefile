run: build
	@./bin/app
build:
	@go build -o ./bin/app ./*.go
test:
	@go test -v ./tests/*.go
