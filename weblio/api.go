package weblio

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var (
	// BaseURL is URL of ejje.weblio.jp.
	BaseURL = "http://ejje.weblio.jp/content"

	re = regexp.MustCompile(`<td class=content-explanation>(.*?)</td>`)
)

// Search is search from ejje.weblio.jp, and returns word's array.
func Search(w string) ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", BaseURL, w))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	/*
		NOTE: If the value is not found, threre doesn't store the cookie "bhEjje".
		***BROWSER ONLY***
	*/

	words, err := extractWords(resp.Body)
	if err != nil {
		return nil, err
	}

	return words, nil
}

// extractWords extracts the results and returns the one's array.
func extractWords(r io.Reader) ([]string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	match := re.FindSubmatch(b)
	if len(match) == 0 {
		return nil, fmt.Errorf("The word is not found.")
	}

	w := string(match[1])

	var words []string

	// delimiter: "、" or "; "
	for _, r := range w {
		if r == '、' {
			words = strings.Split(w, "、")
			break
		} else if r == ';' {
			words = strings.Split(w, ";")
			break
		}
	}

	for i, v := range words {
		words[i] = strings.TrimSpace(v)
	}

	return words, nil
}
