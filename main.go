package main

import (
	"fmt"

	"github.com/zkan/dog-go/fizzbuzz"
)

type person struct {
	name  string
	money float64
}

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

	var kan person
	kan.name = "Kan Ouivirach"
	kan.money = 11.25
	fmt.Println(kan)

	salee := person{
		name:  "Salee",
		money: 130.5,
	}
	fmt.Println(salee)
}
