GO_PATH = github.com/nishgowda/docktor/lib

test:
	- go test $(GO_PATH)/healthcheck
	- go test $(GO_PATH)/autoheal
	- go test $(GO_PATH)/heal

build:
	go build -o bin/docktor bin/main.go
