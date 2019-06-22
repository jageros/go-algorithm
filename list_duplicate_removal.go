package main

func (l *lNode) duplicateRemove() {
	if l.next == nil {
		return
	}

}

func main() {
	l := newList()
	l.duplicateRemove()
	l.print()
}