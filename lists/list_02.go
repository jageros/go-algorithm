package main

func duplicateRemove(l *lNode) {
	if l == nil {
		return
	}
	q := l
	p := l.next
	for p != nil {
		if p.data == l.data {
			q.next = p.next
		}else {
			q = p
		}
		p = q.next
	}
	duplicateRemove(l.next)
}

func main() {
	l := newList(5, 30)
	l.print()
	duplicateRemove(l.p)
	l.print()
}
