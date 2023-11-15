# go-critic

![Build Status](https://github.com/go-critic/go-critic/workflows/Go/badge.svg)
[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/avelino/awesome-go#code-analysis)
[![Go Report Card][go-report-image]][go-report-url]
[![coverage][coverage-image]][coverage-url]
[![PRs Welcome][pr-welcome-image]][pr-welcome-url]

[go-report-image]: https://goreportcard.com/badge/github.com/go-critic/go-critic
[go-report-url]: https://goreportcard.com/report/github.com/go-critic/go-critic
[coverage-image]: https://coveralls.io/repos/github/go-critic/go-critic/badge.svg?branch=master
[coverage-url]: https://coveralls.io/github/go-critic/go-critic?branch=master
[pr-welcome-image]: https://img.shields.io/badge/PRs-welcome-brightgreen.svg
[pr-welcome-url]: https://github.com/go-critic/go-critic/blob/master/CONTRIBUTING.md

Highly extensible Go source code linter providing checks currently missing from other linters.

![Logo](https://avatars1.githubusercontent.com/u/40007520?s=300&u=b44287d8845a63fb0102d5259710c11ea367bb13&v=4)

There is never too much static code analysis. Try it out.

## Features

* Almost 100 diagnostics that check for [bugs](https://go-critic.github.io/overview#checkers-from-the-diagnostic-group), [performance](https://go-critic.github.io/overview#checkers-from-the-performance-group) and [style](https://go-critic.github.io/overview#checkers-from-the-style-group) issues
* Extensible without re-compilation with [dynamic rules](https://quasilyte.dev/blog/post/ruleguard/)
* Includes `#opinionated` checks with very strict and specific requirements
* Self-documented: `gocritic doc <checkname>` gives a checker description

## Documentation

The latest documentation is available at [go-critic.com](https://go-critic.com/overview).

## Installation

For most users, using `go-critic` under [golangci-lint](https://github.com/golangci/golangci-lint) is enough.

Precompiled `go-critic` binaries can be found at [releases](https://github.com/go-critic/go-critic/releases) page.

It can be installed in the usual Go way by running:
```bash
go install -v github.com/go-critic/go-critic/cmd/gocritic@latest
```

To build `go-critic` from sources, clone this repository and run `make gocritic`.

On macOS, you can also install `go-critic` using [MacPorts](https://www.macports.org): `sudo port install go-critic`

## Usage

Be sure `gocritic` executable is under your `$PATH`.

Usage of **gocritic**: `gocritic [sub-command] [sub-command args...]`
Run `gocritic` without arguments to get help output.

```
Supported sub-commands:
	check - run linter over specified targets
		$ gocritic check -help
		$ gocritic check -v -enable='paramTypeCombine,unslice' strings bytes
		$ gocritic check -v -enable='#diagnostic' -disable='#experimental,#opinionated' ./...
	version - print linter version
		$ gocritic version
	doc - get installed checkers documentation
		$ gocritic doc -help
		$ gocritic doc
		$ gocritic doc checkerName
```

`check` sub-command examples:

```bash
# Runs all stable checkers on `fmt` package:
gocritic check fmt

# Run all stable checkers on `pkg1` and `pkg2`
gocritic check pkg1 pkg2

# Run all stable checkers on `fmt` package and configure rangeExprCopy checker
gocritic check -@rangeExprCopy.sizeThreshold 128 fmt

# Runs specified checkers on `fmt` package:
gocritic check -enable elseif,paramName fmt

# Run all stable checkers on current dir and all its children recursively:
gocritic check ./...

# Like above, but without `appendAssign` check:
gocritic check -disable=appendAssign ./...

# Run all stable checkers on `foo.go` file:
gocritic check foo.go

# Run stable diagnostics over `strings` package:
gocritic check -enable='#diagnostic' -disable='#experimental' strings

# Run all stable and non-opinionated checks:
gocritic check -enableAll -disable='#experimental,#opinionated' ./src/...
```

> To get a list of available checker parameters, run `gocritic doc <checkerName>`.

In place of a single name, **tag** can be used. Tag is a named checkers group.

Tags:
* `#diagnostic` - kind of checks that detect various errors in code
* `#style` - kind of checks that find style issues in code
* `#performance` - kind of checks that detect potential performance issues in code
* `#experimental` - check is under testing and development. Disabled by default
* `#opinionated` - check can be unwanted for some people. Disabled by default
* `#security` -  kind of checks that find security issues in code. Disabled by default and empty, so will fail if enabled.

## Contributing

This project aims to be contribution-friendly.

Our chats: [English](https://t.me/go_critic_eng) or
[Russian](https://t.me/go_critic_ru)
([Telegram website](https://telegram.org/))

We're using an optimistic merging strategy most of the time.
In short, this means that if your contribution has some flaws, we can still merge it and then
fix them by ourselves. Experimental and work-in-progress checkers are isolated, so nothing bad will happen.

Code style is the same as in Go project, see [CodeReviewComments](https://github.com/golang/go/wiki/codereviewcomments).

See [CONTRIBUTING.md](CONTRIBUTING.md) for more details.
It also describes how to develop a new checker for the linter.
