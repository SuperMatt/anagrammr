package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

var debug = false

func debugPrint(s string) {
	if debug {
		fmt.Println(s)
	}
}

func findAnag(d *[][]byte, l *string, minLen *int) (s []string) {
	maxLen := len(*l)
	base := make([]int, 26)
	b := bytes.ToLower([]byte(*l))
	for _, v := range b {
		base[v-97]++
	}

	for _, w := range *d {
		debugPrint("Word: " + string(w))
		if len(w) > maxLen {
			debugPrint("Word too long")
			continue
		}
		if len(w) < *minLen {
			debugPrint("Word too short")
			continue
		}
		fail := false
		count := make([]int, 26)
		for _, l := range w {
			debugPrint("Letter: " + string(l))
			bnum := l - 97
			if bnum > 26 {
				debugPrint("Probably not a letter")
				fail = true
				break
			}
			if base[bnum] == 0 {
				debugPrint("Letter does not exist in range")
				fail = true
				break
			}

			if count[bnum] >= base[bnum] {
				debugPrint("Too many of letter")
				fail = true
				break
			} else {
				count[bnum]++
			}
		}
		if !fail {
			s = append(s, string(w))
		}
	}

	return s
}

func readDict(s *string, capsAllowed *bool) (d *[][]byte, _ error) {
	b, err := ioutil.ReadFile(*s)
	if err != nil {
		return nil, err
	}

	ds := bytes.Split(b, []byte{10})
	var td [][]byte
	for k := range ds {
		fail := false
		if len(ds[k]) == 0 {
			break
		}
		if ds[k][0]-97 > 25 && !*capsAllowed {
			fail = true
		}

		if bytes.Contains(ds[k], []byte{39}) {
			fail = true
		}

		if !fail {
			td = append(td, bytes.ToLower(ds[k]))
		}
	}

	d = &td
	return d, nil
}

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

	ls := make([][]string, len(*letters))

	for _, word := range words {
		i := len(word) - 1
		ls[i] = append(ls[i], word)

	}

	for j := len(*letters) - 1; j >= *minLen-1; j-- {
		fmt.Println(j+1, "letter words")
		sort.Strings(ls[j])
		fmt.Println(ls[j])
	}
}
