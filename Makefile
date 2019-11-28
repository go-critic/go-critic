.PHONY: test test-checker test-goroot docs ci ci-tidy ci-tests ci-linter cover install

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
	@if [ "$(TEST_SUITE)" = "linter" ]; then make ci-linter; else make ci-tidy; make ci-tests; fi

ci-tidy:
	go mod tidy
	# If you are testing Go 1.11: https://github.com/golang/go/issues/27868#issuecomment-431413621
	go list all > /dev/null
	git diff --exit-code --quiet || (echo "Please run 'go mod tidy' to clean up the 'go.mod' and 'go.sum' files."; false)

ci-tests:
	go test -v -race -count=1 -coverprofile=coverage.out ./...

ci-linter:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b ${GOPATH}/bin v1.21.0
	golangci-lint run -v
	go install github.com/quasilyte/go-consistent
	go-consistent ./...

cover:
	go install github.com/mattn/goveralls
	goveralls -package github.com/go-critic/go-critic/checkers -coverprofile=coverage.out -service travis-ci -repotoken ${COVERALLS_TOKEN}

gocritic:
	lintpack build -o gocritic -linter.version='v0.4.0' -linter.name='gocritic' github.com/go-critic/go-critic/checkers

install:
	go install ./cmd/gocritic
