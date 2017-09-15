package main

import (
  "fmt"
)

const (
  message = "iota: %d %d\n"
  answer1 = iota*2
  answer2
)

func main() {
  fmt.Printf(message, answer1, answer2)
}
