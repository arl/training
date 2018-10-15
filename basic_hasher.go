package main

import (
	"strconv"
)

type basicHasher struct {
	gen <-chan uint64 // gen generates unique identifiers
}

func newBasicHasher() *basicHasher {
	// generate unique ids forever
	gen := make(chan uint64)
	go func() {
		var i uint64
		for {
			i++
			gen <- i
		}
	}()
	return &basicHasher{
		gen: gen,
	}
}

func (bh *basicHasher) Hash(string) string {
	id := <-bh.gen
	return strconv.FormatUint(id, 10)
}
