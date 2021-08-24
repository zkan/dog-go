package main

import (
	"fmt"

	"github.com/zkan/dog-go/fizzbuzz"
)

func main() {
	fmt.Println("vim-go")

	var x string = "Hello"
	//x = "Hello"
	fmt.Println(x)

	y := 10
	fmt.Println(y)

	var number int = 15
	result := fizzbuzz.FizzBuzz(number)
	fmt.Println(result)
}
