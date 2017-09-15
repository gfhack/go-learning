package main

import (
  "fmt"
)

func main() {
  smoke_say := "two number %d\n"

  nine := uint32(9)

  pi := float32(3.14)

  fmt.Printf("Value: %.2f\n", pi)
  fmt.Printf(smoke_say, nine)
}
