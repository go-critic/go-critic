Finds usage of unnamed results, skips (T, error) and (T, bool) patterns.

**Before:**
```go
func f() (float64, float64)
```

**After:**
```go
func f() (x, y float64)
```
