Detects type switches that cab benefit from type guard clause.

**Before:**
```go
func f() int {
	type point struct { x, y int }
	var v interface{} = point{1, 2}

	switch v.(type) {
	case int:
		return v.(int)
	case point:
		return v.(point).x + v.(point).y
	default:
		return 0
	}
}
```

**After:**
```go
func f() int {
	type point struct { x, y int }
	var v interface{} = point{1, 2}

	switch v := v.(type) {
	case int:
		return v
	case point:
		return v.x + v.y
	default:
		return 0
	}
}
```
