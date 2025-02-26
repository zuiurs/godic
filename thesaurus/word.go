package thesaurus

import (
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const (
	// class of words
	SYNONYM = iota
	ANTONYM
	// use case of words
	COMMON
	INFORMAL
	NONE
)

var (
	// re is thesaurus.com regular expression.
	//
	// Starting Tag's Struct
	// 	class
	// 		common-word
	// 		informal-word
	// 		nil
	// 	data-id
	// 		int
	// 	data-complexity
	// 		int
	// 	data-length
	// 		int
	re = regexp.MustCompile(`(<a href.*?data-id=".*?".*?data-complexity=".*?".*?data-length=".*?">)<span.*?>(.*?)</span>`)
)

// Word depends on thesaurus.com HTML source.
// memo: 7 WORD
type Word struct {
	// Word ID
	// Maybe don't be used
	id int

	// Word use case
	// common or informal word (if not matches, sets None)
	UseCase int

	// Word complexity
	// range: 1 - 3
	Complexity int

	// Word length
	// range: 1 - 3
	LengthRate int

	// Word class
	// synonym or antonym
	Class int

	// Word relevance
	// range: 3 - -3 (0 excepted)
	// More absolute value, more relevance.
	// (Maybe we can judge synonyms and antonyms by this value.
	// synonyms: 3 - 1, antonyms: -1 - -3)
	Relevance int

	// Word spell
	Spell string
}

// GenerateWords analyzes HTML body and returns Word struct array.
// This function can be only used in thesaurus.com HTML source.
// (depends on var "re" regexp)
func GenerateWords(r io.Reader) ([]Word, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	words := make([]Word, 0, 100)

	for i, tag := range re.FindAllSubmatch(b, -1) {
		tagAttrs, err := ParseHTMLStartTagBytes(tag[1])
		if err != nil {
			return nil, err
		}

		var w Word

		w.Spell = string(tag[2])

		for _, attr := range tagAttrs {
			switch attr.key {
			case "class": // "class" not equal Word.class
				if attr.value == "common-word" {
					w.UseCase = COMMON
				} else if attr.value == "informal-word" {
					w.UseCase = INFORMAL
				} else {
					w.UseCase = NONE
				}
			case "data-id":
				w.id, err = strconv.Atoi(attr.value)
			case "data-complexity":
				w.Complexity, err = strconv.Atoi(attr.value)
			case "data-length":
				w.LengthRate, err = strconv.Atoi(attr.value)
			case "data-category":
				rel := strings.TrimPrefix(strings.SplitN(attr.value, "&quot;", -1)[3], "relevant-")
				w.Relevance, err = strconv.Atoi(rel)
			}

			if err != nil {
				return nil, err
			}
		}

		/*
			peculiar specification :-(
			(i: 0, 1, 2, 3, 4, 5 ,6)
			id: 1, 2, 3, 4, 4, 5, 6
			1,2,3,4 is SYN, 4,5,6 is ANT.
		*/
		if i == w.id {
			w.Class = ANTONYM
		} else {
			w.Class = SYNONYM
		}

		if i <= w.id {
			words = append(words, w)
		}
	}

	return words, nil
}

func (w Word) String() string {
	return w.Spell
}

// VerboseString outputs verbose information.
func (w Word) VerboseString() string {
	var use, class string
	if w.UseCase == COMMON {
		use = "Common"
	} else if w.UseCase == INFORMAL {
		use = "Informal"
	} else {
		use = "None"
	}

	if w.Class == SYNONYM {
		class = "Synonym"
	} else {
		class = "Antonym"
	}

	return fmt.Sprintf("ID: %d\nUseCase: %s\nComplexity: %d\nLengthRate: %d\nClass: %s\nRelevance: %d\nSpell: %s\n",
		w.id, use, w.Complexity,
		w.LengthRate, class, w.Relevance, w.Spell)
}
