package main

type MyReader struct{}

func (mr MyReader) Read(in []byte) (int, error) {
	for i, _ := range in {
		in[i] = 'A'
	}
	return len(in), nil
	// errors.New("Error")
}

// https://blog.golang.org/error-handling-and-go
// https://www.golang-book.com/books/intro/9
