package main

import (
	"fmt"
)

func main() {
  dict := Dictionary{[]string{"a", "b", "c", "d"}}
  fmt.Printf("IndexOf b: %v\n", dict.IndexOf("b"))
  fmt.Printf("IndexOf c: %v\n", dict.IndexOf("c"))

  fmt.Printf("ValueAt 2: %v\n", dict.ValueAt(2))
}
