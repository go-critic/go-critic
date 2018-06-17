.PHONY: test docs

test:
	go test -v -count=1 ./...

docs:
	cd ./cmd/makedocs && go run main.go
