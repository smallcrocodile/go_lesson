package queue

import (
	"fmt"
)

type Queue []interface{}

func (q *Queue) Push(v int) {
	fmt.Println(q)

	*q = append(*q, v)

}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
