# Contributing to go-critic

Most of the conventions and rules are derived from [Go](https://github.com/golang/go) project.

Some sections are copied from [ZeroMQ C4](https://rfc.zeromq.org/spec:42/C4/) (Collective Code Construction Contract).

## Getting involved

There are three main ways to contribute:

1. Join [issue tracker](https://github.com/go-critic/go-critic/issues) and help us with
   feature requests, bug reports (including false-positive results), etc.

2. Submit code patches: you can add a new checker, fix or improve already existing checkers
   or make checkers running framework better. This includes patches that improve testing.

3. Improve documentation and/or wiki pages.

The simplest and recommended way to make a first contribution is to fix minor style issue
like typo or missing doc comment. You can also filter issues by using
[help wanted](https://github.com/go-critic/go-critic/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22) and
[good first issue](https://github.com/go-critic/go-critic/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22) labels.

## Code review

### Code review: accepted abbreviations

**LGTM** — looks good to me

**SGTM** — sounds good to me

**PTAL** — please take a look

**s/foo/bar/** — replace foo with bar; this is sed syntax

**s/foo/bar/g** — replace foo with bar throughout your entire change

See [CodeReview](https://github.com/golang/go/wiki/CodeReview).

### Code review: main rules

- Maintainers SHOULD NOT merge their own patches except in exceptional cases, such as non-responsiveness from other Maintainers for an extended period (more than 1-2 days).

- Maintainers SHALL NOT make value judgments on correct patches.

- Maintainers SHALL merge correct patches from other Contributors rapidly.

- Maintainers MAY merge incorrect patches from other Contributors with the goals of (a) ending fruitless discussions, (b) capturing toxic patches in the historical record, (c) engaging with the Contributor on improving their patch quality.

- The user who created an issue SHOULD close the issue after checking the patch is successful.

- Any Contributor who has value judgments on a patch SHOULD express these via their own patches.

## How to write issues
1. Choose correct name to the issue. It should start with path to the checker/folder/doc/etc... it belongs to. Then should be short description of the issue.

Correct naming: 
```
docs/contributing.md: add info about issues.
```

Incorrect naming: 
```
please add some useful content about issues in our contributing.md
```

2. Issue description should contain detailed information about it. If it is bug request, please write steps to reproduce it.
If it is feature request, please describe problem it could solve. Also you could tell us your solution to that problem.

These rules also applies to **pull requests**.

## How to add new checker

1. Come up with checker idea. Concentrate on what it checks.

2. Select one of the base checkers kind (example: expr checker, stmt checker, etc.).

3. Define checker type and constructor function.

4. Add entry to checkers list in `lint.go`. It could be a good idea to mark recently added checker as `experimental`.

5. Add test directory that is named after the checker in `lint/testdata`.

6. Add `positive_tests.go` and `negative_tests.go` files in that directory and add some positive and negative tests there.

7. Run tests. They must fail as your checker does not check anything yet.  
Tests can be run with `go test -v -race -count=1 ./...`.

8. Implement checker itself. Make tests pass.

## Dependencies

These are first-order dependencies:

* [github.com/go-toolsmith/astp](https://github.com/go-toolsmith/astp)
* [github.com/go-toolsmith/astcopy](https://github.com/go-toolsmith/astcopy)
* [github.com/go-toolsmith/astequal](https://github.com/go-toolsmith/astequal)
* [github.com/go-toolsmith/strparse](https://github.com/go-toolsmith/strparse)
* [github.com/go-toolsmith/astfmt](https://github.com/go-toolsmith/astfmt)
* [golang.org/x/tools/go/loader](https://godoc.org/golang.org/x/tools/go/loader)
* [golang.org/x/tools/go/ast/astutil](https://godoc.org/golang.org/x/tools/go/ast/astutil)

If, for whatever reason, you can't build `gocritic` due to missing packages, try to `go get` every one of them.
