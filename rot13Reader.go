package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	reader io.Reader
}

func (r rot13Reader) Read(bytes []byte) (int, error) {
	read := 0
	for {
		n, err := r.reader.Read(bytes)
		if n > 0 {
			for i := read; i < read + n; i++ {
				if bytes[i] > 'm' {
					bytes[i] = (bytes[i] - 'm' - 1) +'a'
				} else if bytes[i] > 'Z' {
					bytes[i] = bytes[i] + 13
				} else if bytes[i] > 'M' {
					bytes[i] = (bytes[i] - 'M' - 1) +'A'
				} else if bytes[i] >= 'A' {
					bytes[i] = bytes[i] + 13
				}
			}
			
			read += n
		} else {
			return read, err
		}
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
