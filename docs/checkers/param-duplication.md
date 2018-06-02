Detects if function parameters could be combined by type and suggest the way to do it.

**before:**
```go
func foo(a, b int, c, d int, e, f int, g int) {}
```

**after:**
```go
func foo(a, b, c, d, e, f, g int) {}
```
