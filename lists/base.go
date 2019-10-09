package main

import (
	"fmt"
	"math/rand"
	"time"
)

type lHeadNode struct {
	p *lNode
	length int
}

type lNode struct {
	data interface{}
	next *lNode
}

func (l *lNode) forEachNode(f func(l *lNode) bool) {
	ctn := f(l)
	if !ctn {
		return
	}
	if l.next == nil {
		return
	}
	l.next.forEachNode(f)
}

func (l *lHeadNode) print() {
	fmt.Printf("List len = %d\n", l.length)
	p := l.p
	for p != nil {
		fmt.Printf("%v ", p.data)
		p = p.next
	}
	fmt.Println()
}

func newList(maxFigure, length int) *lHeadNode {
	l := &lHeadNode{
		p: &lNode{},
		length:length,
	}
	p := l.p
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		num := rand.Int()%maxFigure + 1
		p.data = num
		if i+1 < length {
			p.next = &lNode{}
			p = p.next
		}
	}
	return l
}
