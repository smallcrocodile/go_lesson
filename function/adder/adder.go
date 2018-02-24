package main

import "fmt"

func adder() func(int) int {
	sum := 0 //自由变量
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}

	b := adder2(0)
	s := 0
	for i := 0; i < 10; i++ {
		s, b = b(i)
		fmt.Println(i, s)
	}

}
