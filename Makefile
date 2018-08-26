.PHONY: test test-checker ci cover tools docs

PKG = github.com/go-critic/go-critic/lint github.com/go-critic/go-critic/lint/internal/astwalk github.com/go-critic/go-critic/lint/internal/lintutil

%:      # stubs to get makefile param for `test-checker` command
	@:	# see: https://stackoverflow.com/a/6273809/433041

test:
	go test -v -count=1 ./...

test-checker:
	go test -v -count=1 -run=/$(filter-out $@,$(MAKECMDGOALS)) ./...

docs:
	cd ./cmd/makedocs && go run main.go

ci:
	go get -t -v ./...
	@if [ "$(TEST_SUITE)" = "linter" ]; then make ci-linter; else make ci-tests; fi

ci-tests:
	go tool vet .
	go test -v -race -count=1 ./...

ci-linter:
	gocritic check-project `pwd`
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b ${GOPATH}/bin v1.9.1
	golangci-lint run -v

cover:
	go get -u github.com/mattn/goveralls
	goveralls -package github.com/go-critic/go-critic/lint -service travis-ci -repotoken ${COVERALLS_TOKEN}

install:
	go install ./cmd/gocritic

new:
	cd ./cmd/makenew && go run main.go -name=$(filter-out $@,$(MAKECMDGOALS))
