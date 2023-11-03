.PHONY: build
build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./jinan-darwin-amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./jinan-linux-amd64
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./jinan-windows-amd64.exe
