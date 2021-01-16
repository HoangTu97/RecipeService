BIN_NAME = App
BIN_PATH = ./build/${BIN_NAME}

.PHONY: build

dep:
	go get -u github.com/swaggo/swag/cmd/swag

run:
	${BIN_PATH}

build:
	go build -tags=jsoniter -o ${BIN_PATH}

docs:
	swag init

dev: docs build run
