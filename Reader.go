package main

import (
	"strings"
)

type MyReader struct{}

// Reader, infinite stream of 'A's.

func (myReader MyReader) Read(bytes []byte) (int, error) {
	r := strings.NewReader(strings.Repeat("A", len(bytes)))
	return r.Read(bytes)
}

func main() {

	//reader.Validate(MyReader{})
}