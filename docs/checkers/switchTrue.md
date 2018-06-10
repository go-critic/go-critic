Detects switch-over-bool statements that use explicit `true` tag value.

**Before:**
```go
switch true {
case x > y:
	// ...
}
```

**After:**
```go
switch {
case x > y:
	// ...
}
```
