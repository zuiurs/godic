package main

import (
	"flag"
	"fmt"
	"strings"
	"./lib"
	"./data"
)

const (
	DATA = "dict_data.txt"
)

func main() {
	data, _ := data.Asset(DATA)

	dict := GenerateDict(data)

	flag.Parse()
	words := flag.Args()

	if len(words) == 0 {
		fmt.Println("Error: Required at least an arguments")
		return
	}

	var state string
	var ok bool

	for _, word := range words {
		state, ok = dict[word]
		if ok {
			fmt.Printf("%s: %s\n", word, state)
		} else {
			fmt.Printf("%s: NODATA\n", word)
		}
	}
}

func GenerateDict(data []byte) map[string]string {

	dict := map[string]string{}

	r := lib.NewReader(data)

	var line []string
	var keys []string
	var state string

	for str, err := r.ReadLine(); err == nil; {
		line = strings.SplitN(str, "\t", 2)

		keys = strings.Split(line[0], ",")
		state = line[1]

		for _, key := range keys {
			dict[key] = state
		}

		str, err = r.ReadLine()
	}

	return dict
}
