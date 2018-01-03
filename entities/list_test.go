package entities

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestList_Size(t *testing.T) {
	list := NewList()
	assert.Equal(t, 0, list.Size)
}

func TestList_IsEmpty(t *testing.T) {
	list := NewList()
	assert.True(t, list.IsEmpty())
}

func TestList_InsertAtStart(t *testing.T) {
	list := NewList()
	list.InsertAtStart(6)
	assert.Equal(t, 6, list.Head.Value)
	assert.Nil(t, list.Head.Next)
	assert.False(t, list.IsEmpty())
	assert.Equal(t, 1, list.GetSize())
}

func TestList_InsertAtEnd(t *testing.T) {
	list := NewList()

	list.InsertAtEnd(6)
	assert.Equal(t, 6, list.Head.Value)
	assert.NotNil(t, list.Head)
	assert.False(t, list.IsEmpty())

	list.InsertAtEnd(8)
	assert.Equal(t, 2, list.Size)
	assert.False(t, list.IsEmpty())
	assert.Equal(t, 6, list.Head.Value)
	assert.Equal(t, 8, list.Head.Next.Value)
}

func TestList_InsertSorted(t *testing.T) {
	list := NewList()

	list.InsertSorted(8)
	assert.Equal(t, 1, list.Size)
	assert.False(t, list.IsEmpty())
	assert.Equal(t, 8, list.Head.Value)

	list.InsertSorted(6)
	assert.Equal(t, 2, list.Size)
	assert.Equal(t, 6, list.Head.Value)
	assert.Equal(t, 8, list.Head.Next.Value)

	list.InsertSorted(16)
	assert.False(t, list.IsEmpty())
	assert.Equal(t, 3, list.Size)

	list.InsertSorted(12)
	assert.Equal(t, 4, list.Size)
	assert.Equal(t, 6, list.Head.Value)

	tmp := list.Head

	for tmp.Next != nil {
		assert.True(t, tmp.Value <= tmp.Next.Value)
		fmt.Printf("Comparing %d < %d \n", tmp.Value, tmp.Next.Value)
		tmp = tmp.Next
	}
}

func TestList_IsPresent(t *testing.T) {
	list := NewList()

	list.InsertSorted(4)
	list.InsertSorted(12)
	list.InsertSorted(67)

	assert.True(t, list.IsPresent(4))
	assert.True(t, list.IsPresent(12))
	assert.True(t, list.IsPresent(67))
	assert.False(t, list.IsPresent(54))
}

func TestList_RemoveHeadFromEmptyList(t *testing.T) {
	list := NewList()
	value, removed := list.RemoveHead()
	assert.False(t, removed)
	assert.Equal(t, 0, value)
	assert.Nil(t, list.Head)
}

func TestList_RemoveHeadFromNonEmptyList(t *testing.T) {
	list := NewList()
	list.InsertAtStart(43)
	list.InsertAtStart(12)
	value, removed := list.RemoveHead()

	assert.True(t, removed)
	assert.NotNil(t, list.Head)
	assert.Equal(t, 12, value)
	assert.Equal(t, 1, list.GetSize())
}

func TestList_DeleteNodeFromEmptyList(t *testing.T) {
	list := NewList()
	deleted := list.DeleteNode(23)
	assert.Nil(t, list.Head)
	assert.False(t, deleted)
}

func TestList_DeleteNodeNotInList(t *testing.T) {
	list := NewList()

	list.InsertAtStart(43)
	list.InsertAtStart(12)
	list.InsertAtStart(43)
	list.InsertAtStart(51)

	deleted := list.DeleteNode(2)
	assert.Equal(t, 4, list.GetSize())
	assert.False(t, deleted)
}

func TestList_DeleteNodeHeadOfList(t *testing.T) {
	list := NewList()

	list.InsertAtStart(43)
	list.InsertAtStart(12)
	list.InsertAtStart(43)
	list.InsertAtStart(51)

	deleted := list.DeleteNode(43)
	assert.Equal(t, 3, list.GetSize())
	assert.True(t, deleted)
}

func TestList_DeleteNode(t *testing.T) {
	list := NewList()

	list.InsertAtStart(43)
	list.InsertAtStart(12)
	list.InsertAtStart(18)
	list.InsertAtStart(51)

	deleted := list.DeleteNode(18)
	assert.Equal(t, 3, list.GetSize())
	assert.True(t, deleted)
	assert.False(t, list.IsPresent(18))
}
