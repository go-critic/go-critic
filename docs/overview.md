## Overview

TODO: add some useful content here

## Checkers

Name | Experimental
-----|------------
   big-copy | false
   builtin-shadow | false
   comments | false
   elseif | false
   long-chain | true
   param-duplication | false
   param-name | false
   parenthesis | false
   switchif | false
   type-guard | false
   underef | false
   unexported-call | false
   unslice | false
## big-copy
Detects loops that copy big objects during each iteration.
Suggests to use index access or take address and make use pointer instead.

```go
xs := make([][1024]byte, length)
for _, x := range xs {
	// Loop body.
}
```

```go
xs := make([][1024]byte, length)
for i := range xs {
	x := &amp;xs[i]
	// Loop body.
}
```

## builtin-shadow
Detects when
[predeclared identifiers](https://golang.org/ref/spec#Predeclared_identifiers)
shadowed in assignments.

Avoid situations when using this identifiers like local variables.
Later you can add code trying to access this identifier like a language identifier.
It can take you some time before you figuring it out.

**Before**
```go
func main() {
    // shadowing len function
    len := 10
    println(len)
}
```

**After**
```go
func main() {
    // change identificator name
    length := 10
    println(length)
}
```
## comments
Detects comments that aim to silence go lint complaints about exported symbol not having a doc-comment.

**Before:**
```go
// Foo ...
func Foo() {
     ...
}

```

**After:**
```go
// Foo useful comment for Foo 
func Foo() {
     ...
}

```

## elseif
Detects repeated if-else statements and suggests to replace them with switch statement.

Permits single else or else-if; repeated else-if or else &#43; else-if
will trigger suggestion to use switch statement.

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

## long-chain
Detects repeated expression chains and suggest to refactor them.

**Before:**
```go
a := q.w.e.r.t &#43; 1
b := q.w.e.r.t &#43; 2
c := q.w.e.r.t &#43; 3
v := (a&#43;xs[i&#43;1]) &#43; (b&#43;xs[i&#43;1]) &#43; (c&#43;xs[i&#43;1])
```

**After**
```go
x := xs[i&#43;1]
qwert := q.w.e.r.t
a := qwert &#43; 1
b := qwert &#43; 2
c := qwert &#43; 3
v := (a&#43;x) &#43; (b&#43;x) &#43; (c&#43;x)
```

**Experimental**

Gives false-positives for:
* Cases with re-assignment. See `$GOROOT/src/crypto/md5/md5block.go` for example.

## param-duplication
Detects if function parameters could be combined by type and suggest the way to do it.

**Before:**
```go
func foo(a, b int, c, d int, e, f int, g int) {}
```

**After:**
```go
func foo(a, b, c, d, e, f, g int) {}
```

## param-name
Detects potential issues in function parameter names.

Catches capitalized (exported) parameter names.
Suggests to replace them with non-capitalized versions.

**Before:**
```go
func f(IN int, OUT *int) (ERR error) {}
```

**After:**
```go
func f(in int, out *int) (err error) {}
```

## parenthesis
Detects unneded parenthesis inside type expressions and suggests to remove them.

**Before:**
```go
func foo() [](func([](func()))) {
     ...
}
```

**After:**
```go
func foo() []func([]func()) {
     ...
}
```


## switchif
Detects switch statements that could be better written as if statements.

**Before:**
```go
switch x := x.(type) {
case int:
     ...
}
```

**After:**
```go
if x, ok := x.(int); ok {
   ...
}
```

## type-guard
Detects type switches that cab benefit from type guard clause.

**Before:**
```go
func f() int {
	type point struct { x, y int }
	var v interface{} = point{1, 2}

	switch v.(type) {
	case int:
		return v.(int)
	case point:
		return v.(point).x &#43; v.(point).y
	default:
		return 0
	}
}
```

**After:**
```go
func f() int {
	type point struct { x, y int }
	var v interface{} = point{1, 2}

	switch v := v.(type) {
	case int:
		return v
	case point:
		return v.x &#43; v.y
	default:
		return 0
	}
}
```

## underef
Detects expressions with C style field selection and suggest Go style correction.

**Before:**
```go
(*k).field = 5
_ := (*a)[5] // only if a is array
```

**After:**
```go
k.field = 5
_ := a[5]
```

## unexported-call
Finds calls of unexported method from unexported type outside that type.

**Before:**
```go
type foo struct{}

func (f foo) bar() int { return 1 }
func baz() {
	var fo foo
	fo.bar()
}

```

**After:**
```go
type foo struct{}

func (f foo) Bar() int { return 1 } // now Bar is exported
func baz() {
	var fo foo
	fo.Bar()
}
```

## unslice
Detects slice expressions that can be simplified to sliced expression itself.

**Before:**
```go
f(s[:]) // s is string
copy(b[:], values...) // b is []byte
```

**After:**
```go
f(s)
copy(b, values...)
```
