Detects expressions with C style field selection and suggest Go style correction.

**Before:**
```go
(*k).field = 5
_ := (*a)[5] // a is slice
```

**After:**
```go
k.field = 5
_ := a[5]
```
