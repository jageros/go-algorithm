package main

import "fmt"

type lNode struct {
	data int
	next *lNode
}

func (l *lNode) newNextNode(data int) {
	l.next = &lNode{data:data}
}

func (l *lNode) print() {
	fmt.Printf("List len = %d\n", l.data)
	p := l.next
	for p != nil {
		fmt.Printf("%d ", p.data)
		p = p.next
	}
	fmt.Println()
}

func newList() *lNode {
	l := &lNode{}
	p := l
	for i := 0; i < 100; i += 10 {
		p.newNextNode(i)
		p = p.next
		l.data += 1
	}
	return l
}