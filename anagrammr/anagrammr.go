package anagrammr

import (
	"bytes"
	"io/ioutil"
	"sort"
	"strings"
)

var debug = false
var debugBuffer = bytes.NewBuffer([]byte{})

func debugPrint(s string) {
	if debug {
		debugBuffer.WriteString(s + "\n")
	}
}

//FindAnag performs an anagram lookup of the letters in a dictionary with min length of words
func FindAnag(d *[][]byte, l string, minLen int) (map[int][]string, string) {

	words := make(map[int][]string, len(l))

	maxLen := len(l)
	base := make([]int, 26)
	b := bytes.ToLower([]byte(l))
	for _, v := range b {
		base[v-97]++
	}

	for _, w := range *d {
		debugPrint("Word: " + string(w))
		if len(w) > maxLen {
			debugPrint("Word too long")
			continue
		}
		if len(w) < minLen {
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
			words[len(w)] = append(words[len(w)], string(w))
		}
	}

	return words, debugBuffer.String()
}

//LoadDictFromFile reads a dictionary from a file
func LoadDictFromFile(s string, capsAllowed bool) (d *[][]byte, _ error) {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		return nil, err
	}

	ds := bytes.Split(b, []byte{10})

	return buildDict(&ds, capsAllowed), nil
}

//LoadDictFromStrings takes a slice of strings and turns them into a dictionary
func LoadDictFromStrings(s *[]string, capsAllowed bool) *[][]byte {
	b := make([][]byte, 0)

	for _, v := range *s {
		b = append(b, []byte(v))
	}

	return buildDict(&b, capsAllowed)
}

func buildDict(b *[][]byte, capsAllowed bool) (d *[][]byte) {

	var td [][]byte
	for k := range *b {
		fail := false
		if len((*b)[k]) == 0 {
			break
		}
		if (*b)[k][0]-97 > 25 && !capsAllowed {
			fail = true
		}

		if bytes.Contains((*b)[k], []byte{39}) {
			fail = true
		}

		if !fail {
			td = append(td, bytes.ToLower((*b)[k]))
		}
	}

	d = &td
	return d
}

//DebugEnable enabled debugging
func DebugEnable() {
	debug = true
}

//DebugDisable disable debugging
func DebugDisable() {
	debug = false
}

//FindAnagsInDict finds all anagrams in a dictionary.
func FindAnagsInDict(filename string) (w map[string][]string, _ error) {
	dictBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return w, err
	}

	words := make(map[string][]string)
	anags := make(map[string][]string)

	word := ""
	for itr := 0; itr < len(dictBytes); itr++ {
		lb := dictBytes[itr]
		if lb != 10 {
			word += string([]byte{lb})
		} else {
			spl := strings.Split(word, "")
			sort.Strings(spl)
			sorted := strings.Join(spl, "")
			_, ok := words[sorted]
			if ok {
				words[sorted] = append(words[sorted], word)
			} else {
				words[sorted] = []string{word}
			}
			word = ""
		}
	}

	for k, v := range words {
		if len(v) > 1 {
			anags[k] = v
		}
	}

	return anags, nil

}
