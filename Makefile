GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	mkdir -p api/gtb/v1/gtb && rm -rf api/gtb/v1/gtb/*
	protoc --proto_path=./api/gtb/v1 \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api/gtb/v1/gtb \
 	       --go-http_out=paths=source_relative:./api/gtb/v1/gtb \
 	       --go-grpc_out=paths=source_relative:./api/gtb/v1/gtb \
	       $(API_PROTO_FILES)

.PHONY: swagger
# generate swagger file
swagger:
	protoc --proto_path=. \
		   --proto_path=./third_party \
		   --openapiv2_out . \
		   --openapiv2_opt logtostderr=true \
		   --openapiv2_opt json_names_for_fields=false \
		   $(API_PROTO_FILES)

.PHONY: generate
# generate client code
generate:
	go generate ./...

.PHONY: wire
# generate wire code
wire:
	make wire-tbw

.PHONY: wire-game
# generate wire code
wire-tbw:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	cd cmd/tbw && wire


.PHONY: build
# build
build:
	make wire
	mkdir -p bin/ && go build -ldflags "-s -w -X main.Version=$(VERSION)" -o ./bin/ ./...
	make upx

.PHONY: build-tbw
# -outer
build-tbw:
	make wire-tbw && mkdir -p bin/ && go build -ldflags "-s -w" -o ./bin/tbw ./cmd/tbw && upx bin/tbw

.PHONY: upx
# build
upx:
	upx bin/tbw

.PHONY: test
# test
test:
	go test -v ./... -cover


.PHONY: all
# generate all
all:
	make api;
	make config;
	make wire;
	make generate;
	make validate;
	make swagger;
	make build;
	make test;


.PHONY: run-tbw
# run-tbw
run-tbw:
	make build
	./bin/tbw -conf ./configs


.PHONY: validate
# generate validate code
validate:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--proto_path=$(VALIDATE) \
		--validate_out="lang=go",paths=source_relative:. \
        $(API_PROTO_FILES)

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.PHONY: test-cover
# test coverage
test-cover:
	@go test -v $(go list ./... | grep "service\|internal") -coverprofile=profile.cov ./...
	@echo 'project test coverage:'
	@go tool cover -func profile.cov | tail -n 1 | awk '{print $$3}' | awk '{ gsub(/%/,""); print $$0 }'
	@rm -f ./profile.cov

.DEFAULT_GOAL := help
