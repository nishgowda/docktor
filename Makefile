GO_PATH = github.com/nishgowda/docktor/lib

# tests and builds executable
all: test build

build:
	go build -o bin/docktor bin/main.go

# performs specified tests on docktor library functions
test:
	- go test $(GO_PATH)/healthcheck
	- go test $(GO_PATH)/autoheal
	- go test $(GO_PATH)/heal


