package main

//#include
import "C"

func main() {
	//
}

//export MyStrDup
func MyStrDup(s *C.char) *C.char {
	return C.strdup(s)
}
