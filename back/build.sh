go mod download
go generate ./...
mkdir -p ./../bin
go build  -o ./../bin ./cmd/...