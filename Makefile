BIN_NAME = App

run:
	./${BIN_NAME}

build:
	go build -tags=jsoniter -o ${BIN_NAME}

docs:
	swag init

dev: docs build run
