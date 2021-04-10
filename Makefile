GO_LIB_PATH = github.com/nishgowda/docktor/lib
GO_BIN_PATH=github.com/nishgowda/docktor/bin
# tests and builds executable
all: test-all build

# runs fmt on all non fmt formatted go files
# builds executable in bin directory
build:
	- ./fmt.sh 
	- go build -o bin/docktor bin/main.go

# performs all tests on docktor library functions
test-all: test_healthcheck test_healthcheck test_heal test_suggestions \
			test_scan test-server
test-server:
	- go test $(GO_BIN_PATH)/server

test-coverage: 
	- ./cov.sh

test_healthcheck:
	- go test $(GO_LIB_PATH)/healthcheck

test_autoheal:
	- go test $(GO_LIB_PATH)/autoheal

test_heal:
	- go test $(GO_LIB_PATH)/heal

test_suggestions:
	- go test $(GO_LIB_PATH)/suggestions

test_scan:
	- go test $(GO_LIB_PATH)/scan

clean: 
	rm -f bin/docktor


