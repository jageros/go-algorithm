package main

func (l *lNode) duplicateRemove() {
	if l.next == nil {
		return
	}
	l.next.forEachData(func(data int) bool {
		if l.data == data {
			l.data = 0
		}
		return true
	})

	l.next.duplicateRemove()
}

func main() {
	l := newList(10, 20)
	l.print()
	l.next.duplicateRemove()
	l.print()
}