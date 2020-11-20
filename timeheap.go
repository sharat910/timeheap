package timeheap

import (
	"errors"
	"time"
)

var (
	ErrHeapEmpty = errors.New("heap empty")
)

type Timestamped interface {
	T() time.Time
}

type TimeHeap struct {
	heapArray []Timestamped
	lessThan  func(t1, t2 Timestamped) bool
}

// New creates a timeheap with the initial capacity
func New(capacity int) *TimeHeap {
	return &TimeHeap{
		heapArray: make([]Timestamped, 0, capacity),
		lessThan: func(t1, t2 Timestamped) bool {
			return t1.T().Before(t2.T())
		},
	}
}

// Peek returns a copy of oldest element (in Time)
// If the heap is empty, returns ErrHeapEmpty
// Time complexity: O(1)
func (h *TimeHeap) Peek() (Timestamped, error) {
	if len(h.heapArray) == 0 {
		return nil, ErrHeapEmpty
	}
	return h.heapArray[0], nil
}

// Pop extracts and returns oldest element (in Time)
// If the heap is empty, returns ErrHeapEmpty
// Time complexity: O(log(n))
func (h *TimeHeap) Pop() (Timestamped, error) {
	if len(h.heapArray) == 0 {
		return nil, ErrHeapEmpty
	}
	root := h.heapArray[0]
	h.heapArray[0] = h.heapArray[len(h.heapArray)-1]
	h.heapArray = h.heapArray[:len(h.heapArray)-1]
	h.downHeapify()
	return root, nil
}

// Push inserts a new element into the heap
// Time complexity: O(log(n)) in the worst case,
//                  O(1) (amortized) if objects are generally moving forward in Time
func (h *TimeHeap) Push(item Timestamped) {
	h.heapArray = append(h.heapArray, item)
	h.upHeapify()
}

// internal Functions

func (h *TimeHeap) isLeaf(index int) bool {
	if index >= (len(h.heapArray)/2) && index <= len(h.heapArray) {
		return true
	}
	return false
}

func (h *TimeHeap) parent(index int) int {
	return (index - 1) / 2
}

func (h *TimeHeap) leftChild(index int) int {
	return 2*index + 1
}

func (h *TimeHeap) hasLeftChild(index int) bool {
	return h.leftChild(index) < len(h.heapArray)
}

func (h *TimeHeap) rightChild(index int) int {
	return 2*index + 2
}

func (h *TimeHeap) hasRightChild(index int) bool {
	return h.rightChild(index) < len(h.heapArray)
}

func (h *TimeHeap) swap(first, second int) {
	h.heapArray[first], h.heapArray[second] = h.heapArray[second], h.heapArray[first]
}

func (h *TimeHeap) upHeapify() {
	index := len(h.heapArray) - 1
	for h.lessThan(h.heapArray[index], h.heapArray[h.parent(index)]) {
		h.swap(index, h.parent(index))
		index = h.parent(index)
	}
}

func (h *TimeHeap) downHeapify() {
	index := 0
	for h.hasLeftChild(index) {

		smallChildIdx := h.leftChild(index)
		if h.hasRightChild(index) && h.lessThan(h.heapArray[h.rightChild(index)], h.heapArray[smallChildIdx]) {
			smallChildIdx = h.rightChild(index)
		}

		if h.lessThan(h.heapArray[index], h.heapArray[smallChildIdx]) {
			break
		}

		h.swap(index, smallChildIdx)
		index = smallChildIdx
	}
}
