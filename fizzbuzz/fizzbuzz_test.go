package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	number := 3
	result := FizzBuzz(number)
	if result != "Fizz" {
		t.Errorf("Get %d; want Fizz", number)
	}
}
