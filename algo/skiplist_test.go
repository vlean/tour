package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSkipList(t *testing.T) {
	sp := NewSkipList()
	for i := 1; i < 50; i++ {
		score := rand.Intn(1e3)
		sp.Insert(score, i)
	}
	sp.Insert(618, 19)
	sp.Print()
	tmp := sp.Search(618)
	fmt.Println(tmp.Value)
	fmt.Println(sp)

}
