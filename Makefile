setup:
	go clean --modcache
	go mod tidy -compat=1.17

coverage:
	go mod vendor
	go test ./... -coverprofile=coverage.out -covermode=count -mod=vendor
	go tool cover -func=coverage.out

build:
	go build cmd/main.go