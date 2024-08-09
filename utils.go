package main

import (
	"bufio"
	"crypto/rand"
	"embed"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	// Prevents the CPU from endlessly attempting to fulfill unrealistic
	// password generation parameters. For example, specifying a minimum
	// generated output length of 128 characters, but setting the number of
	// words to 2 will likely cause password generation to fail due to
	// impossible constraints.
	maxAttempts   = 20000
	maxWordLength = 16 // Prevents obnoxiously long words from being used
	minWordLength = 4  // Prevents words like "a", "um", "etc", etc
)

var (
	symbols     = []string{"!", "@", "#", "$", "%", "*", "/", "?", ".", ","}
	symbolCount = len(symbols) // Calculated once on startup instead of repeatedly
)

// Words stores all of the simple and/or complex words in memory. The number of
// words is calculated once and stored as well for later (frequent) use.
type Words struct {
	Simple       *map[int]string
	SimpleCount  int
	Complex      *map[int]string
	ComplexCount int
}

// GetRandomInt generates a random number from 0 to m.
func GetRandomInt(m int) int {
	k, _ := rand.Int(rand.Reader, big.NewInt(int64(m)))

	return int(k.Int64())
}

// GetWords loads all words from the given path into memory. Note that the word
// list must be embedded, using embed.FS.
func GetWords(content embed.FS, path string) (map[int]string, int) {
	readFile, err := content.Open(path)
	if err != nil {
		log.Printf("GetWords error: %v", err.Error())
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result := make(map[int]string)

	i := 0

	for fileScanner.Scan() {
		w := fileScanner.Text()
		result[i] = w
		i++
	}

	readFile.Close()

	return result, i
}

// getRandomWord picks a random word from the provided word list.
// Words that are either too short or too long are ignored. For better
// performance, consider trimming the input word list to match these
// restrictions first.
func getRandomWord(m map[int]string) string {
	for {
		k := GetRandomInt(len(m))

		v, ok := m[k]
		if !ok {
			return ""
		}

		if len(v) < minWordLength || len(v) > maxWordLength {
			continue
		}

		return v
	}
}

// getRandomSymbol returns a random symbol from the list of symbols defined
// in this package.
func getRandomSymbol() string {
	return symbols[GetRandomInt(symbolCount-1)]
}

// GeneratePassword generates a password according to a few rules:
//
// w=dictionary of random words to choose from
//
// n=number of words
//
// s=separator character
//
// maxLen=maximum allowable length of the resulting password
//
// minLen=minimum allowable length of the resulting password.
func GeneratePassword(words *Words, n int, s string, maxLen int, minLen int, extendedWords bool) string {
	startTime := time.Now()

	w := words.Simple
	if extendedWords {
		w = words.Complex
	}

	sb := new(strings.Builder)

	getWords := func() {
		// start by generating the number of requested words
		for pi := 0; pi < n; pi++ {
			sb.WriteString(getRandomWord(*w))
			// don't put the separator character after
			// the last word
			if pi != n-1 {
				sb.WriteString(s)
			}
		}
	}

	// brute-force generate words and ensure that they're within the requested
	// maxLen and minLen
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		// abort if the operation has taken longer than 1 seconds
		if startTime.Add(1 * time.Second).Before(time.Now()) {
			log.Println("debug: operation took too long, canceling")

			return ""
		}

		getWords()

		pwlen := sb.Len() + 2

		if pwlen <= maxLen && pwlen >= minLen {
			sb.WriteString(strconv.Itoa(GetRandomInt(9)))
			sb.WriteString(getRandomSymbol())

			break
		}

		sb.Reset()

		if attempts >= maxAttempts {
			log.Println("debug: exceeded maximum attempts to generate password")

			return ""
		}
	}

	result := sb.String()

	// capitalize the first letter
	rstr := []rune(result)
	rstr[0] = unicode.ToUpper(rstr[0])

	return string(rstr)
}
