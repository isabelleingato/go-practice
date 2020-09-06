package main

import "io"

type rot13Reader struct {
	r io.Reader
}

// Credit with modifications: https://github.com/golang/tour/blob/master/solutions/rot13.go#L15
func rot13(b byte) byte {
	var a byte
	switch {
		// Is b in lowercase ASCII range?
		case 'a' <= b && b <= 'z':
			a = 'a'
		// Is b in uppercase ASCII range?
		case 'A' <= b && b <= 'Z':
			a = 'A'
		default:
			return b
	}
	//return (b-a+13)%(z-a+1) + a
	return (b-a+13)%26 + a
}

func (rReader rot13Reader) Read(in []byte) (int, error) {
	l, err := rReader.r.Read(in)
	for i, value := range in {
		in[i] = rot13(value)
	}
	return l, err
}