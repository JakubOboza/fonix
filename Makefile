.PHONY: build test deps build-release package-release release

build:
	go build -o bin/fonix main.go

test:
	go test -v ./... -count=1

deps:
	go get

release: build-release package-release
	@echo "Release build and packaged"

build-release:
	GOOS=darwin  GOARCH=amd64 go build -o release/osx-amd64/fonix main.go
	GOOS=darwin  GOARCH=arm64 go build -o release/osx-arm64/fonix main.go
	GOOS=linux   GOARCH=amd64 go build -o release/linux-amd64/fonix main.go
	GOOS=linux   GOARCH=arm GOARM=5 go build -o release/linux-armpi/fonix main.go

package-release:
	tar -czvf release/fonix.osx-amd64.tar.gz --directory=release/osx-amd64/ fonix
	tar -czvf release/fonix.osx-arm64.tar.gz --directory=release/osx-arm64/ fonix
	tar -czvf release/fonix.linux-amd64.tar.gz --directory=release/linux-amd64/ fonix
	tar -czvf release/fonix.linux-armpi.tar.gz --directory=release/linux-armpi/ fonix