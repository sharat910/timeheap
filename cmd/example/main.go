package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/sharat910/timeheap"
)

type Object struct {
	t    time.Time
	data int
}

func (o Object) T() time.Time {
	return o.t
}

func main() {
	capacity := 10
	h := timeheap.New(capacity)
	now := time.Now()
	fmt.Println("now", now)
	for i := 0; i < capacity; i++ {
		h.Push(Object{t: now.Add(time.Duration(rand.Intn(60)) * time.Second)})
	}
	for i := 0; i < capacity; i++ {
		obj, err := h.Pop()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(obj.T())
	}
	_, err := h.Pop()
	if err != nil {
		fmt.Println(err)
	}
}
