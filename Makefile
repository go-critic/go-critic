.PHONY: test test-checker ci cover tools docs

export GO111MODULE := on

%:      # stubs to get makefile param for `test-checker` command
	@:	# see: https://stackoverflow.com/a/6273809/433041

test:
	go test -v -count=1 ./...

test-checker:
	go test -v -count=1 -run=/$(filter-out $@,$(MAKECMDGOALS)) ./...

test-goroot:
	go run cmd/gocritic/main.go check-project -enable=$(filter-out $@,$(MAKECMDGOALS)) ${GOROOT}/src

docs:
	cd ./cmd/makedocs && go run main.go

ci:
	@if [ "$(TEST_SUITE)" = "linter" ]; then make ci-linter; else make ci-tests; fi

ci-tests:
	go get -u github.com/go-lintpack/lintpack/...
	go test -v -race -count=1 -coverprofile=coverage.out ./...

ci-linter:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b ${GOPATH}/bin v1.12.3
	golangci-lint run -v
	go get -u github.com/Quasilyte/go-consistent
	go-consistent ./...

cover:
	go get -u github.com/mattn/goveralls
	goveralls -package github.com/go-critic/go-critic/checkers -coverprofile=coverage.out -service travis-ci -repotoken ${COVERALLS_TOKEN}

install:
	go install ./cmd/gocritic
