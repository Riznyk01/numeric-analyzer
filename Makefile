build:
	go build -o numan cmd/main.go

buildwin:
	go build -o numan.exe cmd/main.go

test:
	go test ./...