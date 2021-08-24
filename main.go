package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zkan/dog-go/fizzbuzz"
)

type person struct {
	name  string
	money float64
}

func (p person) say() string {
	return "Hey " + p.name
}

//type some struct {
//    name string
//}
//
//func (s some) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "Hello Dog Go!")
//}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Dog Go from indexHandler!")
}

type Breed struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func breedHandler(w http.ResponseWriter, r *http.Request) {
	const url string = "https://dog.ceo/api/breeds/image/random"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("GET Error")
	}
	defer resp.Body.Close()

	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//    fmt.Println("Read Error")
	//}
	//fmt.Printf("%s", body)

	var b Breed
	if err := json.NewDecoder(resp.Body).Decode(&b); err != nil {
		log.Fatal("Failed!")
	}
	fmt.Println(b.Message)
	fmt.Println(b.Status)

	json.NewEncoder(w).Encode(b)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
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

	//var s some
	//http.Handle("/", s)

	//http.HandleFunc("/", indexHandler)

	//http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
	//    fmt.Fprintf(w, "Hello Dog Go from indexHandler!")
	//})

	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/breeds", breedHandler)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
