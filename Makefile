APP_NAME=example
APP_NAME_LINUX=$(APP_NAME)_linux

# go
go_fmt:
	gofmt -l -w .

go_run:
	go run cmd/main.go

go_build:
	CGO_ENABLED=0 go build -ldflags "-s -w" -o target/$(APP_NAME) cmd/main.go

go_build_linux:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags "-s -w" -o target/$(APP_NAME_LINUX) cmd/main.go

go_test: go_fmt
	CONF_NAME=example-test go test ./...

