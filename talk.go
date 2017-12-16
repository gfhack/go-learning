package main

import (
	"fmt"
)

func add(x int, y int) int {
  return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
  a, b := swap("hello", "world")
  _, c := swap("hello", "world")
  fmt.Println(a, b, c)

  var number int
  number = 42

	if sum := add(number, 13); sum > 50 {
		fmt.Printf("Resultado da soma: %d\n", sum)
	}

	acc := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	carry := 1
	for carry < 1000 {
		carry += carry
	}
}
