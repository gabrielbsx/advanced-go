.PHONY: default run clean

default: run

run:
	@go run cmd/main.go
	
example:
	@go run cmd/example.go

clean:
	@rm -rf ./bin