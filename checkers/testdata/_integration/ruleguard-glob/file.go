package example

import (
	"os"
	"path/filepath"
	"sync"
)

func separator() string {
	return string(os.PathSeparator)
}

func join(p ...string) string {
	return filepath.Join(p...)
}

var mu sync.RWMutex

func badlock() {
	mu.Lock()
	defer mu.RUnlock()
	println("foo")
}
