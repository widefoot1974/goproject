package cmap

import "sync"

type ConcurrentMap struct {
	sync.RWMutex
	items map[string]interface{}
}

func (cm *ConcurrentMap) Set(key string, value interface{}) {
	cm.Lock()
	defer cm.Unlock()
	cm.items[key] = value
}

func (cm *ConcurrentMap) Get(key string) (interface{}, bool) {
	cm.RLock()
	defer cm.RUnlock()
	value, ok := cm.items[key]
	return value, ok
}

func New() *ConcurrentMap {
	return &ConcurrentMap{
		items: make(map[string]interface{}),
	}
}
