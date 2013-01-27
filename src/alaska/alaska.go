package main

import (
	"fmt"
)

func main() {
  dict := Dictionary{[]string{
    "ab in den Urlaub",
    "amazon",
    "chaussures",
    "cheap flights",
    "cheap tickets",
    "ebay",
    "faschingskostume",
    "immobilienscout",
    "karnevalskostume",
    "la redoute",
    "pages jaunes",
    "voyages sncf",
    "yellow pages"}}

  fmt.Printf("IndexOf cheap flights: %v\n", dict.IndexOf("cheap flights"))
  fmt.Printf("IndexOf la redoute: %v\n", dict.IndexOf("la redoute"))

  fmt.Printf("ValueAt 2: %v\n", dict.ValueAt(2))
}
