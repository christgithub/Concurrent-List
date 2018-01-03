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

func (l *List) Display() {
	scroll := l.Head
	for scroll != nil {
		fmt.Println(scroll.Value)
		scroll = scroll.Next

	}
	scroll = nil
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

func (l *List) RemoveHead() (int, bool) {
	l.Mx.Lock()
	defer l.Mx.Unlock()
	if l.IsEmpty() {
		fmt.Println("Error list is empty!")
		return 0, false
	}

	tmpValue := l.Head.Value
	l.Head = l.Head.Next
	l.Size--
	return tmpValue, true
}

func (l *List) DeleteNode(nodeValue int) bool {
	l.Mx.Lock()
	defer l.Mx.Unlock()
	if l.IsEmpty() {
		fmt.Println("Cannot delete from empty list!")
		return false
	}

	if l.Head.Value == nodeValue {
		l.Head = l.Head.Next
		l.Size--
		return true
	}

	scroll := l.Head
	for scroll.Next != nil {
		if scroll.Next.Value == nodeValue {
			scroll.Next = scroll.Next.Next
			l.Size--
			return true
		}
		scroll = scroll.Next
	}
	fmt.Printf("Didn't find node value %d in the list!\n", nodeValue)
	return false
}
