package main

import (
	"fmt"
	"imooc/queue"
)

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q)
	fmt.Println(q.IsEmpty())

	fmt.Println(q)

}
