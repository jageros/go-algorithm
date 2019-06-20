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

func (l *lNode) insertReverseOrder() {
	if l.next == nil {
		return
	}
	p := l.next.next
	l.next.next = nil
	for p != nil {
		next := p.next
		p.next = l.next
		l.next = p
		p = next
	}
}

func main() {
	l := &lNode{}
	p := l
	for i := 0; i < 300; i += 10 {
		p.newNextNode(i)
		p = p.next
	}
	l.print()
	l.insertReverseOrder()
	l.print()
}