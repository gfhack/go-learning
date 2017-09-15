package main

import (
  "fmt"
)

func main() {
  message := "The answer to life is %d!\n"

  answer := 42

  fmt.Printf(message, answer)
}
