package main

import (
	"fmt"
)

func swap(first int, second int) {
	temp := 0
	temp = first
	first = second
	second = temp
	fmt.println(first, second)

}
func main() {
	var a, b int = 10, 20
	fmt.Println(a, b)
	swap(a, b)
	println(a, b)
}
