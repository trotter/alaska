package main

import (
	"fmt"
	"alaska"
)

func main() {
	searchTerm := "ebay"

	searchTermDict := alaska.StringDictionary{[]string{
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
		"yellow pages",
	}}
	fmt.Printf("IndexOf cheap flights: %v\n", searchTermDict.IndexOf("cheap flights"))
	fmt.Printf("IndexOf la redoute: %v\n", searchTermDict.IndexOf("la redoute"))

	intDict := alaska.IntDictionary{[]int{1,2,4,5,12}}
	fmt.Printf("Index of 4: %v\n", intDict.IndexOf(12))

	chunk := alaska.ColumnChunk{intDict, []int{3,2,0,4,0,0,2,1,3,2}}

	column := alaska.StringColumn{searchTermDict, []alaska.ColumnChunk{chunk}}
	fmt.Printf("Matching chunk locations: %s\n", column.Where(searchTerm))

	countryDict := alaska.StringDictionary{[]string{
		"Argentina",
		"Brazil",
		"Canada",
		"Germany",
		"USA",
		"Vietnam",
	}}
	countryIntDict := alaska.IntDictionary{[]int{1,4,5}}
	countryChunk := alaska.ColumnChunk{countryIntDict, []int{0,0,0,1,1,1,2,2,2,2}}
	countryColumn := alaska.StringColumn{countryDict, []alaska.ColumnChunk{countryChunk}}
	fmt.Printf("Matching countries: %s\n", countryColumn.Select(column.Where(searchTerm)))
}
