# Builtin shadow checker

## Description

This checker detects when
[predeclared identifiers](https://golang.org/ref/spec#Predeclared_identifiers)
shadowed in assignments.

Example:
```go
// shadowing len function
len := 10

// ...
// some code
// ...

// compile error: cannot call non-function len
len(sliceIdent)
```

### Rationale

Avoid situations when using this identifiers like local variables.
Later you can add code trying to access this identifier like a language identifier.
It can take you some time before you figure it out.