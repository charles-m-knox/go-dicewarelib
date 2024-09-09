package main

import (
	"embed"
	"log"

	dice "github.com/charles-m-knox/go-dicewarelib"
)

//go:embed words-simple.txt
//go:embed words-complex.txt
var content embed.FS

func main() {
	useExtra := false

	words := dice.Words{}

	// load the simple word dictionary into memory - the memory profile of
	// these words is negligible
	simple, scount := dice.GetWords(content, "words-simple.txt")
	words.Simple = &simple
	words.SimpleCount = scount

	// depending on your application, loading the complex words may take about
	// 30MB of extra RAM
	if useExtra {
		complex, ccount := dice.GetWords(content, "words-complex.txt")
		words.Complex = &complex
		words.ComplexCount = ccount
	} else {
		words.Complex = &map[int]string{} // zero out the ram usage
		words.ComplexCount = 0
	}

	log.Printf("loaded %v simple words and %v complex words", words.SimpleCount, words.ComplexCount)
	log.Println(dice.GeneratePassword(&words, 3, " ", 64, 20, false))
}
