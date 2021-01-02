build:
	go build -tags=jsoniter -o App

dev:
	swag init && make build && ./App
