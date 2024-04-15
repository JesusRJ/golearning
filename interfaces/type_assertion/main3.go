package main

import (
	"io"
	"os"
)

func main() {
	var w io.Writer         // io.Writer só possuí o método Writer
	w = os.Stdout           // os.Stdout retorna um *os.File
	rw := w.(io.ReadWriter) // sucesso: *os.File tem tanto o método Read quanto o método Write

	w = new(ByteCounter)
	rw = (io.ReadWriter) // panic: *ByteCounter não tem o método Read
}
