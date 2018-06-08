# go-critic

[![Build Status](https://travis-ci.org/go-critic/go-critic.svg?branch=master)](https://travis-ci.org/go-critic/go-critic)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-critic/go-critic)](https://goreportcard.com/report/github.com/go-critic/go-critic)

Go source code linter that brings checks that are currently not implemented in other linters.

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

Latest documentation is available at [go-critic.github.io](https://go-critic.github.io/overview.html).

## Installation

```
go get -u github.com/go-critic/go-critic/...
```

## Usage

Add linter to your `$PATH` or run it directly:

```
./bin/gocritic package/path/to/lint
```

Run `gocritic -help` for more information.

Usage of **gocritic**: `gocritic [flags] [package]`

Examples:

| Command | Description |
| --- | --- |
| `gocritic fmt` | Runs all checkers on fmt package |
| `gocritic -enable elseif,param-name fmt` | Runs specified checkers on package |

## Contributing

This project aims to be contributing-friendly.

We're using optimistic merging strategy most of the time.
In short, this means that if your contribution has some flaws, we can still merge it and then
fix them by ourselves. Experimental and work-in-progress checkers are isolated, so nothing bad will happen.

Code style is the same as in Go project, see [CodeReviewComments](https://github.com/golang/go/wiki/codereviewcomments).

See [CONTRIBUTING.md](CONTRIBUTING.md) for more details.
It also describes how to develop a new checker for the linter.
