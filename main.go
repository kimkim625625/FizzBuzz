package main

import (
	fizzbuzz "FizzBuzz/fizz_buzz"
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz.FizzBuzz(i))
	}
}
