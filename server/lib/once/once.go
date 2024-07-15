package once

import (
	"reflect"
	"sync"
)

var onceFuncs = make(map[uintptr]*sync.Once)

// Once will execute unique functions f only once, used for lazy
// loading dependancies
func Once(f func()) {
	ptr := reflect.ValueOf(f).Pointer()
	once, exists := onceFuncs[ptr]
	if !exists {
		once = new(sync.Once)
		onceFuncs[ptr] = once
	}
	once.Do(f)
}
