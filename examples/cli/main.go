package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	dictFile := flag.String("dict", "/usr/share/dict/british-english", "Path to the dictionary file")
	letters := flag.String("letters", "", "Letters in which to find anagrams")
	capsAllowed := flag.Bool("caps", false, "Allow words with caps in dict")
	minLen := flag.Int("minlen", 4, "Minimum word length")
	db := flag.Bool("d", false, "debug mode")

	flag.Parse()

	if *db {
		debug = true
	}

	fmt.Println(*dictFile, *letters)

	d, err := readDict(dictFile, capsAllowed)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	words := findAnag(d, letters, minLen)

}
