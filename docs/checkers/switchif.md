Detects switch statements that could be better written as if statements.

**Before:**
```go
switch x := x.(type) {
case int:
     ...
}
return 0

```

**After:**
```go
if x, ok := x.(int); ok {
   ...
}
```
