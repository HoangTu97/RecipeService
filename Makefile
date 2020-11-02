dev:
	swag init && go build -tags=jsoniter . && ./Food