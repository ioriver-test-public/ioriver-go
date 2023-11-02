export GOPROXY="https://proxy.golang.org,direct"
export GOPRIVATE="github.com/ioriver-test-public"
export GONOSUMDB="github.com/ioriver-test-public/*"
go mod tidy
go mod vendor
go build ./...
go test ./...