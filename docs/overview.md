## Checks overview

This page describes checks supported by [go-critic](https://github.com/go-critic/go-critic) linter.

[//]: # (This is generated file, please don't edit it yourself.)

## Checkers

Total number of checks is 98 :rocket:

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
  <td nowrap>:heavy_check_mark:
    <a href="#argOrder-ref">argOrder</a>
  </td>
  <td>Detects suspicious arguments order</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#badCall-ref">badCall</a>
  </td>
  <td>Detects suspicious function calls</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#badCond-ref">badCond</a>
  </td>
  <td>Detects suspicious condition expressions</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#badLock-ref">badLock</a>
  </td>
  <td>Detects suspicious mutex lock/unlock operations</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#badRegexp-ref">badRegexp</a>
  </td>
  <td>Detects suspicious regexp patterns</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#builtinShadowDecl-ref">builtinShadowDecl</a>
  </td>
  <td>Detects top-level declarations that shadow the predeclared identifiers</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#caseOrder-ref">caseOrder</a>
  </td>
  <td>Detects erroneous case order inside switch statements</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#codegenComment-ref">codegenComment</a>
  </td>
  <td>Detects malformed 'code generated' file comments</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#commentedOutCode-ref">commentedOutCode</a>
  </td>
  <td>Detects commented-out code inside function bodies</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
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
    <a href="#evalOrder-ref">evalOrder</a>
  </td>
  <td>Detects unwanted dependencies on the evaluation order</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#exitAfterDefer-ref">exitAfterDefer</a>
  </td>
  <td>Detects calls to exit/fatal inside functions that use defer</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#filepathJoin-ref">filepathJoin</a>
  </td>
  <td>Detects problems in filepath.Join() function calls</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#flagDeref-ref">flagDeref</a>
  </td>
  <td>Detects immediate dereferencing of `flag` package pointers</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#flagName-ref">flagName</a>
  </td>
  <td>Detects suspicious flag names</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#mapKey-ref">mapKey</a>
  </td>
  <td>Detects suspicious map literal keys</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#nilValReturn-ref">nilValReturn</a>
  </td>
  <td>Detects return statements those results evaluate to nil</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#offBy1-ref">offBy1</a>
  </td>
  <td>Detects various off-by-one kind of errors</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#regexpPattern-ref">regexpPattern</a>
  </td>
  <td>Detects suspicious regexp patterns</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#returnAfterHttpError-ref">returnAfterHttpError</a>
  </td>
  <td>Detects suspicious http.Error call without following return</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#sloppyReassign-ref">sloppyReassign</a>
  </td>
  <td>Detects suspicious/confusing re-assignments</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#sloppyTypeAssert-ref">sloppyTypeAssert</a>
  </td>
  <td>Detects redundant type assertions</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#sortSlice-ref">sortSlice</a>
  </td>
  <td>Detects suspicious sort.Slice calls</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#sprintfQuotedString-ref">sprintfQuotedString</a>
  </td>
  <td>Detects "%s" formatting directives that can be replaced with %q</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#sqlQuery-ref">sqlQuery</a>
  </td>
  <td>Detects issue in Query() and Exec() calls</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#suspiciousSorting-ref">suspiciousSorting</a>
  </td>
  <td>Detects bad usage of sort package</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#syncMapLoadAndDelete-ref">syncMapLoadAndDelete</a>
  </td>
  <td>Detects sync.Map load+delete operations that can be replaced with LoadAndDelete</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#truncateCmp-ref">truncateCmp</a>
  </td>
  <td>Detects potential truncation issues when comparing ints of different sizes</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#unnecessaryDefer-ref">unnecessaryDefer</a>
  </td>
  <td>Detects redundantly deferred calls</td>
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
  <td>Detects when predeclared identifiers are shadowed in assignments</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#captLocal-ref">captLocal</a>
  </td>
  <td>Detects capitalized names for local variables</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
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
    <a href="#deferUnlambda-ref">deferUnlambda</a>
  </td>
  <td>Detects deferred function literals that can be simplified</td>
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
    <a href="#exposedSyncMutex-ref">exposedSyncMutex</a>
  </td>
  <td>Detects exposed methods from sync.Mutex and sync.RWMutex</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#hexLiteral-ref">hexLiteral</a>
  </td>
  <td>Detects hex literals that have mixed case letter digits</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#httpNoBody-ref">httpNoBody</a>
  </td>
  <td>Detects nil usages in http.NewRequest calls, suggesting http.NoBody as an alternative</td>
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
    <a href="#ioutilDeprecated-ref">ioutilDeprecated</a>
  </td>
  <td>Detects deprecated io/ioutil package usages</td>
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
  <td nowrap>:heavy_check_mark:
    <a href="#newDeref-ref">newDeref</a>
  </td>
  <td>Detects immediate dereferencing of `new` expressions</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#octalLiteral-ref">octalLiteral</a>
  </td>
  <td>Detects old-style octal literals</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#paramTypeCombine-ref">paramTypeCombine</a>
  </td>
  <td>Detects if function parameters could be combined by type and suggest the way to do it</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#preferFilepathJoin-ref">preferFilepathJoin</a>
  </td>
  <td>Detects concatenation with os.PathSeparator which can be replaced with filepath.Join</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#ptrToRefParam-ref">ptrToRefParam</a>
  </td>
  <td>Detects input and output parameters that have a type of pointer to referential type</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#redundantSprint-ref">redundantSprint</a>
  </td>
  <td>Detects redundant fmt.Sprint calls</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#regexpMust-ref">regexpMust</a>
  </td>
  <td>Detects `regexp.Compile*` that can be replaced with `regexp.MustCompile*`</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#regexpSimplify-ref">regexpSimplify</a>
  </td>
  <td>Detects regexp patterns that can be simplified</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#ruleguard-ref">ruleguard</a>
  </td>
  <td>Runs user-defined rules using ruleguard linter</td>
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
    <a href="#stringConcatSimplify-ref">stringConcatSimplify</a>
  </td>
  <td>Detects string concat operations that can be simplified</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
    <a href="#switchTrue-ref">switchTrue</a>
  </td>
  <td>Detects switch-over-bool statements that use explicit `true` tag value</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#timeExprSimplify-ref">timeExprSimplify</a>
  </td>
  <td>Detects manual conversion to milli- or microseconds</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#tooManyResultsChecker-ref">tooManyResultsChecker</a>
  </td>
  <td>Detects function with too many results</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#typeAssertChain-ref">typeAssertChain</a>
  </td>
  <td>Detects repeated type assertions and suggests to replace them with type switch statement</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#typeDefFirst-ref">typeDefFirst</a>
  </td>
  <td>Detects method declarations preceding the type definition itself</td>
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
  <td nowrap>:heavy_check_mark:
    <a href="#valSwap-ref">valSwap</a>
  </td>
  <td>Detects value swapping code that are not using parallel assignment</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#whyNoLint-ref">whyNoLint</a>
  </td>
  <td>Ensures that `//nolint` comments include an explanation</td>
</tr><tr>
  <td nowrap>:heavy_check_mark:
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
    <a href="#preferDecodeRune-ref">preferDecodeRune</a>
  </td>
  <td>Detects expressions like []rune(s)[0] that may cause unwanted rune slice allocation</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#preferFprint-ref">preferFprint</a>
  </td>
  <td>Detects fmt.Sprint(f|ln) calls which can be replaced with fmt.Fprint(f|ln)</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#preferStringWriter-ref">preferStringWriter</a>
  </td>
  <td>Detects w.Write or io.WriteString calls which can be replaced with w.WriteString</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#preferWriteByte-ref">preferWriteByte</a>
  </td>
  <td>Detects WriteRune calls with byte literal argument and reports to use WriteByte instead</td>
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
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#sliceClear-ref">sliceClear</a>
  </td>
  <td>Detects slice clear loops, suggests an idiom that is recognized by the Go compiler</td>
</tr><tr>
  <td nowrap>:white_check_mark:
    <a href="#stringXbytes-ref">stringXbytes</a>
  </td>
  <td>Detects redundant conversions between string and []byte</td>
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



  <a name="badCond-ref"></a>
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



  <a name="badLock-ref"></a>
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



  <a name="badRegexp-ref"></a>
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

Detects when predeclared identifiers are shadowed in assignments.





**Before:**
```go
len := 10
```

**After:**
```go
length := 10
```



  <a name="builtinShadowDecl-ref"></a>
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



  <a name="commentFormatting-ref"></a>
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



  <a name="deferUnlambda-ref"></a>
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



  <a name="deprecatedComment-ref"></a>
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



  <a name="evalOrder-ref"></a>
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



  <a name="exitAfterDefer-ref"></a>
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



  <a name="exposedSyncMutex-ref"></a>
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



  <a name="filepathJoin-ref"></a>
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



  <a name="flagDeref-ref"></a>
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



  <a name="flagName-ref"></a>
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



  <a name="httpNoBody-ref"></a>
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



  <a name="ioutilDeprecated-ref"></a>
## ioutilDeprecated

[
  **style**
  **experimental** ]

Detects deprecated io/ioutil package usages.





**Before:**
```go
ioutil.ReadAll(r)
```

**After:**
```go
io.ReadAll(r)
```



  <a name="mapKey-ref"></a>
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



  <a name="offBy1-ref"></a>
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



  <a name="preferDecodeRune-ref"></a>
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



  <a name="preferFilepathJoin-ref"></a>
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



  <a name="preferFprint-ref"></a>
## preferFprint

[
  **performance**
  **experimental** ]

Detects fmt.Sprint(f|ln) calls which can be replaced with fmt.Fprint(f|ln).





**Before:**
```go
w.Write([]byte(fmt.Sprintf("%x", 10)))
```

**After:**
```go
fmt.Fprintf(w, "%x", 10)
```



  <a name="preferStringWriter-ref"></a>
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



  <a name="preferWriteByte-ref"></a>
## preferWriteByte

[
  **performance**
  **experimental** ]

Detects WriteRune calls with byte literal argument and reports to use WriteByte instead.





**Before:**
```go
w.WriteRune('\n')
```

**After:**
```go
w.WriteByte('\n')
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


  <a name="redundantSprint-ref"></a>
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



  <a name="regexpPattern-ref"></a>
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



  <a name="regexpSimplify-ref"></a>
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



  <a name="returnAfterHttpError-ref"></a>
## returnAfterHttpError

[
  **diagnostic**
  **experimental** ]

Detects suspicious http.Error call without following return.





**Before:**
```go
x + string(os.PathSeparator) + y
```

**After:**
```go
filepath.Join(x, y)
```



  <a name="ruleguard-ref"></a>
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

  `@ruleguard.failOnError` Determines the behavior when an error occurs while parsing ruleguard files.
If flag is not set, log error and skip rule files that contain an error.
If flag is set, the value must be a comma-separated list of error conditions.
* 'import': rule refers to a package that cannot be loaded.
* 'dsl':    gorule file does not comply with the ruleguard DSL. (default )

</li>
<li>

  `@ruleguard.rules` comma-separated list of gorule file paths. Glob patterns such as 'rules-*.go' may be specified (default )

</li>

</ul>


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



  <a name="sliceClear-ref"></a>
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



  <a name="sloppyLen-ref"></a>
## sloppyLen

[
  **style** ]

Detects usage of `len` when result is obvious or doesn't make sense.





**Before:**
```go
len(arr) <= 0
```

**After:**
```go
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



  <a name="sloppyTypeAssert-ref"></a>
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



  <a name="sortSlice-ref"></a>
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



  <a name="sprintfQuotedString-ref"></a>
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



  <a name="sqlQuery-ref"></a>
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



  <a name="stringConcatSimplify-ref"></a>
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



  <a name="stringXbytes-ref"></a>
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



  <a name="suspiciousSorting-ref"></a>
## suspiciousSorting

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



  <a name="switchTrue-ref"></a>
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



  <a name="syncMapLoadAndDelete-ref"></a>
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



  <a name="timeExprSimplify-ref"></a>
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



  <a name="tooManyResultsChecker-ref"></a>
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


  <a name="truncateCmp-ref"></a>
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



  <a name="typeDefFirst-ref"></a>
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



  <a name="unnecessaryDefer-ref"></a>
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



  <a name="unslice-ref"></a>
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



  <a name="valSwap-ref"></a>
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



  <a name="whyNoLint-ref"></a>
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



  <a name="wrapperFunc-ref"></a>
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


