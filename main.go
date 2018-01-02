package main

import (
	"sync"

	"github.com/MyProjects/Concurrent-List/entities"
)

func main() {
	wg := &sync.WaitGroup{}
	l := entities.NewList()

	wg.Add(1)
	go func() {
		defer wg.Done()
		l.InsertSorted(47)
		l.InsertSorted(4)
		l.InsertSorted(3)
		l.InsertSorted(14)
		l.InsertSorted(28)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		l.InsertSorted(54)
		l.InsertSorted(49)
		l.InsertSorted(32)
		l.InsertSorted(144)
		l.InsertSorted(280)
	}()
	wg.Wait()
	l.Display()

}
