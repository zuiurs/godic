package thesaurus

import (
	"fmt"
	"strings"
	"text/scanner"
)

// Attr is HTML tag attribute.
type Attr struct {
	key   string
	value string
}

// ParseHTMLStartTagBytes parses HTML Starting tag and returns attribute array of the tag from byte array.
func ParseHTMLStartTagBytes(b []byte) ([]Attr, error) {
	return ParseHTMLStartTagString(string(b))
}

// ParseHTMLStartTagString parses HTML Starting tag and returns attribute array of the tag from string.
func ParseHTMLStartTagString(s string) ([]Attr, error) {
	r := strings.NewReader(s)
	var sc scanner.Scanner
	sc.Init(r)

	var tok rune
	// ignore any token before starting tag
	for tok != rune('<') {
		tok = sc.Scan()
	}

	tok = sc.Scan() // scan tag identifier
	tok = sc.Scan() // scan top attribute

	attr := make([]Attr, 0, 10)

	// TODO: TokenText の返り値が string なのでこうしているが、tokPos, tokEnd あたりを使えばバイト配列で行けるかもしれない
	var bufKey, bufValue string

	var loopCount int

	for ; tok != rune('>'); tok = sc.Scan() {
		for ; tok != rune('='); tok = sc.Scan() {
			bufKey += sc.TokenText()

			if loopCount++; loopCount > 10000 {
				return nil, fmt.Errorf("Error: Uninspected tag format.")
			}
		}
		tok = sc.Scan() // skip equal
		bufValue = sc.TokenText()
		bufValue = bufValue[1 : len(bufValue)-1] // remove quatation

		attr = append(attr, Attr{
			key:   bufKey,
			value: bufValue,
		})

		bufKey, bufValue = "", ""
	}

	return attr, nil
}

func (a Attr) String() string {
	return fmt.Sprintf("(%s, %s)", a.key, a.value)
}

// func main() {
// 	str := `<a href="http://www.thesaurus.com/browse/grant" class="common-word" data-id="1" data-category="{&quot;name&quot;: &quot;relevant-3&quot;, &quot;color&quot;: &quot;#fcbb45&quot;}" data-complexity="1" data-length="1"><span class="text">grant</span>`
//
// 	attr, _ := ParseHTMLStartTagString(str)
//
// 	fmt.Println(attr)
// }
