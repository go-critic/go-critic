Detects comments that aim to silence go lint complaints about exported symbol not having a doc-comment.

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
