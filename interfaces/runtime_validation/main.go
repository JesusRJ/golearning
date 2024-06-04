package main

import (
	"fmt"
)

// ValueEncoder é uma interface de exemplo
type ValueEncoder interface {
	Encode(value interface{}) ([]byte, error)
}

// StructCodec é uma struct de exemplo que implementa ValueEncoder
type StructCodec struct{}

// Encode implementa o método Encode da interface ValueEncoder
func (s *StructCodec) Encode(value interface{}) ([]byte, error) {
	return []byte(fmt.Sprintf("%v", value)), nil
}

// Verificações em tempo de compilação
var _ ValueEncoder = &StructCodec{}
var _ ValueEncoder = (*StructCodec)(nil)

func main() {
	encoder := &StructCodec{}
	data, _ := encoder.Encode("example")
	fmt.Println(string(data)) // Output: example
}
