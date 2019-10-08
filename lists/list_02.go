package main

func (l *lNode) duplicateRemove() {
	if l.next == nil {
		return
	}

}

func main() {
	l := newList(10)
	l.duplicateRemove()
	l.print()
}