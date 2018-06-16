Detects unused params and suggests to name them as `_` (underscore).

**Before:**
```go
func f(a int, b float64) // b isn't used inside function body
```

**After:**
```go
func f(a int, _ float64) // everything is cool
```
