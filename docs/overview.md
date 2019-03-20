## Checks overview

This page describes checks supported by [go-critic](https://github.com/go-critic/go-critic) linter.

[//]: # (This is generated file, please don't edit it yourself.)

## Checkers

* :heavy_check_mark: checker is enabled by default.
* :white_check_mark: checker is disabled by default.

### Checkers from the "diagnostic" group

Diagnostics try to find programming errors in the code.
They also detect code that may be correct, but looks suspicious.

> All diagnostics are enabled by default (unless it has "experimental" tag).

<table>
  <tr>
    <th>Name</th>
    <th>Short description</th>
  </tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#appendAssign-ref">appendAssign</a>
  </td>
  <td>Detects suspicious append result assignments</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#argOrder-ref">argOrder</a>
  </td>
  <td>Detects suspicious arguments order</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#badCall-ref">badCall</a>
  </td>
  <td>Detects suspicious function calls</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#badCond-ref">badCond</a>
  </td>
  <td>Detects suspicious condition expressions</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#caseOrder-ref">caseOrder</a>
  </td>
  <td>Detects erroneous case order inside switch statements</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#codegenComment-ref">codegenComment</a>
  </td>
  <td>Detects malformed 'code generated' file comments</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#commentedOutCode-ref">commentedOutCode</a>
  </td>
  <td>Detects commented-out code inside function bodies</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#deprecatedComment-ref">deprecatedComment</a>
  </td>
  <td>Detects malformed 'deprecated' doc-comments</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#dupArg-ref">dupArg</a>
  </td>
  <td>Detects suspicious duplicated arguments</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#dupBranchBody-ref">dupBranchBody</a>
  </td>
  <td>Detects duplicated branch bodies inside conditional statements</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#dupCase-ref">dupCase</a>
  </td>
  <td>Detects duplicated case clauses inside switch statements</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#dupSubExpr-ref">dupSubExpr</a>
  </td>
  <td>Detects suspicious duplicated sub-expressions</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#exitAfterDefer-ref">exitAfterDefer</a>
  </td>
  <td>Detects calls to exit/fatal inside functions that use defer</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#flagDeref-ref">flagDeref</a>
  </td>
  <td>Detects immediate dereferencing of `flag` package pointers</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#flagName-ref">flagName</a>
  </td>
  <td>Detects flag names with whitespace</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#nilValReturn-ref">nilValReturn</a>
  </td>
  <td>Detects return statements those results evaluate to nil</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#octalLiteral-ref">octalLiteral</a>
  </td>
  <td>Detects octal literals passed to functions</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#offBy1-ref">offBy1</a>
  </td>
  <td>Detects various off-by-one kind of errors</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#sloppyReassign-ref">sloppyReassign</a>
  </td>
  <td>Detects suspicious/confusing re-assignments</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#weakCond-ref">weakCond</a>
  </td>
  <td>Detects conditions that are unsafe due to not being exhaustive</td>
</tr>
</table>

### Checkers from the "style" group

Style checks suggest replacing some form of expression/statement
with another one that is considered more idiomatic or simple.

> Only non-opinionated style checks are enabled by default.

<table>
  <tr>
    <th>Name</th>
    <th>Short description</th>
  </tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#assignOp-ref">assignOp</a>
  </td>
  <td>Detects assignments that can be simplified by using assignment operators</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#boolExprSimplify-ref">boolExprSimplify</a>
  </td>
  <td>Detects bool expressions that can be simplified</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#builtinShadow-ref">builtinShadow</a>
  </td>
  <td>Detects when predeclared identifiers shadowed in assignments</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#captLocal-ref">captLocal</a>
  </td>
  <td>Detects capitalized names for local variables</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#commentFormatting-ref">commentFormatting</a>
  </td>
  <td>Detects comments with non-idiomatic formatting</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#commentedOutImport-ref">commentedOutImport</a>
  </td>
  <td>Detects commented-out imports</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#defaultCaseOrder-ref">defaultCaseOrder</a>
  </td>
  <td>Detects when default case in switch isn't on 1st or last position</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#docStub-ref">docStub</a>
  </td>
  <td>Detects comments that silence go lint complaints about doc-comment</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#dupImport-ref">dupImport</a>
  </td>
  <td>Detects multiple imports of the same package under different aliases</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#elseif-ref">elseif</a>
  </td>
  <td>Detects else with nested if statement that can be replaced with else-if</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#emptyFallthrough-ref">emptyFallthrough</a>
  </td>
  <td>Detects fallthrough that can be avoided by using multi case values</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#emptyStringTest-ref">emptyStringTest</a>
  </td>
  <td>Detects empty string checks that can be written more idiomatically</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#hexLiteral-ref">hexLiteral</a>
  </td>
  <td>Detects hex literals that have mixed case letter digits</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#ifElseChain-ref">ifElseChain</a>
  </td>
  <td>Detects repeated if-else statements and suggests to replace them with switch statement</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#importShadow-ref">importShadow</a>
  </td>
  <td>Detects when imported package names shadowed in the assignments</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#initClause-ref">initClause</a>
  </td>
  <td>Detects non-assignment statements inside if/switch init clause</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#methodExprCall-ref">methodExprCall</a>
  </td>
  <td>Detects method expression call that can be replaced with a method call</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#nestingReduce-ref">nestingReduce</a>
  </td>
  <td>Finds where nesting level could be reduced</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#newDeref-ref">newDeref</a>
  </td>
  <td>Detects immediate dereferencing of `new` expressions</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#paramTypeCombine-ref">paramTypeCombine</a>
  </td>
  <td>Detects if function parameters could be combined by type and suggest the way to do it</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#ptrToRefParam-ref">ptrToRefParam</a>
  </td>
  <td>Detects input and output parameters that have a type of pointer to referential type</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#regexpMust-ref">regexpMust</a>
  </td>
  <td>Detects `regexp.Compile*` that can be replaced with `regexp.MustCompile*`</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#singleCaseSwitch-ref">singleCaseSwitch</a>
  </td>
  <td>Detects switch statements that could be better written as if statement</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#sloppyLen-ref">sloppyLen</a>
  </td>
  <td>Detects usage of `len` when result is obvious or doesn't make sense</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#stringXbytes-ref">stringXbytes</a>
  </td>
  <td>Detects redundant conversions between string and []byte</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#switchTrue-ref">switchTrue</a>
  </td>
  <td>Detects switch-over-bool statements that use explicit `true` tag value</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#typeAssertChain-ref">typeAssertChain</a>
  </td>
  <td>Detects repeated type assertions and suggests to replace them with type switch statement</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#typeSwitchVar-ref">typeSwitchVar</a>
  </td>
  <td>Detects type switches that can benefit from type guard clause with variable</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#typeUnparen-ref">typeUnparen</a>
  </td>
  <td>Detects unneded parenthesis inside type expressions and suggests to remove them</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#underef-ref">underef</a>
  </td>
  <td>Detects dereference expressions that can be omitted</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#unlabelStmt-ref">unlabelStmt</a>
  </td>
  <td>Detects redundant statement labels</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#unlambda-ref">unlambda</a>
  </td>
  <td>Detects function literals that can be simplified</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#unnamedResult-ref">unnamedResult</a>
  </td>
  <td>Detects unnamed results that may benefit from names</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#unnecessaryBlock-ref">unnecessaryBlock</a>
  </td>
  <td>Detects unnecessary braced statement blocks</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#unslice-ref">unslice</a>
  </td>
  <td>Detects slice expressions that can be simplified to sliced expression itself</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#valSwap-ref">valSwap</a>
  </td>
  <td>Detects value swapping code that are not using parallel assignment</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#wrapperFunc-ref">wrapperFunc</a>
  </td>
  <td>Detects function calls that can be replaced with convenience wrappers</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#yodaStyleExpr-ref">yodaStyleExpr</a>
  </td>
  <td>Detects Yoda style expressions and suggests to replace them</td>
</tr>
</table>

### Checkers from the "performance" group

Performance checks tell you about potential issues that
can make your code run slower than it could be.

> All performance checks are disabled by default.

<table>
  <tr>
    <th>Name</th>
    <th>Short description</th>
  </tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#appendCombine-ref">appendCombine</a>
  </td>
  <td>Detects `append` chains to the same slice that can be done in a single `append` call</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#equalFold-ref">equalFold</a>
  </td>
  <td>Detects unoptimal strings/bytes case-insensitive comparison</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#hugeParam-ref">hugeParam</a>
  </td>
  <td>Detects params that incur excessive amount of copying</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#indexAlloc-ref">indexAlloc</a>
  </td>
  <td>Detects strings.Index calls that may cause unwanted allocs</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#rangeExprCopy-ref">rangeExprCopy</a>
  </td>
  <td>Detects expensive copies of `for` loop range expressions</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#rangeValCopy-ref">rangeValCopy</a>
  </td>
  <td>Detects loops that copy big objects during each iteration</td>
</tr>
</table>


  <a name="appendAssign-ref"></a>
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



  <a name="appendCombine-ref"></a>
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



  <a name="argOrder-ref"></a>
## argOrder

[
  **diagnostic**
  **experimental** ]

Detects suspicious arguments order.





**Before:**
```go
strings.HasPrefix("#", userpass)
```

**After:**
```go
strings.HasPrefix(userpass, "#")
```



  <a name="assignOp-ref"></a>
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



  <a name="badCall-ref"></a>
## badCall

[
  **diagnostic**
  **experimental** ]

Detects suspicious function calls.





**Before:**
```go
strings.Replace(s, from, to, 0)
```

**After:**
```go
strings.Replace(s, from, to, -1)
```



  <a name="badCond-ref"></a>
## badCond

[
  **diagnostic**
  **experimental** ]

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



  <a name="boolExprSimplify-ref"></a>
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



  <a name="builtinShadow-ref"></a>
## builtinShadow

[
  **style**
  **opinionated** ]

Detects when predeclared identifiers shadowed in assignments.





**Before:**
```go
len := 10
```

**After:**
```go
length := 10
```



  <a name="captLocal-ref"></a>
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


  <a name="caseOrder-ref"></a>
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



  <a name="codegenComment-ref"></a>
## codegenComment

[
  **diagnostic**
  **experimental** ]

Detects malformed 'code generated' file comments.





**Before:**
```go
// This file was automatically generated by foogen
```

**After:**
```go
// Code generated by foogen. DO NOT EDIT.
```



  <a name="commentFormatting-ref"></a>
## commentFormatting

[
  **style**
  **experimental** ]

Detects comments with non-idiomatic formatting.





**Before:**
```go
//This is a comment
```

**After:**
```go
// This is a comment
```



  <a name="commentedOutCode-ref"></a>
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



  <a name="commentedOutImport-ref"></a>
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



  <a name="defaultCaseOrder-ref"></a>
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



  <a name="deprecatedComment-ref"></a>
## deprecatedComment

[
  **diagnostic**
  **experimental** ]

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



  <a name="docStub-ref"></a>
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



  <a name="dupArg-ref"></a>
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



  <a name="dupBranchBody-ref"></a>
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



  <a name="dupCase-ref"></a>
## dupCase

[
  **diagnostic** ]

Detects duplicated case clauses inside switch statements.





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



  <a name="dupImport-ref"></a>
## dupImport

[
  **style**
  **experimental** ]

Detects multiple imports of the same package under different aliases.





**Before:**
```go
import (
	"fmt"
	priting "fmt" // Imported the second time
)
```

**After:**
```go
import(
	"fmt"
)
```



  <a name="dupSubExpr-ref"></a>
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



  <a name="elseif-ref"></a>
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


  <a name="emptyFallthrough-ref"></a>
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



  <a name="emptyStringTest-ref"></a>
## emptyStringTest

[
  **style**
  **experimental** ]

Detects empty string checks that can be written more idiomatically.




> See https://dmitri.shuralyov.com/idiomatic-go#empty-string-check.

**Before:**
```go
len(s) == 0
```

**After:**
```go
s == ""
```



  <a name="equalFold-ref"></a>
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



  <a name="exitAfterDefer-ref"></a>
## exitAfterDefer

[
  **diagnostic**
  **experimental** ]

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



  <a name="flagDeref-ref"></a>
## flagDeref

[
  **diagnostic** ]

Detects immediate dereferencing of `flag` package pointers.

Suggests to use pointer to array to avoid the copy using `&` on range expression.


> Dereferencing returned pointers will lead to hard to find errors
where flag values are not updated after flag.Parse().

**Before:**
```go
b := *flag.Bool("b", false, "b docs")
```

**After:**
```go
var b bool
flag.BoolVar(&b, "b", false, "b docs")
```



  <a name="flagName-ref"></a>
## flagName

[
  **diagnostic**
  **experimental** ]

Detects flag names with whitespace.





**Before:**
```go
b := flag.Bool(" foo ", false, "description")
```

**After:**
```go
b := flag.Bool("foo", false, "description")
```



  <a name="hexLiteral-ref"></a>
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



  <a name="hugeParam-ref"></a>
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


  <a name="ifElseChain-ref"></a>
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



  <a name="importShadow-ref"></a>
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



  <a name="indexAlloc-ref"></a>
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



  <a name="initClause-ref"></a>
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



  <a name="methodExprCall-ref"></a>
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



  <a name="nestingReduce-ref"></a>
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


  <a name="newDeref-ref"></a>
## newDeref

[
  **style**
  **experimental** ]

Detects immediate dereferencing of `new` expressions.





**Before:**
```go
x := *new(bool)
```

**After:**
```go
x := false
```



  <a name="nilValReturn-ref"></a>
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



  <a name="octalLiteral-ref"></a>
## octalLiteral

[
  **diagnostic**
  **experimental** ]

Detects octal literals passed to functions.





**Before:**
```go
foo(02)
```

**After:**
```go
foo(2)
```



  <a name="offBy1-ref"></a>
## offBy1

[
  **diagnostic**
  **experimental** ]

Detects various off-by-one kind of errors.





**Before:**
```go
xs[len(xs)]
```

**After:**
```go
xs[len(xs)-1]
```



  <a name="paramTypeCombine-ref"></a>
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



  <a name="ptrToRefParam-ref"></a>
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



  <a name="rangeExprCopy-ref"></a>
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


  <a name="rangeValCopy-ref"></a>
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


  <a name="regexpMust-ref"></a>
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



  <a name="singleCaseSwitch-ref"></a>
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



  <a name="sloppyLen-ref"></a>
## sloppyLen

[
  **style** ]

Detects usage of `len` when result is obvious or doesn't make sense.





**Before:**
```go
len(arr) >= 0 // Sloppy
len(arr) <= 0 // Sloppy
len(arr) < 0  // Doesn't make sense at all
```

**After:**
```go
len(arr) > 0
len(arr) == 0
```



  <a name="sloppyReassign-ref"></a>
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



  <a name="stringXbytes-ref"></a>
## stringXbytes

[
  **style**
  **experimental** ]

Detects redundant conversions between string and []byte.





**Before:**
```go
copy(b, []byte(s))
```

**After:**
```go
copy(b, s)
```



  <a name="switchTrue-ref"></a>
## switchTrue

[
  **style** ]

Detects switch-over-bool statements that use explicit `true` tag value.





**Before:**
```go
switch true {
case x > y:
}
```

**After:**
```go
switch {
case x > y:
}
```



  <a name="typeAssertChain-ref"></a>
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



  <a name="typeSwitchVar-ref"></a>
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



  <a name="typeUnparen-ref"></a>
## typeUnparen

[
  **style**
  **opinionated** ]

Detects unneded parenthesis inside type expressions and suggests to remove them.





**Before:**
```go
type foo [](func([](func())))
```

**After:**
```go
type foo []func([]func())
```



  <a name="underef-ref"></a>
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


  <a name="unlabelStmt-ref"></a>
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



  <a name="unlambda-ref"></a>
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



  <a name="unnamedResult-ref"></a>
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


  <a name="unnecessaryBlock-ref"></a>
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



  <a name="unslice-ref"></a>
## unslice

[
  **style** ]

Detects slice expressions that can be simplified to sliced expression itself.





**Before:**
```go
f(s[:])               // s is string
copy(b[:], values...) // b is []byte
```

**After:**
```go
f(s)
copy(b, values...)
```



  <a name="valSwap-ref"></a>
## valSwap

[
  **style**
  **experimental** ]

Detects value swapping code that are not using parallel assignment.





**Before:**
```go
tmp := *x
*x = *y
*y = tmp
```

**After:**
```go
*x, *y = *y, *x
```



  <a name="weakCond-ref"></a>
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



  <a name="wrapperFunc-ref"></a>
## wrapperFunc

[
  **style**
  **experimental** ]

Detects function calls that can be replaced with convenience wrappers.





**Before:**
```go
wg.Add(-1)
```

**After:**
```go
wg.Done()
```



  <a name="yodaStyleExpr-ref"></a>
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


