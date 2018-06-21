Detecs implementation assert pattern in code files and suggests 
to move them to test files/packages to prevent redundant dependency.

**Before:**
```go
// code.go
var _ pkg.MyInterface = (*MyStruct)(nil)
```

**After:**
```go
// code_test.go
var _ pkg.MyInterface = (*MyStruct)(nil)
```