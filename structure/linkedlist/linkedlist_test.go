package linkedlist

import (
	"fmt"
	"testing"
)

var vals []int

func init() {
	vals = []int{1, 2, 3}
}

func TestFromSlice(t *testing.T) {
	list := FromSlice(vals)
	fmt.Println(list.Head)
}

func TestLinkedList_Traverse(t *testing.T) {
	list := FromSlice(vals)
	fmt.Println(list.Traverse())
}
