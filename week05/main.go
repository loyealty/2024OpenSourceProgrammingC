package main

import (
	"fmt" // #include <stdio.h>
	"reflect"
	"strings"
)

func main() {
	//var i int = 55

	i := 55
	f := 3.99
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f))
	fmt.Println(f, math.call(3.49))
	fmt.Println(strings.Title("kim inha"))
	fmt.Println("i는 %d\n", i)
	fmt.Print("i는", i, "\n")
	fmt.Println("i는", i)
}
