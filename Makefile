phony: build, run


build:
	GOOS=darwin GOARCH=amd64 CGO_CFLAGS="-arch x86_64" CGO_ENABLED=1 go build

run:
	GOOS=darwin GOARCH=amd64 CGO_CFLAGS="-arch x86_64" CGO_ENABLED=1 go run .