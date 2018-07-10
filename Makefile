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
	@if [ "$(TEST_SUITE)" = "gometalinter" ]; then make ci-gometalinter; else make ci-tests; fi

ci-tests:
	go tool vet .
	go test -v -race -count=1 ./...
	gocritic check-project `pwd`

ci-gometalinter:
	go get -u gopkg.in/alecthomas/gometalinter.v2
	gometalinter.v2 --install
	gometalinter.v2 --skip=testdata --vendor ./...

cover:
	go get -u github.com/mattn/goveralls
	@for d in ${PKG}; \
		go test -coverprofile=profile.out -covermode=set $$d; \
		goveralls -coverprofile profile.out -service travis-ci -repotoken ${COVERALLS_TOKEN} \
		rm profile.out; \
	done	

install:
	go install ./cmd/gocritic
