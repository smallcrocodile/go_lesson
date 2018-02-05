package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a,b int,op string)  (int,error){
	switch op {
	case "+":
		return a + b,nil
	case "-":
		return a - b,nil
	case "*":
		return a * b,nil
	case "/":
		q,_ := div(a,b)
		return q,nil
	default:
		return 0,fmt.Errorf("unsupported operation:" +op)
	}
}

func apply(op func(int,int)int,a,b int)int  {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args (%d %d)",opName,a,b)
	return op(a,b)
}

//除法

func div(a,b int)(int ,int)  {
	return a/b,a%b
}
func pow(a,b int)int  {
	return int(math.Pow(float64(a),float64(b)))
}

func main()  {
	fmt.Println(eval(3,4,"/-"))
	q,r := div(13,5)
	fmt.Println(q,r)

	fmt.Println(apply(pow,3,4))

}
