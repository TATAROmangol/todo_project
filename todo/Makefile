ifeq ($(MAKECMDGOALS),localRun)
   include .env
   export
   POSTGRES_PORT=${OUT_POSTGRES_PORT}
endif

run: build
	./bin/main

localRun: build
	./bin/main

build:
	go build -o ./bin ./cmd/main.go