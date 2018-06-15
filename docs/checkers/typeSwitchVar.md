Detects type switches that can benefit from type guard clause with variable.

**Before:**
```go
switch v.(type) {
case int:
	return v.(int)
case point:
	return v.(point).x + v.(point).y
default:
	return 0
}
```

**After:**
```go
switch v := v.(type) {
case int:
	return v
case point:
	return v.x + v.y
default:
	return 0
}
```
