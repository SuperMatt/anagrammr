package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/supermatt/anagrammr/anagrammr"
)

func main() {
	dictFile := flag.String("dict", "/usr/share/dict/british-english", "Path to the dictionary file")
	letters := flag.String("letters", "", "Letters in which to find anagrams")
	capsAllowed := flag.Bool("caps", false, "Allow words with caps in dict")
	minLen := flag.Int("minlen", 4, "Minimum word length")
	debug := flag.Bool("d", false, "debug mode")

	flag.Parse()

	if *debug {
		anagrammr.DebugEnable()
	}

	d, err := anagrammr.LoadDictFromFile(*dictFile, *capsAllowed)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	words, debugInfo := anagrammr.FindAnag(d, *letters, *minLen)

	if *debug {
		fmt.Println(debugInfo)
	}

	for i := len(*letters); i >= *minLen; i-- {
		fmt.Println(i, "letter words:")
		for _, v := range words[i] {
			fmt.Printf("%s ", v)
		}

		fmt.Println()
		fmt.Println()

	}

}
