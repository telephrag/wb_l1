package main

import (
	"fmt"
	"sync"
)

// More or less a copy of an `OrdersCache` type I used in `L0`.
type ConcMap struct {
	data map[string]any

	// Not all operations require full lock to be thread safe so, `sync.RWMutex` is optimal.
	mu sync.RWMutex
}

// Initializes internal fields of an object of the type `ConcMap`.
func (cm *ConcMap) Init() (self *ConcMap) {
	cm.data = make(map[string]interface{})
	cm.mu = sync.RWMutex{}

	return cm
}

// Returns records count of the underlying map.
func (cm *ConcMap) Len() int {
	return len(cm.data)
}

// Returns value of underlying map at key
func (cm *ConcMap) Get(k string) (any, bool) {
	// Make map read-only for a time we read to avoid reading invalid data.
	// Reads elsewhere are still possible.
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	v, ok := cm.data[k]
	return v, ok
}

func (cm *ConcMap) Set(k string, v any) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.data[k] = v
}

func main() {
	// Personally, I would rather have creation of an object look ugly
	// than expose fields that are not meant to be exposed.
	m := (&ConcMap{}).Init()

	go func() {
		m.Set("a", "str")
	}()

	go func() {
		m.Set("b", 123)
	}()
	fmt.Println(m.Get("a"))

	go func() {
		m.Set("c", struct{}{})
	}()
	fmt.Println(m.Get("b"))

	fmt.Println(m.Get("c"))

	// `go run -race main.go` to see that there are no dataraces.
}
