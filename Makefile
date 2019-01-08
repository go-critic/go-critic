.PHONY: test test-checker ci cover tools docs

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
	GO111MODULE=on go mod vendor
	@if [ "$(TEST_SUITE)" = "linter" ]; then make ci-linter; else make ci-tests; fi

ci-tests:
	go test -v -race -count=1 ./...

ci-linter:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b ${GOPATH}/bin v1.12.3
	golangci-lint run -v

cover:
	GO111MODULE=on go mod vendor
	go get -u github.com/mattn/goveralls
	goveralls -package github.com/go-critic/go-critic/checkers -service travis-ci -repotoken ${COVERALLS_TOKEN}

install:
	go install ./cmd/gocritic
