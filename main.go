package main

import (
  "fmt"
)

var (
  answer = 41
  message = "The answer to life is %d!\n"
)

func main() {
  answer += 1
  fmt.Printf(message, answer)
}
