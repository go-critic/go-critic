# kfulint

[![Build Status](https://travis-ci.org/PieselBois/kfulint.svg?branch=master)](https://travis-ci.org/PieselBois/kfulint)
[![Go Report Card](https://goreportcard.com/badge/github.com/PieselBois/kfulint)](https://goreportcard.com/report/github.com/PieselBois/kfulint)

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

Latest documentation is available in [docs/overview.md](docs/overview.md).

## Installation

```
go get -u github.com/PieselBois/kfulint/...
```

## Usage

Add linter to your `$PATH` or run it directly:

```
./bin/kfulint package/path/to/lint
```

Run `kfulint -help` for more information.

Usage of **kfulint**: `kfulint [flags] [package]`

Examples:  
&nbsp; | &nbsp;
------ | -----
`kfulint fmt` | Runs all checkers on fmt package.
`kfulint -enable elseif,param-name fmt` | Runs specified checkers on package.

## Contributing

This project aims to be contributing-friendly.

We're using optimistic merging strategy most of the time.
In short, this means that if your contribution has some flaws, we can still merge it and then
fix them by ourselves. Experimental and work-in-progress checkers are isolated, so nothing bad will happen.

Code style is the same as in Go project, see [CodeReviewComments](https://github.com/golang/go/wiki/codereviewcomments).

See [CONTRIBUTING.md](CONTRIBUTING.md) for more details.
It also describes how to develop a new checker for the linter.
