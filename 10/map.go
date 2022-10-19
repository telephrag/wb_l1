package main

import (
	"fmt"
	"sync"
)

type ConcMap struct {
	data map[int64][]float64

	// Not all operations require full lock to be thread safe so, `sync.RWMutex` is optimal.
	mu sync.RWMutex
}

// Initializes internal fields of an object of the type `ConcMap`.
func (cm *ConcMap) Init() (self *ConcMap) {
	cm.data = make(map[int64][]float64)
	cm.mu = sync.RWMutex{}

	return cm
}

// Returns records count of the underlying map.
func (cm *ConcMap) Len() int {
	return len(cm.data)
}

// Returns value of underlying map at key
func (cm *ConcMap) Get(k int64) ([]float64, bool) {
	// Make map read-only for a time we read to avoid reading invalid data.
	// Reads elsewhere are still possible.
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	v, ok := cm.data[k]
	return v, ok
}

func (cm *ConcMap) AppendAt(k int64, v float64) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if _, ok := cm.data[k]; ok {
		cm.data[k] = append(cm.data[k], v)
	} else {
		cm.data[k] = make([]float64, 1)
		cm.data[k][0] = v
	}
}

func (cm *ConcMap) Print() {
	for k, v := range cm.data {
		fmt.Printf("%d: %v\n", k, v)
	}
}
