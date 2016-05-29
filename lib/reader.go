package lib

import (
	"errors"
)

// Read Only
type Reader struct {
	data     []byte
	index    int
	capacity int
}

func NewReader(data []byte) *Reader {
	r := new(Reader)
	r.data = data
	r.index = 0
	r.capacity = len(data)

	return r
}

func (r *Reader) ReadLine() (string, error) {
	if r.index >= r.capacity {
		return "", errors.New("Reach the EOF")
	}

	current_char := r.data[r.index]

	var b []byte

	for current_char != '\n' {
		b = append(b, current_char)
		r.index++
		current_char = r.data[r.index]
	}

	r.index++

	str := string(b)

	return str, nil
}
