# Project manifest

This document exists to make it clear what fits `gocritic` project and what is not.<br>
In a sense, this manifest describes project philosophy.

This document may change over time.<br>
Contributors and maintainers are encouraged to discuss it and propose changes.<br>

## What gocritic is

It's hard to classify a particular linter due to similarities between them.

There are at least 4 linters-related categories:

1. Ones that you run frequently: on your CI, from your editor or IDE.
   They usually have a false positive rate that is close to zero.
   Warnings they give are not advices, but directives.
   [staticcheck](https://github.com/dominikh/go-tools) is a good example.
2. Rule-based checkers. An example of these is project-local codestyle checkers.
   They are highly customizable and usually check for local conventions
   as opposed to generally accepted guidelines.
   [checkstyle](https://github.com/qiniu/checkstyle) may be close to this category.
3. Linter runners. They typically don't implement checks by themselves, but are
   able to run other linters and aggregate their output.
   Examples: [gometalinter](https://github.com/alecthomas/gometalinter) and [golangci-lint](https://github.com/golangci/golangci-lint).
4. Linters that are made for code audit. Ones that help you to do code review,
   3rd-party library code quality evaluation and so on.
   False positive rate is quite high to make them worthwhile inside CI.
   You don't want to break a build due to the warnings they produce.
   There are advantages of this, see below.
   Example: [gocritic](https://github.com/go-critic/go-critic).

The advantages of not being bound to CI and close-to-zero false positive rate requirement is freedom.
We can find code that looks like a bug or suspicious.
Some checks can't be done without false positives, but they are still useful.
You *have* to check most `gocritic` warnings manually, this is why `quickfix`-like features doesn't
play well with `gocritic`. You may not want to fix issues found by `gocritic` automatically.
