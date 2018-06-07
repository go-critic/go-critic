Detects immediate dereferencing of `flag` package pointers.
Suggests using `XxxVar` functions to achieve desired effect.

**Before:**
```go
b := *flag.Bool("b", false, "b docs")
```

**After:**
```go
var b bool
flag.BoolVar(&b, "b", false, "b docs")
```
