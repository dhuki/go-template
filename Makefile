mocks:
	go generate ./...

run-http:
	go run ./cmd/api

run-grpc:
	go run ./cmd/grpc

run-http-dev:
	go run ./cmd/api -env=DEV

run-http-uat:
	go run ./cmd/api -env=UAT

run-http-prod:
	go run ./cmd/api -env=PROD

delete_hidden_files:
	find . -type f -name "._*" -print -delete

generate-proto-health:
	protoc -I . --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative internal/adapter/grpc/v1/pb/health_check.proto

