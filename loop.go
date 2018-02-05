package main

import (
	"strconv"
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
)

func convertToBin(n int) string  {
	result := ""
	for ; n >0; n /=2{
		lsb := n%2
		result = strconv.Itoa(lsb)+result
	}
	return result
}

func printFile(filename string)  {
	file,err :=os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func forever()  {
	for{
		fmt.Println("abc")
	}

}

func main()  {
	fmt.Println("convertToBin results:")
	fmt.Println(
		convertToBin(15),
		convertToBin(2),
		convertToBin(5),
	)
	fmt.Println("abc.txt contents:")
	printFile("abc.txt")


	s := `abc "ad"
	kddd
	123
	p
	22333'`
	printFileContents(strings.NewReader(s))
}



