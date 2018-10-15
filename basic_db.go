package main

import (
	"fmt"
	"sync"
)

// basicKVS is a very basic Key-Value Store.
type basicKVS struct {
	mutex sync.Mutex
	kvs   map[string]string
}

func newBasicKVS() *basicKVS {
	return &basicKVS{
		kvs: make(map[string]string),
	}
}

// store stores a key value pair.
func (db *basicKVS) store(k, v string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.kvs[k] = v
	return nil
}

// load returns the value associated with a given key.
func (db *basicKVS) load(k string) (string, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	v, ok := db.kvs[k]
	if !ok {
		return "", fmt.Errorf("basicKVS: unknown key '%s'", k)
	}
	return v, nil
}
