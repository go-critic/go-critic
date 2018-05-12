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

## Quick start / installation

```
go get -u github.com/PieselBois/kfulint/...
```

Add linter to your `$PATH` or run it directly:

```
./bin/kfulint package/path/to/lint
```

Run `./bin/kfulint -help` for more info.
