# Compiling Go to WebAssembly

* [Instructions](https://github.com/golang/go/wiki/WebAssembly#getting-started)
* GOOS=js GOARCH=wasm go build -o main.wasm
* Move main.wasm to the subdirectory
* In httpserver dir, go run http.go
* localhost:8080/index.html, view console logs