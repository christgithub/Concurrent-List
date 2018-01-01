package main

import (
	"fmt"

	"github.com/MyProjects/LinkedList/entities"
)

func main() {
	l := entities.NewList()

	displayList(l)
}

func displayList(l *entities.List) {
	scroll := l.Head
	for scroll != nil {
		fmt.Println(scroll.Value)
		scroll = scroll.Next

	}
	scroll = nil
}
