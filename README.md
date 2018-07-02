# go-critic

[![Build Status][travis-image]][travis-url]
[![Go Report Card][go-report-image]][go-report-url]
[![GoDoc][godoc-image]][godoc-url]
[![codecov][codecov-image]][codecov-url]

[travis-image]: https://travis-ci.org/go-critic/go-critic.svg?branch=master
[travis-url]: https://travis-ci.org/go-critic/go-critic
[go-report-image]: https://goreportcard.com/badge/github.com/go-critic/go-critic
[go-report-url]: https://goreportcard.com/report/github.com/go-critic/go-critic
[godoc-image]: https://godoc.org/github.com/go-critic/go-critic/lint?status.svg
[godoc-url]: https://godoc.org/github.com/go-critic/go-critic/lint
[codecov-image]: https://codecov.io/gh/go-critic/go-critic/branch/master/graph/badge.svg
[codecov-url]: https://codecov.io/gh/go-critic/go-critic

Go source code linter providing checks currently missing from other linters.

![Logo](https://avatars1.githubusercontent.com/u/40007520?s=400&u=b44287d8845a63fb0102d5259710c11ea367bb13&v=4)


**Project goals:**

- Provide as many useful checks as possible;
  We're prototyping and experimenting here.

- When a specific check implementation is mature and proven useful,
  propose its integration into other linters.

- If a good checker can't find a better home, it stays here.

We say **yes!** to most checks that feel right and can help someone.

You can use this tool to make whole codebase checks periodically.

There is never too much static code analysis. Try it out.

## Documentation

The latest documentation is available at [go-critic.github.io](https://go-critic.github.io/overview).

## Installation

```
go get -u github.com/go-critic/go-critic/...
```

## Usage

Be sure `gocritic` executable is under your `$PATH`.

Usage of **gocritic**: `gocritic [sub-command] [sub-command args...]`
Run `gocritic` without arguments to get help output.

Examples:

| Command | Description |
| --- | --- |
| `gocritic check-package fmt` | Runs all stable checkers on fmt package |
| `gocritic check-package pkg1 pkg2` | Run all stable checkers on pkg1 and pkg2 |
| `gocritic check-package -enable elseif,param-name fmt` | Runs specified checkers on fmt package |
| `gocritic check-project $GOROOT/src` | Run all stable checkers on entire GOROOT |
| `gocritic check-project $GOPATH/src` | Run all stable checkers on entire GOPATH |
| `gocritic check-project $GOPATH/src/foo` | Run all stable checkers on all packages under GOPATH/src/foo |

> Note: `check-project $GOPATH/xyz` won't work it you're using multiple paths under `GOPATH`.

## Contributing

This project aims to be contribution-friendly.

We're using an optimistic merging strategy most of the time.
In short, this means that if your contribution has some flaws, we can still merge it and then
fix them by ourselves. Experimental and work-in-progress checkers are isolated, so nothing bad will happen.

Code style is the same as in Go project, see [CodeReviewComments](https://github.com/golang/go/wiki/codereviewcomments).

See [CONTRIBUTING.md](CONTRIBUTING.md) for more details.
It also describes how to develop a new checker for the linter.
