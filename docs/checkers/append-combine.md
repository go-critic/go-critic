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
