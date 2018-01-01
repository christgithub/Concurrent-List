package entities

import (
	"fmt"
	"sync"
)

type List struct {
	Head *Node
	Size int
	Mx   *sync.Mutex
}

type Node struct {
	Value int
	Next  *Node
}

func NewList() *List {
	return &List{
		Head: nil,
		Size: 0,
		Mx:   &sync.Mutex{},
	}
}

func (l *List) GetSize() int {
	return l.Size
}

func (l *List) IsEmpty() bool {
	return l.Size == 0
}

func (l *List) InsertAtStart(value int) {
	l.Mx.Lock()
	defer l.Mx.Unlock()
	l.Head = &Node{value, l.Head}
	l.Size++
}

func (l *List) InsertAtEnd(value int) {
	newNode := &Node{value, nil}
	l.Mx.Lock()
	defer l.Mx.Unlock()

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
	l.Mx.Lock()
	defer l.Mx.Unlock()

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
	l.Mx.Lock()
	defer l.Mx.Unlock()
	tmp := l.Head

	for tmp != nil {
		if value == tmp.Value {
			return true
		}
		tmp = tmp.Next
	}
	return false
}
