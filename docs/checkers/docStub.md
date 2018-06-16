Detects comments that silence go lint complaints about doc-comment.

**Before:**
```go
// Foo ...
func Foo() {
     // ...
}

```

**After:**
```go
func Foo() {
     // ...
}
```

> You can either remove a comment to let go lint find it or change stub to useful comment.
> This checker makes it easier to detect stubs, the action is up to you.
