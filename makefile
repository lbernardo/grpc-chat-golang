VERSION=v1

service:
	protoc --proto_path=proto/api/v1  proto/api/v1/*.proto --go_out=plugins=grpc:pkg/api/v1

start-server:
	go run server/main.go

start-client:
	go run client/main.go
