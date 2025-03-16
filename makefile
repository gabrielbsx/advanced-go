.PHONY: default run clean

default: run

run:
	@go run cmd/main.go

clean:
	@rm -rf ./bin