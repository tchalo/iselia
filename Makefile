.PHONY: all deploy

test:
	go test ./...

dep:
	dep ensure

pretty:
	gofmt -w **/**/*.go
	gofmt -w **/*.go
	gofmt -w *.go
