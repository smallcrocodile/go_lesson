package main

import "fmt"

func main()  {
	cache := make(map[string]string)
	cache["name"] = "jack"
	fmt.Println(cache)
}
