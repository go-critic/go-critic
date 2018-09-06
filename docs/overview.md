## Checks overview

This page describes checks supported by [go-critic](https://github.com/go-critic/go-critic) linter.

[//]: # (This is generated file, please don't edit it yourself.)

## Checkers:

<table>
  <tr>
    <th>Name</th>
    <th>Short description</th>
  </tr>
      <tr>
        <td><a href="#appendCombine-ref">appendCombine</a></td>
        <td>Detects `append` chains to the same slice that can be done in a single `append` call</td>
      </tr>
      <tr>
        <td><a href="#builtinShadow-ref">builtinShadow</a></td>
        <td>Detects when predeclared identifiers shadowed in assignments</td>
      </tr>
      <tr>
        <td><a href="#flagDeref-ref">flagDeref</a></td>
        <td>Detects immediate dereferencing of `flag` package pointers.</td>
      </tr>
      <tr>
        <td><a href="#ifElseChain-ref">ifElseChain</a></td>
        <td>Detects repeated if-else statements and suggests to replace them with switch statement</td>
      </tr>
      <tr>
        <td><a href="#paramTypeCombine-ref">paramTypeCombine</a></td>
        <td>Detects if function parameters could be combined by type and suggest the way to do it</td>
      </tr>
      <tr>
        <td><a href="#rangeExprCopy-ref">rangeExprCopy</a></td>
        <td>Detects expensive copies of `for` loop range expressions</td>
      </tr>
      <tr>
        <td><a href="#rangeValCopy-ref">rangeValCopy</a></td>
        <td>Detects loops that copy big objects during each iteration</td>
      </tr>
      <tr>
        <td><a href="#singleCaseSwitch-ref">singleCaseSwitch</a></td>
        <td>Detects switch statements that could be better written as if statements</td>
      </tr>
      <tr>
        <td><a href="#switchTrue-ref">switchTrue</a></td>
        <td>Detects switch-over-bool statements that use explicit `true` tag value</td>
      </tr>
      <tr>
        <td><a href="#typeSwitchVar-ref">typeSwitchVar</a></td>
        <td>Detects type switches that can benefit from type guard clause with variable</td>
      </tr>
      <tr>
        <td><a href="#typeUnparen-ref">typeUnparen</a></td>
        <td>Detects unneded parenthesis inside type expressions and suggests to remove them</td>
      </tr>
      <tr>
        <td><a href="#underef-ref">underef</a></td>
        <td>Detects dereference expressions that can be omitted</td>
      </tr>
      <tr>
        <td><a href="#unslice-ref">unslice</a></td>
        <td>Detects slice expressions that can be simplified to sliced expression itself</td>
      </tr>
</table>

**Experimental:**

<table>
  <tr>
    <th>Name</th>
    <th>Short description</th>
  </tr>
      <tr>
        <td><a href="#appendAssign-ref">appendAssign</a></td>
        <td>Detects suspicious append result assignments</td>
      </tr>
      <tr>
        <td><a href="#assignOp-ref">assignOp</a></td>
        <td>Detects assignments that can be simplified by using assignment operators</td>
      </tr>
      <tr>
        <td><a href="#boolExprSimplify-ref">boolExprSimplify</a></td>
        <td>Detects bool expressions that can be simplified</td>
      </tr>
      <tr>
        <td><a href="#boolFuncPrefix-ref">boolFuncPrefix</a> :nerd_face:</td>
        <td>Detects function returning only bool and suggests to add Is/Has/Contains prefix to it's name</td>
      </tr>
      <tr>
        <td><a href="#busySelect-ref">busySelect</a></td>
        <td>Detects default statement inside a select without a sleep that might waste a CPU time</td>
      </tr>
      <tr>
        <td><a href="#captLocal-ref">captLocal</a></td>
        <td>Detects capitalized names for local variables</td>
      </tr>
      <tr>
        <td><a href="#caseOrder-ref">caseOrder</a></td>
        <td>Detects erroneous case order inside switch statements</td>
      </tr>
      <tr>
        <td><a href="#commentedOutCode-ref">commentedOutCode</a></td>
        <td>Detects commented-out code inside function bodies</td>
      </tr>
      <tr>
        <td><a href="#deadCodeAfterLogFatal-ref">deadCodeAfterLogFatal</a></td>
        <td>Detects dead code that follow panic/fatal logging</td>
      </tr>
      <tr>
        <td><a href="#defaultCaseOrder-ref">defaultCaseOrder</a></td>
        <td>Detects when default case in switch isn't on 1st or last position</td>
      </tr>
      <tr>
        <td><a href="#deferInLoop-ref">deferInLoop</a></td>
        <td>Detects defer in loop and warns that it will not be executed till the end of function's scope</td>
      </tr>
      <tr>
        <td><a href="#deprecatedComment-ref">deprecatedComment</a></td>
        <td>Detects malformed "deprecated" doc-comments</td>
      </tr>
      <tr>
        <td><a href="#docStub-ref">docStub</a></td>
        <td>Detects comments that silence go lint complaints about doc-comment</td>
      </tr>
      <tr>
        <td><a href="#dupArg-ref">dupArg</a></td>
        <td>Detects suspicious duplicated arguments</td>
      </tr>
      <tr>
        <td><a href="#dupBranchBody-ref">dupBranchBody</a></td>
        <td>Detects duplicated branch bodies inside conditional statements</td>
      </tr>
      <tr>
        <td><a href="#dupCase-ref">dupCase</a></td>
        <td>Detects duplicated case clauses inside switch statements</td>
      </tr>
      <tr>
        <td><a href="#dupSubExpr-ref">dupSubExpr</a></td>
        <td>Detects suspicious duplicated sub-expressions</td>
      </tr>
      <tr>
        <td><a href="#elseif-ref">elseif</a> :nerd_face:</td>
        <td>Detects else with nested if statement that can be replaced with else-if</td>
      </tr>
      <tr>
        <td><a href="#emptyFallthrough-ref">emptyFallthrough</a></td>
        <td>Detects fallthrough that can be avoided by using multi case values</td>
      </tr>
      <tr>
        <td><a href="#emptyFmt-ref">emptyFmt</a></td>
        <td>Detects usages of formatting functions without formatting arguments</td>
      </tr>
      <tr>
        <td><a href="#evalOrder-ref">evalOrder</a></td>
        <td>Detects potentially unsafe dependencies on evaluation order</td>
      </tr>
      <tr>
        <td><a href="#floatCompare-ref">floatCompare</a></td>
        <td>Detects fragile float variables comparisons</td>
      </tr>
      <tr>
        <td><a href="#hugeParam-ref">hugeParam</a></td>
        <td>Detects params that incur excessive amount of copying</td>
      </tr>
      <tr>
        <td><a href="#importPackageName-ref">importPackageName</a></td>
        <td>Detects when imported package names are unnecessary renamed</td>
      </tr>
      <tr>
        <td><a href="#importShadow-ref">importShadow</a></td>
        <td>Detects when imported package names shadowed in assignments</td>
      </tr>
      <tr>
        <td><a href="#indexOnlyLoop-ref">indexOnlyLoop</a></td>
        <td>Detects for loops that can benefit from rewrite to range loop</td>
      </tr>
      <tr>
        <td><a href="#initClause-ref">initClause</a></td>
        <td>Detects non-assignment statements inside if/switch init clause</td>
      </tr>
      <tr>
        <td><a href="#longChain-ref">longChain</a></td>
        <td>Detects repeated expression chains and suggest to refactor them</td>
      </tr>
      <tr>
        <td><a href="#methodExprCall-ref">methodExprCall</a> :nerd_face:</td>
        <td>Detects method expression call that can be replaced with a method call</td>
      </tr>
      <tr>
        <td><a href="#namedConst-ref">namedConst</a></td>
        <td>Detects literals that can be replaced with defined named const</td>
      </tr>
      <tr>
        <td><a href="#nestingReduce-ref">nestingReduce</a></td>
        <td>Finds where nesting level could be reduced</td>
      </tr>
      <tr>
        <td><a href="#nilValReturn-ref">nilValReturn</a></td>
        <td>Detects return statements those results evaluate to nil</td>
      </tr>
      <tr>
        <td><a href="#ptrToRefParam-ref">ptrToRefParam</a></td>
        <td>Detects input and output parameters that have a type of pointer to referential type</td>
      </tr>
      <tr>
        <td><a href="#regexpMust-ref">regexpMust</a></td>
        <td>Detects `regexp.Compile*` that can be replaced with `regexp.MustCompile*`</td>
      </tr>
      <tr>
        <td><a href="#sloppyLen-ref">sloppyLen</a></td>
        <td>Detects usage of `len` when result is obvious or doesn't make sense</td>
      </tr>
      <tr>
        <td><a href="#sqlRowsClose-ref">sqlRowsClose</a></td>
        <td>Detects uses of *sql.Rows without call Close method</td>
      </tr>
      <tr>
        <td><a href="#stdExpr-ref">stdExpr</a></td>
        <td>Detects constant expressions that can be replaced by a stdlib const</td>
      </tr>
      <tr>
        <td><a href="#unexportedCall-ref">unexportedCall</a> :nerd_face:</td>
        <td>Detects calls of unexported method from unexported type outside that type</td>
      </tr>
      <tr>
        <td><a href="#unlambda-ref">unlambda</a></td>
        <td>Detects function literals that can be simplified</td>
      </tr>
      <tr>
        <td><a href="#unnamedResult-ref">unnamedResult</a></td>
        <td>Detects unnamed results that may benefit from names</td>
      </tr>
      <tr>
        <td><a href="#unnecessaryBlock-ref">unnecessaryBlock</a></td>
        <td>Detects unnecessary braced statement blocks</td>
      </tr>
      <tr>
        <td><a href="#blankParam-ref">blankParam</a></td>
        <td>Detects blank params and suggests to name them as `_` (underscore)</td>
      </tr>
      <tr>
        <td><a href="#yodaStyleExpr-ref">yodaStyleExpr</a> :nerd_face:</td>
        <td>Detects Yoda style expressions that suggest to replace them</td>
      </tr>
</table>



<a name="appendAssign-ref"></a>
## appendAssign
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



<a name="assignOp-ref"></a>
## assignOp
Detects assignments that can be simplified by using assignment operators.



**Before:**
```go
x = x * 2
```

**After:**
```go
x *= 2
```



<a name="boolExprSimplify-ref"></a>
## boolExprSimplify
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



<a name="boolFuncPrefix-ref"></a>
## boolFuncPrefix
Detects function returning only bool and suggests to add Is/Has/Contains prefix to it's name.



**Before:**
```go
func Enabled() bool
```

**After:**
```go
func IsEnabled() bool
```


`boolFuncPrefix` is very opinionated.
<a name="builtinShadow-ref"></a>
## builtinShadow
Detects when predeclared identifiers shadowed in assignments.



**Before:**
```go
len := 10
println(len)
```

**After:**
```go
length := 10 // Changed variable name
println(length)
```


`builtinShadow` is syntax-only checker (fast).
<a name="busySelect-ref"></a>
## busySelect
Detects default statement inside a select without a sleep that might waste a CPU time.



**Before:**
```go
for {
	select {
	case <-ch:
		// ...
	default:
		// will waste CPU time
	}
}
```

**After:**
```go
for {
	select {
	case <-ch:
		// ...
	default:
		time.Sleep(100 * time.Millisecond)
	}
}
```



<a name="captLocal-ref"></a>
## captLocal
Detects capitalized names for local variables.



**Before:**
```go
func f(IN int, OUT *int) (ERR error) {}
```

**After:**
```go
func f(in int, out *int) (err error) {}
```


`captLocal` is syntax-only checker (fast).
<a name="caseOrder-ref"></a>
## caseOrder
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



<a name="commentedOutCode-ref"></a>
## commentedOutCode
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



<a name="deadCodeAfterLogFatal-ref"></a>
## deadCodeAfterLogFatal
Detects dead code that follow panic/fatal logging.



**Before:**
```go
log.Fatal("exits function")
return
```

**After:**
```go
log.Fatal("exits function")
```



<a name="defaultCaseOrder-ref"></a>
## defaultCaseOrder
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


`defaultCaseOrder` is syntax-only checker (fast).
<a name="deferInLoop-ref"></a>
## deferInLoop
Detects defer in loop and warns that it will not be executed till the end of function's scope.



**Before:**
```go
for i := range [10]int{} {
	defer f(i) // will be executed only at the end of func
}
```

**After:**
```go
for i := range [10]int{} {
	func(i int) {
		defer f(i)
	}(i)
}
```



<a name="deprecatedComment-ref"></a>
## deprecatedComment
Detects malformed "deprecated" doc-comments.



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


`deprecatedComment` is syntax-only checker (fast).
<a name="docStub-ref"></a>
## docStub
Detects comments that silence go lint complaints about doc-comment.



**Before:**
```go
// Foo ...
func Foo() {
}
```

**After:**
```go
func Foo() {
}
```

You can either remove a comment to let go lint find it or change stub to useful comment.
This checker makes it easier to detect stubs, the action is up to you.
`docStub` is syntax-only checker (fast).
<a name="dupArg-ref"></a>
## dupArg
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



<a name="dupSubExpr-ref"></a>
## dupSubExpr
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


`elseif` is very opinionated.
<a name="emptyFallthrough-ref"></a>
## emptyFallthrough
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



<a name="emptyFmt-ref"></a>
## emptyFmt
Detects usages of formatting functions without formatting arguments.



**Before:**
```go
fmt.Sprintf("whatever")
fmt.Errorf("wherever")
```

**After:**
```go
fmt.Sprint("whatever")
errors.New("wherever")
```



<a name="evalOrder-ref"></a>
## evalOrder
Detects potentially unsafe dependencies on evaluation order.



**Before:**
```go
return mayModifySlice(&xs), xs[0]
```

**After:**
```go
// A)
v := mayModifySlice(&xs)
return v, xs[0]
// B)
v := xs[0]
return mayModifySlice(&xs), v
```



<a name="flagDeref-ref"></a>
## flagDeref
Detects immediate dereferencing of `flag` package pointers..



**Before:**
```go
b := *flag.Bool("b", false, "b docs")
```

**After:**
```go
var b bool
flag.BoolVar(&b, "b", false, "b docs")
```

Dereferencing returned pointers will lead to hard to find errors
where flag values are not updated after flag.Parse().
`flagDeref` is syntax-only checker (fast).
<a name="floatCompare-ref"></a>
## floatCompare
Detects fragile float variables comparisons.



**Before:**
```go
// x and y are floats
return x == y
```

**After:**
```go
// x and y are floats
return math.Abs(x - y) < eps
```



<a name="hugeParam-ref"></a>
## hugeParam
Detects params that incur excessive amount of copying.



**Before:**
```go
func f(x [1024]int) {}
```

**After:**
```go
func f(x *[1024]int) {}
```



<a name="ifElseChain-ref"></a>
## ifElseChain
Detects repeated if-else statements and suggests to replace them with switch statement.



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

Permits single else or else-if; repeated else-if or else + else-if
will trigger suggestion to use switch statement.
See [EffectiveGo#switch](https://golang.org/doc/effective_go.html#switch).
`ifElseChain` is syntax-only checker (fast).
<a name="importPackageName-ref"></a>
## importPackageName
Detects when imported package names are unnecessary renamed.



**Before:**
```go
import lint "github.com/go-critic/go-critic/lint"
```

**After:**
```go
import "github.com/go-critic/go-critic/lint"
```



<a name="importShadow-ref"></a>
## importShadow
Detects when imported package names shadowed in assignments.



**Before:**
```go
// "path/filepath" is imported.
func myFunc(filepath string) {
}
```

**After:**
```go
func myFunc(filename string) {
}
```



<a name="indexOnlyLoop-ref"></a>
## indexOnlyLoop
Detects for loops that can benefit from rewrite to range loop.

Suggests to use for key, v := range container form.

**Before:**
```go
for i := range files {
	if files[i] != nil {
		files[i].Close()
	}
}
```

**After:**
```go
for _, f := range files {
	if f != nil {
		f.Close()
	}
}
```



<a name="initClause-ref"></a>
## initClause
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


`initClause` is syntax-only checker (fast).
<a name="longChain-ref"></a>
## longChain
Detects repeated expression chains and suggest to refactor them.



**Before:**
```go
a := q.w.e.r.t + 1
b := q.w.e.r.t + 2
c := q.w.e.r.t + 3
v := (a + xs[i+1]) + (b + xs[i+1]) + (c + xs[i+1])
```

**After:**
```go
x := xs[i+1]
qwert := q.w.e.r.t
a := qwert + 1
b := qwert + 2
c := qwert + 3
v := (a + x) + (b + x) + (c + x)
```



<a name="methodExprCall-ref"></a>
## methodExprCall
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


`methodExprCall` is very opinionated.
<a name="namedConst-ref"></a>
## namedConst
Detects literals that can be replaced with defined named const.



**Before:**
```go
// pos has type of token.Pos.
return pos != 0
```

**After:**
```go
return pos != token.NoPos
```



<a name="nestingReduce-ref"></a>
## nestingReduce
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



<a name="nilValReturn-ref"></a>
## nilValReturn
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
	return nil
}
```



<a name="paramTypeCombine-ref"></a>
## paramTypeCombine
Detects if function parameters could be combined by type and suggest the way to do it.



**Before:**
```go
func foo(a, b int, c, d int, e, f int, g int) {}
```

**After:**
```go
func foo(a, b, c, d, e, f, g int) {}
```


`paramTypeCombine` is syntax-only checker (fast).
<a name="ptrToRefParam-ref"></a>
## ptrToRefParam
Detects input and output parameters that have a type of pointer to referential type.



**Before:**
```go
func f(m *map[string]int) (ch *chan *int)
```

**After:**
```go
func f(m map[string]int) (ch chan *int)
```

Slices are not as referential as maps or channels, but it's usually
better to return them by value rather than modyfing them by pointer.

<a name="rangeExprCopy-ref"></a>
## rangeExprCopy
Detects expensive copies of `for` loop range expressions.

Suggests to use pointer to array to avoid the copy using `&` on range expression.

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

See Go issue for details: https://github.com/golang/go/issues/15812

<a name="rangeValCopy-ref"></a>
## rangeValCopy
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



<a name="regexpMust-ref"></a>
## regexpMust
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
Detects switch statements that could be better written as if statements.



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


`singleCaseSwitch` is syntax-only checker (fast).
<a name="sloppyLen-ref"></a>
## sloppyLen
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


`sloppyLen` is syntax-only checker (fast).
<a name="sqlRowsClose-ref"></a>
## sqlRowsClose
Detects uses of *sql.Rows without call Close method.



**Before:**
```go
rows, _ := db.Query( /**/ )
for rows.Next {
}
```

**After:**
```go
rows, _ := db.Query( /**/ )
for rows.Next {
}
rows.Close()
```



<a name="stdExpr-ref"></a>
## stdExpr
Detects constant expressions that can be replaced by a stdlib const.



**Before:**
```go
intBytes := make([]byte, unsafe.Sizeof(0))
maxVal := 1<<7 - 1
```

**After:**
```go
intBytes := make([]byte, bits.IntSize)
maxVal := math.MaxInt8
```



<a name="switchTrue-ref"></a>
## switchTrue
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


`switchTrue` is syntax-only checker (fast).
<a name="typeSwitchVar-ref"></a>
## typeSwitchVar
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
Detects unneded parenthesis inside type expressions and suggests to remove them.



**Before:**
```go
type foo [](func([](func())))
```

**After:**
```go
type foo []func([]func())
```


`typeUnparen` is syntax-only checker (fast).
<a name="underef-ref"></a>
## underef
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



<a name="unexportedCall-ref"></a>
## unexportedCall
Detects calls of unexported method from unexported type outside that type.



**Before:**
```go
func baz(f foo) {
	fo.bar()
}
```

**After:**
```go
func baz(f foo) {
	fo.Bar() // Made method exported
}
```


`unexportedCall` is very opinionated.
<a name="unlambda-ref"></a>
## unlambda
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
Detects unnamed results that may benefit from names.



**Before:**
```go
func f() (float64, float64)
```

**After:**
```go
func f() (x, y float64)
```



<a name="unnecessaryBlock-ref"></a>
## unnecessaryBlock
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


`unnecessaryBlock` is syntax-only checker (fast).
<a name="unslice-ref"></a>
## unslice
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



<a name="blankParam-ref"></a>
## blankParam
Detects blank params and suggests to name them as `_` (underscore).



**Before:**
```go
func f(a int, b float64) // b isn't used inside function body
```

**After:**
```go
func f(a int, _ float64) // everything is cool
```



<a name="yodaStyleExpr-ref"></a>
## yodaStyleExpr
Detects Yoda style expressions that suggest to replace them.



**Before:**
```go
return nil != ptr
```

**After:**
```go
return ptr != nil
```


`yodaStyleExpr` is very opinionated.
