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

Add linter to your `$PATH` or run it directly:

```
./bin/gocritic package package/path/to/lint # check single package
./bin/gocritic project $GOPATH/src          # check all packages under $GOPATH/src
```

Run `gocritic` without arguments for more info.

Usage of **gocritic**: `gocritic [sub-command] [sub-command args...]`

Examples:

| Command | Description |
| --- | --- |
| `gocritic package fmt` | Runs all (stable) checkers on fmt package |
| `gocritic package -enable elseif,param-name fmt` | Runs specified checkers on fmt package |
| `gocritic project $GOROOT/src` | Run all (stable) checkers on entire GOROOT |

## Contributing

This project aims to be contributing-friendly.

We're using optimistic merging strategy most of the time.
In short, this means that if your contribution has some flaws, we can still merge it and then
fix them by ourselves. Experimental and work-in-progress checkers are isolated, so nothing bad will happen.

Code style is the same as in Go project, see [CodeReviewComments](https://github.com/golang/go/wiki/codereviewcomments).

See [CONTRIBUTING.md](CONTRIBUTING.md) for more details.
It also describes how to develop a new checker for the linter.
