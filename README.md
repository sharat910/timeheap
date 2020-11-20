# timeheap

A golang package to store and retrieve sorted timestamped objects.

Example usage: `cmd/example`

It exposes the following public APIs
```go
// New creates a timeheap with the initial capacity
func New(capacity int) *TimeHeap

// Peek returns a copy of oldest element (in Time)
// If the heap is empty, returns ErrHeapEmpty
// Time complexity: O(1)
func (h *TimeHeap) Peek() (Timestamped, error) 

// Pop extracts and returns oldest element (in Time)
// If the heap is empty, returns ErrHeapEmpty
// Time complexity: O(log(n))
func (h *TimeHeap) Pop() (Timestamped, error)

// Push inserts a new element into the heap
// Time complexity: O(log(n)) in the worst case,
//                  O(1) (amortized) if objects are generally moving forward in Time
func (h *TimeHeap) Push(item Timestamped)
```

