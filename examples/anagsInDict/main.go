package main

import (
	"fmt"
	"os"

	"github.com/supermatt/anagrammr/anagrammr"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: command <dictionary path>")
		os.Exit(1)
	}

	dict := args[1]

	words, err := anagrammr.FindAnagsInDict(dict)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	for k, v := range words {
		fmt.Println(k, v)
	}
}
