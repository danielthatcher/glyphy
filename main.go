package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/gobuffalo/packr/v2"
	flag "github.com/spf13/pflag"
)

type replacement struct {
	// The index in the original string to start replacing at
	index int

	// The characters in the original string to replace
	characters string
}

// All generates all possible homoglyphs for the replacement. newIndices is a list whose i-th element gives the index in the
// homoglyph string which i would have referred to in the original string. This can be passed to r.All for another replacement
// to allow for multiple replacements to keep their own string
//
// If the string is unmodified, then nil can be passed in for the indices
//
// All performs no checks to see if replacements are compatible with previous replacements, e.g. the replacements for "m" and
// "p.m." are incompatible.
//
// N.B. s will be modified by usage of append
func (r *replacement) All(s []rune, indices []int) (homoglyphs [][]rune, newIndices []int) {
	// If indices not supplied, use the identity map
	if indices == nil {
		indices = make([]int, len(s))
		for i := range indices {
			indices[i] = i
		}
	}

	// Do all homoglyph replacements
	l := len(r.characters)
	realIndex := indices[r.index]
	homoglyphs = make([][]rune, len(replacements[r.characters]))
	for i, v := range replacements[r.characters] {
		new := append(s[:realIndex], v...)
		new = append(new, s[realIndex+l:]...)
		homoglyphs[i] = make([]rune, len(new))
		copy(homoglyphs[i], new)
	}

	// Update newIndices to be correct
	if l > 1 {
		update := make([]int, l)
		for i := range update {
			update[i] = realIndex
		}
		newIndices = append(indices[:r.index], update...)

		// How many characters shorter the string now is
		diff := l - 1
		for i := range indices[l+1:] {
			newIndices = append(newIndices, indices[i+l+1]-diff)
		}
	} else {
		newIndices = indices
	}

	return
}

// IsCompatible returns whether two replacements can be applied to the same base string without modifying overlapping areas.
func (r *replacement) IsCompatible(w *replacement) bool {
	rl := len(r.characters)
	wl := len(w.characters)
	return ((w.index + wl) <= r.index) || ((r.index + rl) <= w.index)
}

func RecursiveReplace(target []rune, indices []int, startIndex int, depth int, applied []*replacement, out chan<- []rune) {
	if depth <= 0 {
		return
	}

OUTER:
	for i := range positions {
		if i < startIndex {
			continue
		}

		r := positions[i]
		targetCopy := make([]rune, len(target))
		copy(targetCopy, target)

		// Check if we can apply it
		for j := range applied {
			if !applied[j].IsCompatible(r) {
				continue OUTER
			}
		}

		// Send for this level
		glyphs, newIndices := r.All(targetCopy, indices)
		for k := range glyphs {
			out <- glyphs[k]
		}

		// Recurse
		newApplied := make([]*replacement, len(applied)+1)
		copy(newApplied, applied)
		newApplied[len(newApplied)-1] = r
		for k := range glyphs {
			RecursiveReplace(glyphs[k], newIndices, i, depth-1, newApplied, out)
		}
	}

	if len(applied) == 0 {
		close(out)
	}
}

var replacements map[string][][]rune
var positions []*replacement
var target []rune

func main() {
	var depth int
	var urlencode bool

	flag.IntVarP(&depth, "max-replacements", "n", 1, "The maximum number of positions to replace in")
	flag.BoolVarP(&urlencode, "urlencode", "u", false, "URL encode special characters in the data")
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	target := flag.Arg(0)

	// Load the replacements dictionary from a JSON file included using packr2
	box := packr.New("pwd", "./")
	jsonBytes, err := box.Find("replacements.json")
	if err != nil {
		fmt.Printf("Failed to find replacements.json file: %v.\nWas this binary built using packr2?", err)
		os.Exit(1)
	}
	replacements = make(map[string][][]rune)
	replacementsStr := make(map[string][]string)
	json.Unmarshal(jsonBytes, &replacementsStr)
	for k, v := range replacementsStr {
		runes := make([][]rune, len(v))
		for i := range runes {
			runes[i] = []rune(v[i])
		}
		replacements[k] = runes
	}

	// Create a map of the indexes of our replacement target occurrences in the string
	positions = make([]*replacement, 0)
	for k := range replacements {
		i := -1
		for {
			loc := strings.Index(target[i+1:], k)
			if loc == -1 {
				break
			}
			i += loc + 1
			r := &replacement{i, k}
			positions = append(positions, r)
		}
	}

	applied := make([]*replacement, 0)
	out := make(chan []rune)

	go RecursiveReplace([]rune(target), nil, 0, depth, applied, out)
	for r := range out {
		s := string(r)
		if urlencode {
			s = url.QueryEscape(s)
		}
		fmt.Println(s)
	}
}
