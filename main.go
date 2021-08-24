package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/zkan/dog-go/fizzbuzz"
)

type person struct {
	name  string
	money float64
}

func (p person) say() string {
	return "Hey " + p.name
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

	fmt.Println(kan.say())
	fmt.Println(salee.say())

	const url string = "https://dog.ceo/api/breeds/image/random"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("GET Error")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Error")
	}
	fmt.Printf("%s", body)
}
