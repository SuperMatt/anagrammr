package anagrammr

import (
	"bytes"
	"io/ioutil"
	"sort"
	"strings"
)

var separatorBytes = map[byte]bool{9: true, //tab
	10: true, //line feed
	13: true, //carriage return
	32: true} //space

var ignoreBytes = map[byte]bool{33: true, //!
	34: true, //"
	39: true, //'
	40: true, //(
	41: true, //)
	44: true, //,
	46: true, //.
	47: true, //\/
	58: true, //:
	59: true, //;
	63: true, //?
	91: true, //[
	92: true, //\
	93: true} //]

var debug = false
var debugBuffer = bytes.NewBuffer([]byte{})

func debugPrint(s string) {
	if debug {
		debugBuffer.WriteString(s + "\n")
	}
}

//FindAnag performs an anagram lookup of the letters in a dictionary with min length of words
func FindAnag(d *[][]byte, letters string, minLen int) (map[int][]string, string) {

	words := make(map[int][]string, len(letters))

	maxLen := len(letters)
	base := make([]int, 256)
	for _, v := range []byte(letters) {
		base[v]++
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
		count := make([]int, 256)
		fail := false
		for _, l := range w {
			debugPrint("Letter: " + string(l))
			if count[l] >= base[l] {
				debugPrint("Too many of letter")
				fail = true
				break
			} else {
				count[l]++
			}
		}
		if !fail {
			words[len(w)] = append(words[len(w)], string(w))
		}
	}

	return words, debugBuffer.String()
}

//LoadDictFromFile reads a dictionary from a file
func LoadDictFromFile(s string) (*[][]byte, error) {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		return nil, err

	}

	var wds [][]byte
	var wd []byte
	var blankWd []byte
	ignoreTilSeparator := false

	for _, v := range b {
		if _, ok := separatorBytes[v]; ok {
			if !ignoreTilSeparator {
				wds = append(wds, wd)
			}
			wd = blankWd
			ignoreTilSeparator = false
		} else if _, ok := ignoreBytes[v]; ok {
			ignoreTilSeparator = true
			continue
		} else {
			wd = append(wd, v)
		}
	}

	return &wds, nil
}

//LoadDictFromStrings takes a slice of strings and turns them into a dictionary
func LoadDictFromStrings(s *[]string) *[][]byte {
	b := make([][]byte, 0)

	for _, v := range *s {
		b = append(b, []byte(v))
	}

	return &b
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
	dictBytes, err := LoadDictFromFile(filename)
	if err != nil {
		return w, err
	}

	words := make(map[string][]string)
	foundWords := make(map[string]bool)

	for _, v := range *dictBytes {
		word := string(v)
		_, found := foundWords[word]
		if found {
			debugPrint("Already found " + word)
			continue
		}

		debugPrint("New word: " + word)
		foundWords[word] = true

		spl := strings.Split(word, "")
		sort.Strings(spl)
		sorted := strings.Join(spl, "")

		_, ok := words[sorted]
		if ok {
			words[sorted] = append(words[sorted], word)

		} else {
			words[sorted] = []string{word}
		}
	}

	anags := make(map[string][]string)
	for k, v := range words {
		if len(v) > 1 {
			anags[k] = v
		}
	}

	return anags, nil

}
