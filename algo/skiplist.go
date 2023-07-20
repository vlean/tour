package main

import (
	"fmt"
	"math/rand"
)

const (
	MaxSkipLevel = 30
	SkipRate     = 0.25
)

func init() {
}

type SkipListNode struct {
	Next  []*SkipListNode
	Score int
	Value int
}

type SkipList struct {
	Head   *SkipListNode
	Length int
	Level  int
}

func NewSkipList() *SkipList {
	return &SkipList{
		Head: &SkipListNode{
			Next:  make([]*SkipListNode, MaxSkipLevel),
			Score: 0,
			Value: 0,
		},
		Length: 0,
		Level:  0,
	}
}

func (s *SkipList) Insert(score int, value int) {
	var next *SkipListNode
	prev := s.Head

	update := make([]*SkipListNode, MaxSkipLevel)

	for i := s.Level - 1; i >= 0; i-- { // fix: 2
		next = prev.Next[i]
		for next != nil && next.Score < score { // 替换值
			prev, next = next, next.Next[i]
		}
		update[i] = prev
	}

	if next != nil && next.Score == score {
		next.Value = value
		return
	}

	l := level()
	if l > s.Level { // fix: 1
		l = s.Level + 1
		s.Level = l
		update[s.Level-1] = s.Head
	}

	node := &SkipListNode{
		Next:  make([]*SkipListNode, l),
		Score: score,
		Value: value,
	}
	// 替换指针指向
	for i := 0; i < l; i++ {
		node.Next[i] = update[i].Next[i]
		update[i].Next[i] = node
	}
	s.Length++
}

func level() int {
	n := 0
	for {
		n++
		if rand.Int()%4 != 0 {
			continue
		}
		break
	}
	return n
}

func (s *SkipList) Delete(score int) {

}

func (s *SkipList) Search(score int) *SkipListNode {
	h := s.Head

LOOP:
	lv := s.Level
	if lv > len(h.Next) {
		lv = len(h.Next)
	}
	for i := lv - 1; i >= 0; i-- {
		node := h.Next[i]
		if node == nil {
			fmt.Println("current score:", h.Score, "node score: nil")
			continue
		}
		if node.Score > score {
			fmt.Println("current score:", h.Score, "node score:", node.Score)
			continue
		}
		if node.Score == score { // 替换值
			fmt.Println("find")
			return node
		}
		fmt.Println("current score:", h.Score, "node score:", node.Score)
		h = node
		goto LOOP
	}
	return nil
}

func (s *SkipList) Print() {
	vs := make([][]string, s.Length+1)
	next := s.Head
	for j := 0; j < s.Length+1; j++ {
		lv := make([]string, s.Level)
		rg := s.Level
		if rg > len(next.Next) {
			rg = len(next.Next)
		}
		fmt.Printf("[%03d-%d](%d)\t\t", next.Score, next.Value, len(next.Next))
		for i := 0; i < rg; i++ {
			nt := next.Next[i]
			if nt == nil {
				fmt.Printf("nil-nil\t\t")
				lv[i] = "nil-nil"
			} else {
				lv[i] = fmt.Sprintf("%03d-%d\t\t", next.Next[i].Score, next.Next[i].Value)
				fmt.Print(lv[i])
			}
		}
		fmt.Println()
		vs[j] = lv
		next = next.Next[0]
	}

}
