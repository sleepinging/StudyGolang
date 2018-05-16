package main

//
import "C"
import "fmt"

//export Sum
func Sum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("123")
}
