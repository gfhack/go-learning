package main

import (
  "fmt"
)

func main() {
  atoz := "the quick brown fox jumps over the lazy dog\n"
  literal := `the quick brown fox jumps over the lazy dog\n`

  fmt.Printf("%s\n", atoz[0:9])
  fmt.Printf("%s\n", atoz[:9])
  fmt.Printf("%s\n", atoz[15:19])
  fmt.Printf("%s\n", atoz[15:])

  for index, character := range atoz {
    fmt.Printf("index, character: %d, %c\n", index, character)
  }

  for _, character := range atoz {
    fmt.Printf("character: %c\n", character)
  }

  for index := range atoz {
    fmt.Printf("index: %d\n", index)
  }

  fmt.Printf("length: %d\n", len(atoz))
  fmt.Printf("literal: %s\n", literal)
}
