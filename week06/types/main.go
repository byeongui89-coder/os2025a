package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	//var name strings
	//var name = "Kim Inha"
	//name="Kim Inha"
	//var name = "kim Inha"
	var name float64
	name = 2.71
	fmt.Println(name, reflect.TypeOf(name))
	fmt.Println(name)
	fmt.Println(math.Ceil(2.71))
	fmt.Println(strings.Title("head first go"))
}
