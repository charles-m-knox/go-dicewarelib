# go-dicewarelib

Contains a few basic functions that allow for diceware password generation, with sensible defaults that work with most websites.

## Setup

Install:

```bash
go get git.cmcode.dev/cmcode/go-dicewarelib
```

## Usage

```go
package main

import (
    "log"
    "embed"

    dice "git.cmcode.dev/cmcode/go-dicewarelib"
)

//go:embed words-simple.txt
//go:embed words-complex.txt
var content embed.FS

func main() {
    // load the simple word dictionary into memory
    words := dice.Words{}
	*words.Simple, words.SimpleCount = dice.GetWords(content, "words-simple.txt")
	log.Println(GeneratePassword(&words, 3, " ", 64, 20, false))
}
```
