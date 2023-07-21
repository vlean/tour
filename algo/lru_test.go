package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRU(t *testing.T) {
	cache := NewLRUCache(10)
	for i := 0; i < 15; i++ {
		cache.Put(fmt.Sprintf("%3d", rand.Intn(1e3)), i)
		cache.Print()
	}
	cache.Put("100", 100)
	cache.Print()
	val, _ := cache.Get("100")
	assert.Equal(t, val, 100)

	cache.Put("101", 101)
	cache.Print()
	cache.Put("100", 100)
	cache.Print()
}
