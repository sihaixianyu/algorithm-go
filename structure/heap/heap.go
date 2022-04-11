package heap

import "golang.org/x/exp/constraints"

type Heap[T any] struct {
	len  uint
	vals []T
	less func(T, T) bool
}

func New[T any](comparator func(T, T) bool) Heap[T] {
	return Heap[T]{
		len: 0,
		// Leave T[0] unused for calculation simplicity
		vals: make([]T, 1),
		less: comparator,
	}
}

func NewMin[T constraints.Ordered]() Heap[T] {
	return Heap[T]{
		len:  0,
		vals: make([]T, 1),
		less: func(v1, v2 T) bool {
			return v1 < v2
		},
	}
}

func NewMax[T constraints.Ordered]() Heap[T] {
	return Heap[T]{
		len:  0,
		vals: make([]T, 1),
		less: func(v1, v2 T) bool {
			return v1 > v2
		},
	}
}

func (h *Heap[T]) Len() uint {
	return h.len
}

func (h *Heap[T]) Empty() bool {
	return h.len == 0
}

func (h *Heap[T]) Push(val T) {
	h.len += 1
	h.vals = append(h.vals, val)

	// Heapify from bottom to top
	h.heapifyUp(h.len)
}

func (h *Heap[T]) Pop() T {
	if h.len == 0 {
		panic("the heap is empty")
	}

	h.vals[1], h.vals[h.len] = h.vals[h.len], h.vals[1]
	res := h.vals[h.len]
	h.vals = h.vals[:h.len]
	h.len -= 1

	h.heapifyDown(1)

	return res
}

func (h *Heap[T]) heapifyUp(idx uint) {
	for idx > 1 && h.less(h.vals[idx], h.vals[idx/2]) {
		h.vals[idx/2], h.vals[idx] = h.vals[idx], h.vals[idx/2]
		idx = idx / 2
	}
}

func (h *Heap[T]) heapifyDown(idx uint) {
	for 2*idx <= h.len {
		childIdx := 2 * idx
		if childIdx < h.len && h.less(h.vals[childIdx+1], h.vals[childIdx]) {
			childIdx += 1
		}

		if h.less(h.vals[idx], h.vals[childIdx]) {
			break
		}

		h.vals[idx], h.vals[childIdx] = h.vals[childIdx], h.vals[idx]
		idx = childIdx
	}
}
