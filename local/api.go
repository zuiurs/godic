package local

import (
	"bufio"
	"bytes"
	"errors"
	"strings"
)

//go-bindata.exe -o local/data.go -pkg local data/dict_data.txt

const (
	// Data is a path of dict_data.txt from application root.
	Data = "data/dict_data.txt"
)

var (
	// Dict is cache of dictionary map.
	Dict = make(map[string]string)
)

// Search is search from local dictionary, and returns words array.
func Search(w string) (string, error) {
	// If have a cache, don't generate dictionary map.
	if len(Dict) == 0 {
		data, _ := Asset(Data)
		Dict = generateDict(data)
	}

	if v, ok := Dict[w]; ok {
		return v, nil
	}
	return "", errors.New("The word is not registered.")
}

func generateDict(data []byte) map[string]string {
	dict := make(map[string]string)

	sc := bufio.NewScanner(bytes.NewReader(data))

	var line []string
	var words []string
	var desc string

	var buf string
	for sc.Scan() {
		buf = sc.Text()

		// word, word \t description
		line = strings.SplitN(buf, "\t", 2)

		words = strings.Split(line[0], ",")
		desc = strings.TrimSpace(line[1])

		for _, word := range words {
			dict[strings.TrimSpace(word)] = desc
		}
	}

	return dict
}
