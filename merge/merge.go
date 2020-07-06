package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: " + os.Args[0] + "[file1.json] [file2.json]... [out.json]")
		os.Exit(1)
	}

	combined := make(map[string][]string)
	for _, fname := range os.Args[1 : len(os.Args)-1] {
		f, err := os.Open(fname)
		if err != nil {
			fmt.Printf("Coudln't load JSON file %s: %v\n", fname, err)
			os.Exit(1)
		}
		defer f.Close()
		jsonBytes, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Printf("Error reading contents of %s: %v\n", fname, err)
			os.Exit(1)
		}

		replacements := make(map[string][]string)
		json.Unmarshal(jsonBytes, &replacements)

		for k, v := range replacements {
			if len(k) == 0 {
				// Firefox appears to output a 0 length key
				continue
			}
			cv, exists := combined[k]
			if !exists {
				cv = replacements[k]
			} else {
			LOOP:
				for _, repl := range v {
					for _, test := range cv {
						if repl == test {
							continue LOOP
						}
					}
					v = append(v, repl)
				}
			}

			combined[k] = v
		}
	}

	outName := os.Args[len(os.Args)-1]
	j, err := json.MarshalIndent(combined, "", "  ")
	if err != nil {
		fmt.Printf("Couldn't marshall combined map to JSON: %v\n", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(outName, j, 0644)
	if err != nil {
		fmt.Printf("Failed to write JSON to file %s: %v\n", outName, err)
		os.Exit(1)
	}
}
