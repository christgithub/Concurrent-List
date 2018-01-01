package entities

import "fmt"

type List struct {
	Head *Node
	Size int
}

type Node struct {
	Value int
	Next  *Node
}

func NewList() *List {
	return &List{
		Head: nil,
		Size: 0,
	}
}

func (l *List) GetSize() int {
	return l.Size
}

func (l *List) IsEmpty() bool {
	return l.Size == 0
}

func (l *List) InsertAtStart(value int) {
	l.Head = &Node{value, l.Head}
	l.Size++
}

func (l *List) InsertAtEnd(value int) {
	newNode := &Node{value, nil}

	if l.Head == nil {
		l.Head = newNode
		l.Size++
		return
	}

	scroll := l.Head

	for scroll.Next != nil {
		scroll = scroll.Next
	}

	scroll.Next = newNode
	scroll = nil
	l.Size++
}

func (l *List) InsertSorted(value int) {
	newNode := &Node{value, nil}
	scroll := l.Head

	if scroll == nil || scroll.Value > value {
		newNode.Next = l.Head
		l.Head = newNode
		l.Size++
		scroll = nil
		return

	}

	for scroll.Next != nil && scroll.Next.Value < value {
		scroll = scroll.Next
	}

	newNode.Next = scroll.Next
	scroll.Next = newNode
	l.Size++
	scroll = nil
}

func (l *List) InsertAtPosition(value int, position int) {

}

func (l *List) displayList() {
	scroll := l.Head
	for scroll != nil {
		fmt.Println(scroll.Value)
		scroll = scroll.Next

	}
	scroll = nil
}

func (l *List) isSortedAscending() bool {
	return false
}

func (l *List) isSortedDescending() bool {
	return false
}

func (l *List) IsPresent(value int) bool {
	tmp := l.Head

	for tmp != nil {
		if value == tmp.Value {
			return true
		}
		tmp = tmp.Next
	}
	return false
}
