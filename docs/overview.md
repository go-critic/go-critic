## Checks overview

This page describes checks supported by [go-critic](https://github.com/go-critic/go-critic) linter.

[//]: # (This is generated file, please don't edit it yourself.)

## Checkers

Total number of checks is 105 :rocket:

* :heavy_check_mark: checker is enabled by default.
* :white_check_mark: checker is disabled by default.

### Checkers from the "diagnostic" group

Diagnostics try to find programming errors in the code.
They also detect code that may be correct, but looks suspicious.

> All diagnostics are enabled by default (unless it has "experimental" tag).

| Name | Short description |
|------|-------------------|
|:heavy_check_mark:[appendAssign](#appendassign)|Detects suspicious append result assignments|
|:heavy_check_mark:[argOrder](#argorder)|Detects suspicious arguments order|
|:heavy_check_mark:[badCall](#badcall)|Detects suspicious function calls|
|:heavy_check_mark:[badCond](#badcond)|Detects suspicious condition expressions|
|:white_check_mark:[badLock](#badlock)|Detects suspicious mutex lock/unlock operations|
|:white_check_mark:[badRegexp](#badregexp)|Detects suspicious regexp patterns|
|:white_check_mark:[badSorting](#badsorting)|Detects bad usage of sort package|
|:white_check_mark:[badSyncOnceFunc](#badsynconcefunc)|Detects bad usage of sync.OnceFunc|
|:white_check_mark:[builtinShadowDecl](#builtinshadowdecl)|Detects top-level declarations that shadow the predeclared identifiers|
|:heavy_check_mark:[caseOrder](#caseorder)|Detects erroneous case order inside switch statements|
|:heavy_check_mark:[codegenComment](#codegencomment)|Detects malformed 'code generated' file comments|
|:white_check_mark:[commentedOutCode](#commentedoutcode)|Detects commented-out code inside function bodies|
|:white_check_mark:[deferInLoop](#deferinloop)|Detects loops inside functions that use defer|
|:heavy_check_mark:[deprecatedComment](#deprecatedcomment)|Detects malformed 'deprecated' doc-comments|
|:heavy_check_mark:[dupArg](#duparg)|Detects suspicious duplicated arguments|
|:heavy_check_mark:[dupBranchBody](#dupbranchbody)|Detects duplicated branch bodies inside conditional statements|
|:heavy_check_mark:[dupCase](#dupcase)|Detects duplicated case clauses inside switch or select statements|
|:heavy_check_mark:[dupSubExpr](#dupsubexpr)|Detects suspicious duplicated sub-expressions|
|:white_check_mark:[dynamicFmtString](#dynamicfmtstring)|Detects suspicious formatting strings usage|
|:white_check_mark:[emptyDecl](#emptydecl)|Detects suspicious empty declarations blocks|
|:white_check_mark:[evalOrder](#evalorder)|Detects unwanted dependencies on the evaluation order|
|:heavy_check_mark:[exitAfterDefer](#exitafterdefer)|Detects calls to exit/fatal inside functions that use defer|
|:white_check_mark:[externalErrorReassign](#externalerrorreassign)|Detects suspicious reassignment of error from another package|
|:white_check_mark:[filepathJoin](#filepathjoin)|Detects problems in filepath.Join() function calls|
|:heavy_check_mark:[flagDeref](#flagderef)|Detects immediate dereferencing of `flag` package pointers|
|:heavy_check_mark:[flagName](#flagname)|Detects suspicious flag names|
|:heavy_check_mark:[mapKey](#mapkey)|Detects suspicious map literal keys|
|:white_check_mark:[nilValReturn](#nilvalreturn)|Detects return statements those results evaluate to nil|
|:heavy_check_mark:[offBy1](#offby1)|Detects various off-by-one kind of errors|
|:white_check_mark:[regexpPattern](#regexppattern)|Detects suspicious regexp patterns|
|:white_check_mark:[returnAfterHttpError](#returnafterhttperror)|Detects suspicious http.Error call without following return|
|:heavy_check_mark:[sloppyLen](#sloppylen)|Detects usage of `len` when result is obvious or doesn't make sense|
|:white_check_mark:[sloppyReassign](#sloppyreassign)|Detects suspicious/confusing re-assignments|
|:heavy_check_mark:[sloppyTypeAssert](#sloppytypeassert)|Detects redundant type assertions|
|:white_check_mark:[sortSlice](#sortslice)|Detects suspicious sort.Slice calls|
|:white_check_mark:[sprintfQuotedString](#sprintfquotedstring)|Detects "%s" formatting directives that can be replaced with %q|
|:white_check_mark:[sqlQuery](#sqlquery)|Detects issue in Query() and Exec() calls|
|:white_check_mark:[syncMapLoadAndDelete](#syncmaploadanddelete)|Detects sync.Map load+delete operations that can be replaced with LoadAndDelete|
|:white_check_mark:[truncateCmp](#truncatecmp)|Detects potential truncation issues when comparing ints of different sizes|
|:white_check_mark:[uncheckedInlineErr](#uncheckedinlineerr)|Detects unchecked errors in if statements|
|:white_check_mark:[unnecessaryDefer](#unnecessarydefer)|Detects redundantly deferred calls|
|:white_check_mark:[weakCond](#weakcond)|Detects conditions that are unsafe due to not being exhaustive|

### Checkers from the "style" group

Style checks suggest replacing some form of expression/statement
with another one that is considered more idiomatic or simple.

> Only non-opinionated style checks are enabled by default.

| Name | Short description |
|------|-------------------|
|:heavy_check_mark:[assignOp](#assignop)|Detects assignments that can be simplified by using assignment operators|
|:white_check_mark:[boolExprSimplify](#boolexprsimplify)|Detects bool expressions that can be simplified|
|:white_check_mark:[builtinShadow](#builtinshadow)|Detects when predeclared identifiers are shadowed in assignments|
|:heavy_check_mark:[captLocal](#captlocal)|Detects capitalized names for local variables|
|:heavy_check_mark:[commentFormatting](#commentformatting)|Detects comments with non-idiomatic formatting|
|:white_check_mark:[commentedOutImport](#commentedoutimport)|Detects commented-out imports|
|:heavy_check_mark:[defaultCaseOrder](#defaultcaseorder)|Detects when default case in switch isn't on 1st or last position|
|:white_check_mark:[deferUnlambda](#deferunlambda)|Detects deferred function literals that can be simplified|
|:white_check_mark:[docStub](#docstub)|Detects comments that silence go lint complaints about doc-comment|
|:white_check_mark:[dupImport](#dupimport)|Detects multiple imports of the same package under different aliases|
|:heavy_check_mark:[elseif](#elseif)|Detects else with nested if statement that can be replaced with else-if|
|:white_check_mark:[emptyFallthrough](#emptyfallthrough)|Detects fallthrough that can be avoided by using multi case values|
|:white_check_mark:[emptyStringTest](#emptystringtest)|Detects empty string checks that can be written more idiomatically|
|:white_check_mark:[exposedSyncMutex](#exposedsyncmutex)|Detects exposed methods from sync.Mutex and sync.RWMutex|
|:white_check_mark:[hexLiteral](#hexliteral)|Detects hex literals that have mixed case letter digits|
|:white_check_mark:[httpNoBody](#httpnobody)|Detects nil usages in http.NewRequest calls, suggesting http.NoBody as an alternative|
|:heavy_check_mark:[ifElseChain](#ifelsechain)|Detects repeated if-else statements and suggests to replace them with switch statement|
|:white_check_mark:[importShadow](#importshadow)|Detects when imported package names shadowed in the assignments|
|:white_check_mark:[initClause](#initclause)|Detects non-assignment statements inside if/switch init clause|
|:white_check_mark:[methodExprCall](#methodexprcall)|Detects method expression call that can be replaced with a method call|
|:white_check_mark:[nestingReduce](#nestingreduce)|Finds where nesting level could be reduced|
|:heavy_check_mark:[newDeref](#newderef)|Detects immediate dereferencing of `new` expressions|
|:white_check_mark:[octalLiteral](#octalliteral)|Detects old-style octal literals|
|:white_check_mark:[paramTypeCombine](#paramtypecombine)|Detects if function parameters could be combined by type and suggest the way to do it|
|:white_check_mark:[preferFilepathJoin](#preferfilepathjoin)|Detects concatenation with os.PathSeparator which can be replaced with filepath.Join|
|:white_check_mark:[ptrToRefParam](#ptrtorefparam)|Detects input and output parameters that have a type of pointer to referential type|
|:white_check_mark:[redundantSprint](#redundantsprint)|Detects redundant fmt.Sprint calls|
|:heavy_check_mark:[regexpMust](#regexpmust)|Detects `regexp.Compile*` that can be replaced with `regexp.MustCompile*`|
|:white_check_mark:[regexpSimplify](#regexpsimplify)|Detects regexp patterns that can be simplified|
|:white_check_mark:[ruleguard](#ruleguard)|Runs user-defined rules using ruleguard linter|
|:heavy_check_mark:[singleCaseSwitch](#singlecaseswitch)|Detects switch statements that could be better written as if statement|
|:white_check_mark:[stringConcatSimplify](#stringconcatsimplify)|Detects string concat operations that can be simplified|
|:white_check_mark:[stringsCompare](#stringscompare)|Detects strings.Compare usage|
|:heavy_check_mark:[switchTrue](#switchtrue)|Detects switch-over-bool statements that use explicit `true` tag value|
|:white_check_mark:[timeExprSimplify](#timeexprsimplify)|Detects manual conversion to milli- or microseconds|
|:white_check_mark:[todoCommentWithoutDetail](#todocommentwithoutdetail)|Detects TODO comments without detail/assignee|
|:white_check_mark:[tooManyResultsChecker](#toomanyresultschecker)|Detects function with too many results|
|:white_check_mark:[typeAssertChain](#typeassertchain)|Detects repeated type assertions and suggests to replace them with type switch statement|
|:white_check_mark:[typeDefFirst](#typedeffirst)|Detects method declarations preceding the type definition itself|
|:heavy_check_mark:[typeSwitchVar](#typeswitchvar)|Detects type switches that can benefit from type guard clause with variable|
|:white_check_mark:[typeUnparen](#typeunparen)|Detects unneeded parenthesis inside type expressions and suggests to remove them|
|:heavy_check_mark:[underef](#underef)|Detects dereference expressions that can be omitted|
|:white_check_mark:[unlabelStmt](#unlabelstmt)|Detects redundant statement labels|
|:heavy_check_mark:[unlambda](#unlambda)|Detects function literals that can be simplified|
|:white_check_mark:[unnamedResult](#unnamedresult)|Detects unnamed results that may benefit from names|
|:white_check_mark:[unnecessaryBlock](#unnecessaryblock)|Detects unnecessary braced statement blocks|
|:heavy_check_mark:[unslice](#unslice)|Detects slice expressions that can be simplified to sliced expression itself|
|:heavy_check_mark:[valSwap](#valswap)|Detects value swapping code that are not using parallel assignment|
|:white_check_mark:[whyNoLint](#whynolint)|Ensures that `//nolint` comments include an explanation|
|:heavy_check_mark:[wrapperFunc](#wrapperfunc)|Detects function calls that can be replaced with convenience wrappers|
|:white_check_mark:[yodaStyleExpr](#yodastyleexpr)|Detects Yoda style expressions and suggests to replace them|

### Checkers from the "performance" group

Performance checks tell you about potential issues that
can make your code run slower than it could be.

> All performance checks are disabled by default.

| Name | Short description |
|------|-------------------|
|:white_check_mark:[appendCombine](#appendcombine)|Detects `append` chains to the same slice that can be done in a single `append` call|
|:white_check_mark:[equalFold](#equalfold)|Detects unoptimal strings/bytes case-insensitive comparison|
|:white_check_mark:[hugeParam](#hugeparam)|Detects params that incur excessive amount of copying|
|:white_check_mark:[indexAlloc](#indexalloc)|Detects strings.Index calls that may cause unwanted allocs|
|:white_check_mark:[preferDecodeRune](#preferdecoderune)|Detects expressions like []rune(s)[0] that may cause unwanted rune slice allocation|
|:white_check_mark:[preferFprint](#preferfprint)|Detects fmt.Sprint(f/ln) calls which can be replaced with fmt.Fprint(f/ln)|
|:white_check_mark:[preferStringWriter](#preferstringwriter)|Detects w.Write or io.WriteString calls which can be replaced with w.WriteString|
|:white_check_mark:[preferWriteByte](#preferwritebyte)|Detects WriteRune calls with rune literal argument that is single byte and reports to use WriteByte instead|
|:white_check_mark:[rangeExprCopy](#rangeexprcopy)|Detects expensive copies of `for` loop range expressions|
|:white_check_mark:[rangeValCopy](#rangevalcopy)|Detects loops that copy big objects during each iteration|
|:white_check_mark:[sliceClear](#sliceclear)|Detects slice clear loops, suggests an idiom that is recognized by the Go compiler|
|:white_check_mark:[stringXbytes](#stringxbytes)|Detects redundant conversions between string and []byte|

## appendAssign

[
  **diagnostic** ]

Detects suspicious append result assignments.





**Before:**
```go
p.positives = append(p.negatives, x)
p.negatives = append(p.negatives, y)
```

**After:**
```go
p.positives = append(p.positives, x)
p.negatives = append(p.negatives, y)
```


## appendCombine

[
  **performance** ]

Detects `append` chains to the same slice that can be done in a single `append` call.





**Before:**
```go
xs = append(xs, 1)
xs = append(xs, 2)
```

**After:**
```go
xs = append(xs, 1, 2)
```


## argOrder

[
  **diagnostic** ]

Detects suspicious arguments order.





**Before:**
```go
strings.HasPrefix("#", userpass)
```

**After:**
```go
strings.HasPrefix(userpass, "#")
```


## assignOp

[
  **style** ]

Detects assignments that can be simplified by using assignment operators.





**Before:**
```go
x = x * 2
```

**After:**
```go
x *= 2
```


## badCall

[
  **diagnostic** ]

Detects suspicious function calls.





**Before:**
```go
strings.Replace(s, from, to, 0)
```

**After:**
```go
strings.Replace(s, from, to, -1)
```


## badCond

[
  **diagnostic** ]

Detects suspicious condition expressions.





**Before:**
```go
for i := 0; i > n; i++ {
	xs[i] = 0
}
```

**After:**
```go
for i := 0; i < n; i++ {
	xs[i] = 0
}
```


## badLock

[
  **diagnostic**
  **experimental** ]

Detects suspicious mutex lock/unlock operations.





**Before:**
```go
mu.Lock(); mu.Unlock()
```

**After:**
```go
mu.Lock(); defer mu.Unlock()
```


## badRegexp

[
  **diagnostic**
  **experimental** ]

Detects suspicious regexp patterns.





**Before:**
```go
regexp.MustCompile(`(?:^aa|bb|cc)foo[aba]`)
```

**After:**
```go
regexp.MustCompile(`^(?:aa|bb|cc)foo[ab]`)
```


## badSorting

[
  **diagnostic**
  **experimental** ]

Detects bad usage of sort package.





**Before:**
```go
xs = sort.StringSlice(xs)
```

**After:**
```go
sort.Strings(xs)
```


## badSyncOnceFunc

[
  **diagnostic**
  **experimental** ]

Detects bad usage of sync.OnceFunc.





**Before:**
```go
sync.OnceFunc(foo)()
```

**After:**
```go
fooOnce := sync.OnceFunc(foo); ...; fooOnce()
```


## boolExprSimplify

[
  **style**
  **experimental** ]

Detects bool expressions that can be simplified.





**Before:**
```go
a := !(elapsed >= expectElapsedMin)
b := !(x) == !(y)
```

**After:**
```go
a := elapsed < expectElapsedMin
b := (x) == (y)
```


## builtinShadow

[
  **style**
  **opinionated** ]

Detects when predeclared identifiers are shadowed in assignments.





**Before:**
```go
len := 10
```

**After:**
```go
length := 10
```


## builtinShadowDecl

[
  **diagnostic**
  **experimental** ]

Detects top-level declarations that shadow the predeclared identifiers.





**Before:**
```go
type int struct {}
```

**After:**
```go
type myInt struct {}
```


## captLocal

[
  **style** ]

Detects capitalized names for local variables.





**Before:**
```go
func f(IN int, OUT *int) (ERR error) {}
```

**After:**
```go
func f(in int, out *int) (err error) {}
```


Checker parameters:
<ul>
<li>

  `@captLocal.paramsOnly` whether to restrict checker to params only (default true)

</li>

</ul>

## caseOrder

[
  **diagnostic** ]

Detects erroneous case order inside switch statements.





**Before:**
```go
switch x.(type) {
case ast.Expr:
	fmt.Println("expr")
case *ast.BasicLit:
	fmt.Println("basic lit") // Never executed
}
```

**After:**
```go
switch x.(type) {
case *ast.BasicLit:
	fmt.Println("basic lit") // Now reachable
case ast.Expr:
	fmt.Println("expr")
}
```


## codegenComment

[
  **diagnostic** ]

Detects malformed 'code generated' file comments.





**Before:**
```go
// This file was automatically generated by foogen
```

**After:**
```go
// Code generated by foogen. DO NOT EDIT.
```


## commentFormatting

[
  **style** ]

Detects comments with non-idiomatic formatting.





**Before:**
```go
//This is a comment
```

**After:**
```go
// This is a comment
```


## commentedOutCode

[
  **diagnostic**
  **experimental** ]

Detects commented-out code inside function bodies.





**Before:**
```go
// fmt.Println("Debugging hard")
foo(1, 2)
```

**After:**
```go
foo(1, 2)
```


Checker parameters:
<ul>
<li>

  `@commentedOutCode.minLength` min length of the comment that triggers a warning (default 15)

</li>

</ul>

## commentedOutImport

[
  **style**
  **experimental** ]

Detects commented-out imports.





**Before:**
```go
import (
	"fmt"
	//"os"
)
```

**After:**
```go
import (
	"fmt"
)
```


## defaultCaseOrder

[
  **style** ]

Detects when default case in switch isn't on 1st or last position.





**Before:**
```go
switch {
case x > y:
	// ...
default: // <- not the best position
	// ...
case x == 10:
	// ...
}
```

**After:**
```go
switch {
case x > y:
	// ...
case x == 10:
	// ...
default: // <- last case (could also be the first one)
	// ...
}
```


## deferInLoop

[
  **diagnostic**
  **experimental** ]

Detects loops inside functions that use defer.





**Before:**
```go
for _, filename := range []string{"foo", "bar"} {
	 f, err := os.Open(filename)
	
	defer f.Close()
}
```

**After:**
```go
func process(filename string) {
	 f, err := os.Open(filename)
	
	defer f.Close()
}
/* ... */
for _, filename := range []string{"foo", "bar"} {
	process(filename)
}
```


## deferUnlambda

[
  **style**
  **experimental** ]

Detects deferred function literals that can be simplified.





**Before:**
```go
defer func() { f() }()
```

**After:**
```go
defer f()
```


## deprecatedComment

[
  **diagnostic** ]

Detects malformed 'deprecated' doc-comments.





**Before:**
```go
// deprecated, use FuncNew instead
func FuncOld() int
```

**After:**
```go
// Deprecated: use FuncNew instead
func FuncOld() int
```


## docStub

[
  **style**
  **experimental** ]

Detects comments that silence go lint complaints about doc-comment.





**Before:**
```go
// Foo ...
func Foo() {
}
```

**After:**
```go
// (A) - remove the doc-comment stub
func Foo() {}
// (B) - replace it with meaningful comment
// Foo is a demonstration-only function.
func Foo() {}
```


## dupArg

[
  **diagnostic** ]

Detects suspicious duplicated arguments.





**Before:**
```go
copy(dst, dst)
```

**After:**
```go
copy(dst, src)
```


## dupBranchBody

[
  **diagnostic** ]

Detects duplicated branch bodies inside conditional statements.





**Before:**
```go
if cond {
	println("cond=true")
} else {
	println("cond=true")
}
```

**After:**
```go
if cond {
	println("cond=true")
} else {
	println("cond=false")
}
```


## dupCase

[
  **diagnostic** ]

Detects duplicated case clauses inside switch or select statements.





**Before:**
```go
switch x {
case ys[0], ys[1], ys[2], ys[0], ys[4]:
}
```

**After:**
```go
switch x {
case ys[0], ys[1], ys[2], ys[3], ys[4]:
}
```


## dupImport

[
  **style**
  **experimental** ]

Detects multiple imports of the same package under different aliases.





**Before:**
```go
import (
	"fmt"
	printing "fmt" // Imported the second time
)
```

**After:**
```go
import(
	"fmt"
)
```


## dupSubExpr

[
  **diagnostic** ]

Detects suspicious duplicated sub-expressions.





**Before:**
```go
sort.Slice(xs, func(i, j int) bool {
	return xs[i].v < xs[i].v // Duplicated index
})
```

**After:**
```go
sort.Slice(xs, func(i, j int) bool {
	return xs[i].v < xs[j].v
})
```


## dynamicFmtString

[
  **diagnostic**
  **experimental** ]

Detects suspicious formatting strings usage.





**Before:**
```go
fmt.Errorf(msg)
```

**After:**
```go
fmt.Errorf("%s", msg)
```


## elseif

[
  **style** ]

Detects else with nested if statement that can be replaced with else-if.





**Before:**
```go
if cond1 {
} else {
	if x := cond2; x {
	}
}
```

**After:**
```go
if cond1 {
} else if x := cond2; x {
}
```


Checker parameters:
<ul>
<li>

  `@elseif.skipBalanced` whether to skip balanced if-else pairs (default true)

</li>

</ul>

## emptyDecl

[
  **diagnostic**
  **experimental** ]

Detects suspicious empty declarations blocks.





**Before:**
```go
var()
```

**After:**
```go
/* nothing */
```


## emptyFallthrough

[
  **style**
  **experimental** ]

Detects fallthrough that can be avoided by using multi case values.





**Before:**
```go
switch kind {
case reflect.Int:
	fallthrough
case reflect.Int32:
	return Int
}
```

**After:**
```go
switch kind {
case reflect.Int, reflect.Int32:
	return Int
}
```


## emptyStringTest

[
  **style**
  **experimental** ]

Detects empty string checks that can be written more idiomatically.





**Before:**
```go
len(s) == 0
```

**After:**
```go
s == ""
```


## equalFold

[
  **performance**
  **experimental** ]

Detects unoptimal strings/bytes case-insensitive comparison.





**Before:**
```go
strings.ToLower(x) == strings.ToLower(y)
```

**After:**
```go
strings.EqualFold(x, y)
```


## evalOrder

[
  **diagnostic**
  **experimental** ]

Detects unwanted dependencies on the evaluation order.





**Before:**
```go
return x, f(&x)
```

**After:**
```go
err := f(&x)
return x, err
```


## exitAfterDefer

[
  **diagnostic** ]

Detects calls to exit/fatal inside functions that use defer.





**Before:**
```go
defer os.Remove(filename)
if bad {
	log.Fatalf("something bad happened")
}
```

**After:**
```go
defer os.Remove(filename)
if bad {
	log.Printf("something bad happened")
	return
}
```


## exposedSyncMutex

[
  **style**
  **experimental** ]

Detects exposed methods from sync.Mutex and sync.RWMutex.





**Before:**
```go
type Foo struct{ ...; sync.Mutex; ... }
```

**After:**
```go
type Foo struct{ ...; mu sync.Mutex; ... }
```


## externalErrorReassign

[
  **diagnostic**
  **experimental** ]

Detects suspicious reassignment of error from another package.





**Before:**
```go
io.EOF = nil
```

**After:**
```go
/* don't do it */
```


## filepathJoin

[
  **diagnostic**
  **experimental** ]

Detects problems in filepath.Join() function calls.





**Before:**
```go
filepath.Join("dir/", filename)
```

**After:**
```go
filepath.Join("dir", filename)
```


## flagDeref

[
  **diagnostic** ]

Detects immediate dereferencing of `flag` package pointers.





**Before:**
```go
b := *flag.Bool("b", false, "b docs")
```

**After:**
```go
var b bool; flag.BoolVar(&b, "b", false, "b docs")
```


## flagName

[
  **diagnostic** ]

Detects suspicious flag names.




> https://github.com/golang/go/issues/41792

**Before:**
```go
b := flag.Bool(" foo ", false, "description")
```

**After:**
```go
b := flag.Bool("foo", false, "description")
```


## hexLiteral

[
  **style**
  **experimental** ]

Detects hex literals that have mixed case letter digits.





**Before:**
```go
x := 0X12
y := 0xfF
```

**After:**
```go
x := 0x12
// (A)
y := 0xff
// (B)
y := 0xFF
```


## httpNoBody

[
  **style**
  **experimental** ]

Detects nil usages in http.NewRequest calls, suggesting http.NoBody as an alternative.





**Before:**
```go
http.NewRequest("GET", url, nil)
```

**After:**
```go
http.NewRequest("GET", url, http.NoBody)
```


## hugeParam

[
  **performance** ]

Detects params that incur excessive amount of copying.





**Before:**
```go
func f(x [1024]int) {}
```

**After:**
```go
func f(x *[1024]int) {}
```


Checker parameters:
<ul>
<li>

  `@hugeParam.sizeThreshold` size in bytes that makes the warning trigger (default 80)

</li>

</ul>

## ifElseChain

[
  **style** ]

Detects repeated if-else statements and suggests to replace them with switch statement.




> Permits single else or else-if; repeated else-if or else + else-if
will trigger suggestion to use switch statement.
See [EffectiveGo#switch](https://golang.org/doc/effective_go.html#switch).

**Before:**
```go
if cond1 {
	// Code A.
} else if cond2 {
	// Code B.
} else {
	// Code C.
}
```

**After:**
```go
switch {
case cond1:
	// Code A.
case cond2:
	// Code B.
default:
	// Code C.
}
```


Checker parameters:
<ul>
<li>

  `@ifElseChain.minThreshold` min number of if-else blocks that makes the warning trigger (default 2)

</li>

</ul>

## importShadow

[
  **style**
  **opinionated** ]

Detects when imported package names shadowed in the assignments.





**Before:**
```go
// "path/filepath" is imported.
filepath := "foo.txt"
```

**After:**
```go
filename := "foo.txt"
```


## indexAlloc

[
  **performance** ]

Detects strings.Index calls that may cause unwanted allocs.




> See Go issue for details: https://github.com/golang/go/issues/25864

**Before:**
```go
strings.Index(string(x), y)
```

**After:**
```go
bytes.Index(x, []byte(y))
```


## initClause

[
  **style**
  **opinionated**
  **experimental** ]

Detects non-assignment statements inside if/switch init clause.





**Before:**
```go
if sideEffect(); cond {
}
```

**After:**
```go
sideEffect()
if cond {
}
```


## mapKey

[
  **diagnostic** ]

Detects suspicious map literal keys.





**Before:**
```go
_ = map[string]int{
	"foo": 1,
	"bar ": 2,
}
```

**After:**
```go
_ = map[string]int{
	"foo": 1,
	"bar": 2,
}
```


## methodExprCall

[
  **style**
  **experimental** ]

Detects method expression call that can be replaced with a method call.





**Before:**
```go
f := foo{}
foo.bar(f)
```

**After:**
```go
f := foo{}
f.bar()
```


## nestingReduce

[
  **style**
  **opinionated**
  **experimental** ]

Finds where nesting level could be reduced.





**Before:**
```go
for _, v := range a {
	if v.Bool {
		body()
	}
}
```

**After:**
```go
for _, v := range a {
	if !v.Bool {
		continue
	}
	body()
}
```


Checker parameters:
<ul>
<li>

  `@nestingReduce.bodyWidth` min number of statements inside a branch to trigger a warning (default 5)

</li>

</ul>

## newDeref

[
  **style** ]

Detects immediate dereferencing of `new` expressions.





**Before:**
```go
x := *new(bool)
```

**After:**
```go
x := false
```


## nilValReturn

[
  **diagnostic**
  **experimental** ]

Detects return statements those results evaluate to nil.





**Before:**
```go
if err == nil {
	return err
}
```

**After:**
```go
// (A) - return nil explicitly
if err == nil {
	return nil
}
// (B) - typo in "==", change to "!="
if err != nil {
	return err
}
```


## octalLiteral

[
  **style**
  **experimental**
  **opinionated** ]

Detects old-style octal literals.





**Before:**
```go
foo(02)
```

**After:**
```go
foo(0o2)
```


## offBy1

[
  **diagnostic** ]

Detects various off-by-one kind of errors.





**Before:**
```go
xs[len(xs)]
```

**After:**
```go
xs[len(xs)-1]
```


## paramTypeCombine

[
  **style**
  **opinionated** ]

Detects if function parameters could be combined by type and suggest the way to do it.





**Before:**
```go
func foo(a, b int, c, d int, e, f int, g int) {}
```

**After:**
```go
func foo(a, b, c, d, e, f, g int) {}
```


## preferDecodeRune

[
  **performance**
  **experimental** ]

Detects expressions like []rune(s)[0] that may cause unwanted rune slice allocation.




> See Go issue for details: https://github.com/golang/go/issues/45260

**Before:**
```go
r := []rune(s)[0]
```

**After:**
```go
r, _ := utf8.DecodeRuneInString(s)
```


## preferFilepathJoin

[
  **style**
  **experimental** ]

Detects concatenation with os.PathSeparator which can be replaced with filepath.Join.





**Before:**
```go
x + string(os.PathSeparator) + y
```

**After:**
```go
filepath.Join(x, y)
```


## preferFprint

[
  **performance**
  **experimental** ]

Detects fmt.Sprint(f/ln) calls which can be replaced with fmt.Fprint(f/ln).





**Before:**
```go
w.Write([]byte(fmt.Sprintf("%x", 10)))
```

**After:**
```go
fmt.Fprintf(w, "%x", 10)
```


## preferStringWriter

[
  **performance**
  **experimental** ]

Detects w.Write or io.WriteString calls which can be replaced with w.WriteString.





**Before:**
```go
w.Write([]byte("foo"))
```

**After:**
```go
w.WriteString("foo")
```


## preferWriteByte

[
  **performance**
  **experimental**
  **opinionated** ]

Detects WriteRune calls with rune literal argument that is single byte and reports to use WriteByte instead.





**Before:**
```go
w.WriteRune('\n')
```

**After:**
```go
w.WriteByte('\n')
```


## ptrToRefParam

[
  **style**
  **opinionated**
  **experimental** ]

Detects input and output parameters that have a type of pointer to referential type.





**Before:**
```go
func f(m *map[string]int) (*chan *int)
```

**After:**
```go
func f(m map[string]int) (chan *int)
```


## rangeExprCopy

[
  **performance** ]

Detects expensive copies of `for` loop range expressions.

Suggests to use pointer to array to avoid the copy using `&` on range expression.


> See Go issue for details: https://github.com/golang/go/issues/15812.

**Before:**
```go
var xs [2048]byte
for _, x := range xs { // Copies 2048 bytes
	// Loop body.
}
```

**After:**
```go
var xs [2048]byte
for _, x := range &xs { // No copy
	// Loop body.
}
```


Checker parameters:
<ul>
<li>

  `@rangeExprCopy.sizeThreshold` size in bytes that makes the warning trigger (default 512)

</li>
<li>

  `@rangeExprCopy.skipTestFuncs` whether to check test functions (default true)

</li>

</ul>

## rangeValCopy

[
  **performance** ]

Detects loops that copy big objects during each iteration.

Suggests to use index access or take address and make use pointer instead.



**Before:**
```go
xs := make([][1024]byte, length)
for _, x := range xs {
	// Loop body.
}
```

**After:**
```go
xs := make([][1024]byte, length)
for i := range xs {
	x := &xs[i]
	// Loop body.
}
```


Checker parameters:
<ul>
<li>

  `@rangeValCopy.sizeThreshold` size in bytes that makes the warning trigger (default 128)

</li>
<li>

  `@rangeValCopy.skipTestFuncs` whether to check test functions (default true)

</li>

</ul>

## redundantSprint

[
  **style**
  **experimental** ]

Detects redundant fmt.Sprint calls.





**Before:**
```go
fmt.Sprint(x)
```

**After:**
```go
x.String()
```


## regexpMust

[
  **style** ]

Detects `regexp.Compile*` that can be replaced with `regexp.MustCompile*`.





**Before:**
```go
re, _ := regexp.Compile("const pattern")
```

**After:**
```go
re := regexp.MustCompile("const pattern")
```


## regexpPattern

[
  **diagnostic**
  **experimental** ]

Detects suspicious regexp patterns.





**Before:**
```go
regexp.MustCompile(`google.com|yandex.ru`)
```

**After:**
```go
regexp.MustCompile(`google\.com|yandex\.ru`)
```


## regexpSimplify

[
  **style**
  **experimental**
  **opinionated** ]

Detects regexp patterns that can be simplified.





**Before:**
```go
regexp.MustCompile(`(?:a|b|c)   [a-z][a-z]*`)
```

**After:**
```go
regexp.MustCompile(`[abc] {3}[a-z]+`)
```


## returnAfterHttpError

[
  **diagnostic**
  **experimental** ]

Detects suspicious http.Error call without following return.





**Before:**
```go
if err != nil { http.Error(...); }
```

**After:**
```go
if err != nil { http.Error(...); return; }
```


## ruleguard

[
  **style**
  **experimental** ]

Runs user-defined rules using ruleguard linter.

Reads a rules file and turns them into go-critic checkers.


> See https://github.com/quasilyte/go-ruleguard.

**Before:**
```go
N/A
```

**After:**
```go
N/A
```


Checker parameters:
<ul>
<li>

  `@ruleguard.debug` enable debug for the specified named rules group (default )

</li>
<li>

  `@ruleguard.disable` comma-separated list of disabled groups or skip empty to enable everything (default )

</li>
<li>

  `@ruleguard.enable` comma-separated list of enabled groups or skip empty to enable everything (default &lt;all&gt;)

</li>
<li>

  `@ruleguard.failOn` Determines the behavior when an error occurs while parsing ruleguard files.
If flag is not set, log error and skip rule files that contain an error.
If flag is set, the value must be a comma-separated list of error conditions.
* 'import': rule refers to a package that cannot be loaded.
* 'dsl':    gorule file does not comply with the ruleguard DSL. (default )

</li>
<li>

  `@ruleguard.failOnError` deprecated, use failOn param; if set to true, identical to failOn='all', otherwise failOn='' (default false)

</li>
<li>

  `@ruleguard.rules` comma-separated list of gorule file paths. Glob patterns such as 'rules-*.go' may be specified (default )

</li>

</ul>

## singleCaseSwitch

[
  **style** ]

Detects switch statements that could be better written as if statement.





**Before:**
```go
switch x := x.(type) {
case int:
	body()
}
```

**After:**
```go
if x, ok := x.(int); ok {
	body()
}
```


## sliceClear

[
  **performance**
  **experimental** ]

Detects slice clear loops, suggests an idiom that is recognized by the Go compiler.





**Before:**
```go
for i := 0; i < len(buf); i++ { buf[i] = 0 }
```

**After:**
```go
for i := range buf { buf[i] = 0 }
```


## sloppyLen

[
  **diagnostic** ]

Detects usage of `len` when result is obvious or doesn't make sense.





**Before:**
```go
len(arr) <= 0
```

**After:**
```go
len(arr) == 0
```


## sloppyReassign

[
  **diagnostic**
  **experimental** ]

Detects suspicious/confusing re-assignments.





**Before:**
```go
if err = f(); err != nil { return err }
```

**After:**
```go
if err := f(); err != nil { return err }
```


## sloppyTypeAssert

[
  **diagnostic** ]

Detects redundant type assertions.





**Before:**
```go
func f(r io.Reader) interface{} {
	return r.(interface{})
}
```

**After:**
```go
func f(r io.Reader) interface{} {
	return r
}
```


## sortSlice

[
  **diagnostic**
  **experimental** ]

Detects suspicious sort.Slice calls.





**Before:**
```go
sort.Slice(xs, func(i, j) bool { return keys[i] < keys[j] })
```

**After:**
```go
sort.Slice(kv, func(i, j) bool { return kv[i].key < kv[j].key })
```


## sprintfQuotedString

[
  **diagnostic**
  **experimental** ]

Detects "%s" formatting directives that can be replaced with %q.





**Before:**
```go
fmt.Sprintf(`"%s"`, s)
```

**After:**
```go
fmt.Sprintf(`%q`, s)
```


## sqlQuery

[
  **diagnostic**
  **experimental** ]

Detects issue in Query() and Exec() calls.





**Before:**
```go
_, err := db.Query("UPDATE ...")
```

**After:**
```go
_, err := db.Exec("UPDATE ...")
```


## stringConcatSimplify

[
  **style**
  **experimental** ]

Detects string concat operations that can be simplified.





**Before:**
```go
strings.Join([]string{x, y}, "_")
```

**After:**
```go
x + "_" + y
```


## stringXbytes

[
  **performance** ]

Detects redundant conversions between string and []byte.





**Before:**
```go
copy(b, []byte(s))
```

**After:**
```go
copy(b, s)
```


## stringsCompare

[
  **style**
  **experimental** ]

Detects strings.Compare usage.





**Before:**
```go
strings.Compare(x, y)
```

**After:**
```go
x < y
```


## switchTrue

[
  **style** ]

Detects switch-over-bool statements that use explicit `true` tag value.





**Before:**
```go
switch true {...}
```

**After:**
```go
switch {...}
```


## syncMapLoadAndDelete

[
  **diagnostic**
  **experimental** ]

Detects sync.Map load+delete operations that can be replaced with LoadAndDelete.





**Before:**
```go
v, ok := m.Load(k); if ok { m.Delete($k); f(v); }
```

**After:**
```go
v, deleted := m.LoadAndDelete(k); if deleted { f(v) }
```


## timeExprSimplify

[
  **style**
  **experimental** ]

Detects manual conversion to milli- or microseconds.





**Before:**
```go
t.Unix() / 1000
```

**After:**
```go
t.UnixMilli()
```


## todoCommentWithoutDetail

[
  **style**
  **opinionated**
  **experimental** ]

Detects TODO comments without detail/assignee.





**Before:**
```go
// TODO
fiiWithCtx(nil, a, b)
```

**After:**
```go
// TODO(admin): pass context.TODO() instead of nil
fiiWithCtx(nil, a, b)
```


## tooManyResultsChecker

[
  **style**
  **opinionated**
  **experimental** ]

Detects function with too many results.





**Before:**
```go
func fn() (a, b, c, d float32, _ int, _ bool)
```

**After:**
```go
func fn() (resultStruct, bool)
```


Checker parameters:
<ul>
<li>

  `@tooManyResultsChecker.maxResults` maximum number of results (default 5)

</li>

</ul>

## truncateCmp

[
  **diagnostic**
  **experimental** ]

Detects potential truncation issues when comparing ints of different sizes.





**Before:**
```go
func f(x int32, y int16) bool {
  return int16(x) < y
}
```

**After:**
```go
func f(x int32, int16) bool {
  return x < int32(y)
}
```


Checker parameters:
<ul>
<li>

  `@truncateCmp.skipArchDependent` whether to skip int/uint/uintptr types (default true)

</li>

</ul>

## typeAssertChain

[
  **style**
  **experimental** ]

Detects repeated type assertions and suggests to replace them with type switch statement.





**Before:**
```go
if x, ok := v.(T1); ok {
	// Code A, uses x.
} else if x, ok := v.(T2); ok {
	// Code B, uses x.
} else if x, ok := v.(T3); ok {
	// Code C, uses x.
}
```

**After:**
```go
switch x := v.(T1) {
case cond1:
	// Code A, uses x.
case cond2:
	// Code B, uses x.
default:
	// Code C, uses x.
}
```


## typeDefFirst

[
  **style**
  **experimental** ]

Detects method declarations preceding the type definition itself.





**Before:**
```go
func (r rec) Method() {}
type rec struct{}
```

**After:**
```go
type rec struct{}
func (r rec) Method() {}
```


## typeSwitchVar

[
  **style** ]

Detects type switches that can benefit from type guard clause with variable.





**Before:**
```go
switch v.(type) {
case int:
	return v.(int)
case point:
	return v.(point).x + v.(point).y
default:
	return 0
}
```

**After:**
```go
switch v := v.(type) {
case int:
	return v
case point:
	return v.x + v.y
default:
	return 0
}
```


## typeUnparen

[
  **style**
  **opinionated** ]

Detects unneeded parenthesis inside type expressions and suggests to remove them.





**Before:**
```go
type foo [](func([](func())))
```

**After:**
```go
type foo []func([]func())
```


## uncheckedInlineErr

[
  **diagnostic**
  **experimental** ]

Detects unchecked errors in if statements.





**Before:**
```go
if err := expr(); err2 != nil { /*...*/ }
```

**After:**
```go
if err := expr(); err != nil { /*...*/ }
```


## underef

[
  **style** ]

Detects dereference expressions that can be omitted.





**Before:**
```go
(*k).field = 5
v := (*a)[5] // only if a is array
```

**After:**
```go
k.field = 5
v := a[5]
```


Checker parameters:
<ul>
<li>

  `@underef.skipRecvDeref` whether to skip (*x).method() calls where x is a pointer receiver (default true)

</li>

</ul>

## unlabelStmt

[
  **style**
  **experimental** ]

Detects redundant statement labels.





**Before:**
```go
derp:
for x := range xs {
	if x == 0 {
		break derp
	}
}
```

**After:**
```go
for x := range xs {
	if x == 0 {
		break
	}
}
```


## unlambda

[
  **style** ]

Detects function literals that can be simplified.





**Before:**
```go
func(x int) int { return fn(x) }
```

**After:**
```go
fn
```


## unnamedResult

[
  **style**
  **opinionated**
  **experimental** ]

Detects unnamed results that may benefit from names.





**Before:**
```go
func f() (float64, float64)
```

**After:**
```go
func f() (x, y float64)
```


Checker parameters:
<ul>
<li>

  `@unnamedResult.checkExported` whether to check exported functions (default false)

</li>

</ul>

## unnecessaryBlock

[
  **style**
  **opinionated**
  **experimental** ]

Detects unnecessary braced statement blocks.





**Before:**
```go
x := 1
{
	print(x)
}
```

**After:**
```go
x := 1
print(x)
```


## unnecessaryDefer

[
  **diagnostic**
  **experimental** ]

Detects redundantly deferred calls.





**Before:**
```go
func() {
	defer os.Remove(filename)
}
```

**After:**
```go
func() {
	os.Remove(filename)
}
```


## unslice

[
  **style** ]

Detects slice expressions that can be simplified to sliced expression itself.





**Before:**
```go
copy(b[:], values...)
```

**After:**
```go
copy(b, values...)
```


## valSwap

[
  **style** ]

Detects value swapping code that are not using parallel assignment.





**Before:**
```go
*tmp = *x; *x = *y; *y = *tmp
```

**After:**
```go
*x, *y = *y, *x
```


## weakCond

[
  **diagnostic**
  **experimental** ]

Detects conditions that are unsafe due to not being exhaustive.





**Before:**
```go
xs != nil && xs[0] != nil
```

**After:**
```go
len(xs) != 0 && xs[0] != nil
```


## whyNoLint

[
  **style**
  **experimental** ]

Ensures that `//nolint` comments include an explanation.





**Before:**
```go
//nolint
```

**After:**
```go
//nolint // reason
```


## wrapperFunc

[
  **style** ]

Detects function calls that can be replaced with convenience wrappers.





**Before:**
```go
wg.Add(-1)
```

**After:**
```go
wg.Done()
```


## yodaStyleExpr

[
  **style**
  **experimental** ]

Detects Yoda style expressions and suggests to replace them.





**Before:**
```go
return nil != ptr
```

**After:**
```go
return ptr != nil
```


