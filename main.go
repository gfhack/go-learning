package main

import (
  "fmt"
)

const (
  answer = 42
  message = "The answer to life is %d!\n"
)

func main() {
  fmt.Printf(message, answer)
}
