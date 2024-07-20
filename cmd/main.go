package main

import (
	"fmt"
	"github.com/j4ck4L0up3/pokemonscraper/pokescraper"
	// "os"
)

func main() {
	url := "https://serebii.net/pokedex-sv/"
	numRegions := 9
	pokemonMatrix := pokescraper.ProcessPokemonMatrix(url, numRegions)
	fmt.Printf("Pokemon Matrix:\n%v\n", pokemonMatrix)

	// TODO: retieve all the page urls from {elem: "option", attrKey: "value"}
	// do this in batches of 151,100, 135, 107, 156, 72, 88, 96, 120
	// for Kanto, Johto, Hoenn, Sinnoh, Unova, Kalos, Alola, Galar/Hisui, Paldea, respectively
	// store in map

	// filename := "playgroud.txt"
	// file, osErr := os.Create(filename)
	// if osErr != nil {
	// 	fmt.Printf("error creating or opening file {%v}: %v", filename, osErr)
	// 	return
	// }
	//
	// defer file.Close()

	// _, writeErr := file.WriteString(htmlRawStr)
	// if writeErr != nil {
	// 	fmt.Printf("error writing to file {%v}: %v", filename, writeErr)
	// }
	//
	// for _, attrVal := range pageUrls {
	// 	_, osErr := file.WriteString(attrVal + "\n")
	// 	if osErr != nil {
	// 		fmt.Printf("error writing to file {%v}: %v", filename, osErr)
	// 	}
	// }

	// TODO: make a processing file to process and parse the returned strings

	// for _, idName := range pokeIdNames {
	// 	_, osErr := file.WriteString(idName + "\n")
	// 	if osErr != nil {
	// 		fmt.Printf("error writing to file {%v}: %v", filename, osErr)
	// 	}
	// }

	// pokeByRegion := make(map[string][]string)
	// pokeGroups := make([][]string, 1025)

	// batches := []int{151,100, 135, 107, 156, 72, 88, 96, 120}
	// for i := 0; i <= len(batches); i++ {
	// 	for j := 0; j <= batches[i]; j++ {
	// 		pokeGroups[i] = append(pokeGroups[i], )
	// 	}
	// }

	// TODO: to get id, name separated in a tuple like item

	// TODO: to get types, search {elem: <img>, attrKey: alt}, then filter pattern \b-type (or something)
	// use GetDOMAttrVals

}
