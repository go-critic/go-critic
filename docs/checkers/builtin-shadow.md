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