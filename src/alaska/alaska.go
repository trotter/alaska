package main

import (
	"fmt"
)

func main() {
	dict := StringDictionary{[]string{
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

	intDict := IntDictionary{[]int{1,2,4,5,12}}
	fmt.Printf("Index of 4: %v\n", intDict.IndexOf(12))
	fmt.Printf("ValueAt 2: %v\n", dict.ValueAt(2))

	chunk := ColumnChunk{intDict, []int{3,2,0,4,0,0,2,1,3,2}}
	fmt.Printf("String ValueAt 2: %v\n", dict.ValueAt(chunk.GetGlobalId(3)))

	column := StringColumn{dict, []ColumnChunk{chunk}}
	fmt.Printf("Matching chunk locations: %s\n", column.Where("amazon"))
}
