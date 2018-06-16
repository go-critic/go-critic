For functions with multiple return values, detects unnamed results
that do not match `(T, error)` or `(T, bool)` pattern.

**Before:**
```go
func f() (float64, float64)
```

**After:**
```go
func f() (x, y float64)
```
