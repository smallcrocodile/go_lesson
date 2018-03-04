package queue

import (
	"fmt"
)

//A FIFO queue.
type Queue []interface{}

func (q *Queue) Push(v int) {
	fmt.Println(q)

	*q = append(*q, v)

}

// Pops element from head.
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

// Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
