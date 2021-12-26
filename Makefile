.PHONY: test test-checker test-goroot docs ci ci-tidy ci-tests ci-linter cover install

GOPATH_DIR=`go env GOPATH`

export GO111MODULE := on

%:      # stubs to get makefile param for `test-checker` command
	@:	# see: https://stackoverflow.com/a/6273809/433041

build-release:
	mkdir -p bin
	go build -o bin/gocritic -ldflags "-X 'main.Version=${GOCRITIC_VERSION}'" ./cmd/gocritic

test:
	go test -v -count=1 ./...

test-checker:
	go test -v -count=1 -run=/$(filter-out $@,$(MAKECMDGOALS)) ./checkers

test-goroot:
	go run cmd/gocritic/main.go check-project -enable=$(filter-out $@,$(MAKECMDGOALS)) ${GOROOT}/src

docs:
	cd ./cmd/makedocs && go run main.go

ci:
	@if [ "$(TEST_SUITE)" = "linter" ]; then make ci-linter; else make ci-tidy; make ci-tests; fi

ci-tidy:
	go mod tidy
	# If you are testing Go 1.11: https://github.com/golang/go/issues/27868#issuecomment-431413621
	go list all > /dev/null
	git diff --exit-code --quiet || (echo "Please run 'go mod tidy' to clean up the 'go.mod' and 'go.sum' files."; false)

ci-tests:
	GOCRITIC_EXTERNAL_TESTS=1 go test -v -race -count=1 -coverprofile=coverage.out ./...

ci-linter:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH_DIR)/bin v1.30.0
	@$(GOPATH_DIR)/bin/golangci-lint run -v
	cd tools && go install github.com/quasilyte/go-consistent
	@$(GOPATH_DIR)/bin/go-consistent ./...
	go build -o gocritic ./cmd/gocritic
	./gocritic check -enableAll ./...

cover:
	cd tools && go install github.com/mattn/goveralls
	goveralls -package github.com/go-critic/go-critic/checkers -coverprofile=coverage.out -service travis-ci -repotoken ${COVERALLS_TOKEN}

gocritic:
	go build -o gocritic ./cmd/gocritic

install:
	go install ./cmd/gocritic
