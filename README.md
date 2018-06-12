# go-critic

[![Build Status](https://travis-ci.org/go-critic/go-critic.svg?branch=master)](https://travis-ci.org/go-critic/go-critic)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-critic/go-critic)](https://goreportcard.com/report/github.com/go-critic/go-critic)

Go source code linter that brings checks that are currently not implemented in other linters.

![Logo](https://avatars1.githubusercontent.com/u/40007520?s=400&u=b44287d8845a63fb0102d5259710c11ea367bb13&v=4)


**Project goals:**

- Provide as much useful checks as possible.
  We're prototyping and experimenting here

- When specific check implementation is mature and proven useful,
  propose it's integration into other linter

- If good checker can't find a better home, it stays here

We say **yes!** to most checks that feel right and can help someone.

You can use this tool to make whole codebase checks from time to time.

There is never too much static code analysis. Try it out.

## Documentation

Latest documentation is available at [go-critic.github.io](https://go-critic.github.io/overview).

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

## Contributing

This project aims to be contributing-friendly.

We're using optimistic merging strategy most of the time.
In short, this means that if your contribution has some flaws, we can still merge it and then
fix them by ourselves. Experimental and work-in-progress checkers are isolated, so nothing bad will happen.

Code style is the same as in Go project, see [CodeReviewComments](https://github.com/golang/go/wiki/codereviewcomments).

See [CONTRIBUTING.md](CONTRIBUTING.md) for more details.
It also describes how to develop a new checker for the linter.

## Dependencies

* [github.com/go-toolsmith/astp](https://github.com/go-toolsmith/astp)
* [github.com/go-toolsmith/astcopy](https://github.com/go-toolsmith/astcopy)
* [github.com/go-toolsmith/astequal](https://github.com/go-toolsmith/astequal)
* [github.com/go-toolsmith/strparse](https://github.com/go-toolsmith/strparse)
* [golang.org/x/tools/go/loader](https://godoc.org/golang.org/x/tools/go/loader)
* [golang.org/x/tools/go/ast/astutil](https://godoc.org/golang.org/x/tools/go/ast/astutil)
