package thesaurus

import (
	"fmt"
	"net/http"
)

const (
	// BaseURL is URL of thesaurus.com.
	BaseURL = "http://www.thesaurus.com/browse"
)

// Search is search from thesaurus.com, and returns words array.
func Search(w string) ([]Word, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", BaseURL, w))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	/*
		If the value is not found, threre returns "301 Moved Permanenly".
		Location: http://www.thesaurus.com/misspelling?term=misu
	*/
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("The word is not registered.")
	}

	words, err := GenerateWords(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(words) == 0 {
		return nil, fmt.Errorf("0 synonyms found")
	}

	return words, nil
}

// SynSort sorts Word array by relevance in ascending order, and returns sorted new slices.
func SynSort(ws []Word) []Word {
	return countingSort(ws, true)
}

// AntSort sorts Word array by relevance in descending order, and returns sorted new slices.
func AntSort(ws []Word) []Word {
	return countingSort(ws, false)
}

// countingSort sorts Word array by relevance.
// The data range of relevance is decided, so I use counting sort.
// (Insertion sort is maybe more faser because initial word's order is almost sorted.)
func countingSort(ws []Word, isAsc bool) []Word {
	type Data []Word
	// data range 3 - -3, so reserves size 7.
	dataCase := make([]Data, 7)

	offset := +3

	for _, v := range ws {
		dataCase[v.Relevance+offset] = append(dataCase[v.Relevance+offset], v)
	}

	sorted := make([]Word, 0, len(ws))
	if isAsc {
		for i := len(dataCase) - 1; i >= 0; i-- {
			for _, v := range dataCase[i] {
				sorted = append(sorted, v)
			}
		}
	} else {
		for i := 0; i < len(dataCase); i++ {
			for _, v := range dataCase[i] {
				sorted = append(sorted, v)
			}
		}
	}

	return sorted
}

// FilterRelevance extracts Words matched value of relevance.
func FilterRelevance(ws []Word, rel int) ([]Word, error) {
	words := make([]Word, 0, len(ws))

	if rel > 3 || rel < -3 || rel == 0 {
		return nil, fmt.Errorf("thesaurus.FilterRelevance: out of range \"%d\"", rel)
	}

	for _, w := range ws {
		if w.Relevance == rel {
			words = append(words, w)
		}
	}

	return words, nil
}

// FilterLength extracts Words matched value of lengthRate.
func FilterLength(ws []Word, length int) ([]Word, error) {
	words := make([]Word, 0, len(ws))

	if length > 3 || length < 1 {
		return nil, fmt.Errorf("thesaurus.FilterLength: out of range \"%d\"", length)
	}

	for _, w := range ws {
		if w.LengthRate == length {
			words = append(words, w)
		}
	}

	return words, nil
}

// FilterComplexity extracts Words matched value of complexity.
func FilterComplexity(ws []Word, comp int) ([]Word, error) {
	words := make([]Word, 0, len(ws))

	if comp > 3 || comp < 1 {
		return nil, fmt.Errorf("thesaurus.FilterComplexity: out of range \"%d\"", comp)
	}

	for _, w := range ws {
		if w.Complexity == comp {
			words = append(words, w)
		}
	}

	return words, nil
}

// FilterUseCase extracts Words matched value of useCase.
func FilterUseCase(ws []Word, use int) ([]Word, error) {
	words := make([]Word, 0, len(ws))

	if !(use == COMMON || use == INFORMAL || use == NONE) {
		return nil, fmt.Errorf("thesaurus.FilterUseCase: out of range \"%d\"", use)
	}

	for _, w := range ws {
		if w.UseCase == use {
			words = append(words, w)
		}
	}

	return words, nil
}

// Synonyms extracts synonym words.
func Synonyms(ws []Word) []Word {
	return filterClass(ws, SYNONYM)
}

// Antonyms extracts antonym words.
func Antonyms(ws []Word) []Word {
	return filterClass(ws, ANTONYM)
}

// filterClass extracts Words matched value of class.
func filterClass(ws []Word, class int) []Word {
	words := make([]Word, 0, len(ws))

	for _, w := range ws {
		if w.Class == class {
			words = append(words, w)
		}
	}

	return words
}
