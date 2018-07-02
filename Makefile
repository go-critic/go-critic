.PHONY: test test-checker ci cover tools docs

PKG = $(shell go list ./...)

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
	go tool vet .
	go test -v -race -count=1 ./...
	gometalinter.v2 --skip=testdata --vendor ./...
	gocritic check-project `pwd`

cover:
	@echo "" > coverage.out
	@for d in ${PKG}; \
		do echo "" > profile.out; \
		go test -coverprofile=profile.out -covermode=set $$d; \
		cat profile.out >> coverage.out; \
		rm profile.out; \
	done

tools:
	go get -u gopkg.in/alecthomas/gometalinter.v2
	gometalinter.v2 --install
	go get -u github.com/go-critic/go-critic/...

install:
	go install ./cmd/gocritic
