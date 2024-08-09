package main

import (
	"embed"
	"log"
	// dice "git.cmcode.dev/cmcode/go-dicewarelib"
)

//go:embed words-simple.txt
//go:embed words-complex.txt
var content embed.FS

func main() {
	// load the simple word dictionary into memory
	words := Words{}
	*words.Simple, words.SimpleCount = GetWords(content, "words-simple.txt")
	log.Println(GeneratePassword(&words, 3, " ", 64, 20, false))
}
