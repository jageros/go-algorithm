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

func (l *lNode) reverseInsert() {
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

func (l *lNode) reverseOnSpot() {
	var pre *lNode
	cur := l.next
	for cur != nil {
		next := cur.next
		cur.next = pre
		pre = cur
		cur = next
	}
	l.next = pre
}

func (l *lNode) reverseRecursive() *lNode {
	if l.next == nil {
		return l
	}
	h := l.next.reverseRecursive()
	l.next.next = l
	l.next = nil
	return h
}

func (l *lNode) reverse() {
	l.next = l.next.reverseRecursive()
}

func (l *lNode) reversePrint() {
	if l.next == nil {
		fmt.Print(l.data, " ")
		return
	}
	l.next.reversePrint()
	fmt.Print(l.data, " ")
}

func main() {
	l := &lNode{}
	p := l
	for i := 0; i < 100; i += 10 {
		p.newNextNode(i)
		p = p.next
		l.data += 1
	}
	l.print()
	l.reverseInsert()
	l.print()
	l.reverseOnSpot()
	l.print()
	l.reverse()

	l.print()
	fmt.Println("List len =", l.data)
	l.next.reversePrint()
	fmt.Println()
}