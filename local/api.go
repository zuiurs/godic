package local

import (
	"errors"
	"strings"

	"github.com/zuiurs/godic/lib"
)

//go-bindata.exe -o local/data.go -pkg local data/dict_data.txt

const (
	DATA = "data/dict_data.txt"
)

var (
	Dict = make(map[string]string)
)

func Search(w string) (string, error) {
	if len(Dict) == 0 {
		data, _ := Asset(DATA)
		Dict = generateDict(data)
	}

	if v, ok := Dict[w]; ok {
		return v, nil
	}
	return "", errors.New("The word is not registered.")
}

func generateDict(data []byte) map[string]string {

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
