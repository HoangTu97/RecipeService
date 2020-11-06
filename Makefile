build:
	go build -tags=jsoniter .

dev:
	swag init && make build && ./Food
