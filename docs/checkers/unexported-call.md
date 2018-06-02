Finds calls of unexported method from unexported type outside that type.

**Before:**
```go
type foo struct{}

func (f foo) f() int { return 1 }
func f() {
	var fo foo
	fo.f()
}

```

**After:**
```go
type foo struct{}

func (f foo) unexported() int { return 1 }
func f() {
	var fo foo
	fo.F()
}
```
