Detects function returning only bool and suggests to add Is/Has/Contains prefix to it's name.

**Before:**
```go
func Enabled() bool
```

**After:**
```go
func IsEnabled() bool
```
