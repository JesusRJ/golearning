# Context

# Dicas

## Use empty struct (`struct{}`) como chaves para valores de contexto em Go

> Referência: [Use struct{} as keys for context.Value() in Go](https://gist.github.com/ww9/4ad7b2ddfb94816a30dfdf2218e02f48)

### TLDR

Usar um struct vazia como chave para o value de um contexto evita dois problemas:

1. Alocação desnecessária;
2. Ambiguidade e margem para uso indevido;

Além disso pode ser uma escolha mais eficiente se comparado a `strings` e `uint8` [benchmarking](https://gist.github.com/SammyOina/6eb54babd618ab6a850e8f1af4f4ac7d).

```Go
type key struct{}
ctx = context.WithValue(ctx, key{}, "my value") // Set value
myValue, ok := ctx.Value(key{}).(string) // Get value
```

### Explicação detalhada

```Go
package main

import (
	"fmt"
	"unsafe"
)

type StructKey1 struct{}
type StructKey2 struct{}

type IntKey1 int
type IntKey2 int

// Why use empty struct{} and not int as context.Value() key types:
func main() {

	// First of all, plain string keys look simple but they do not protect context
	// values from being tampered by other packages:
	// Also they cause allocation, so you probably don't want to use them.
	compareKeys("dontUseString", "dontUseString") // true

	// Now let's see how struct{} types behave.
	// Same struct{} keys are always equal:
	compareKeys(StructKey1{}, StructKey1{}) // true
	// Different struct{} keys are never equal even if they appear to be of the same struct{} type:
	compareKeys(StructKey1{}, StructKey2{}) // false

	// This also applies to int keys. Same type means equal:
	compareKeys(IntKey1(0), IntKey1(0)) // true
	// And different keys are never equal even though they have 0 value:
	compareKeys(IntKey1(0), IntKey2(0)) // false
	// However, unlike struct{}, an int typed key allows for mistakes with the value:
	compareKeys(IntKey1(0), IntKey1(1)) // false
	// To add to why you shouldn't use int typed keys, when seeing that the package define a
	// key as int, the user might be tempted to pass a primitive 0 instead. Which doesn't work:
	compareKeys(IntKey1(0), 0) // false

	// Whereas with struct{} there's no ambiguity or margin for missuse.
	// So if for any reason your package wants to allow other packages to set/get a
	// context.Value() that's specific to your package, either export a struct{} key
	// like `type StructKey1 struct{}` or, even better, export methods that interact
	// with these values but keep the key unexported:

	// func WithUserIP(ctx context.Context, userIP net.IP) context.Context {
	// 	return context.WithValue(ctx, userIPKey{}, userIP)
	// }
	// func UserIPFromContext(ctx context.Context) (net.IP, bool) {
	// 	userIP, ok := ctx.Value(userIPKey{}).(net.IP)
	// 	return userIP, ok
	// }
}

func compareKeys(key1 interface{}, key2 interface{}) {
	type ifaceHdr struct {
		T unsafe.Pointer
		V unsafe.Pointer
	}

	fmt.Println("\nkey1 == key2?", key1 == key2)
	fmt.Printf("key1 %+v\n", *(*ifaceHdr)(unsafe.Pointer(&key1)))
	fmt.Printf("key2 %+v\n", *(*ifaceHdr)(unsafe.Pointer(&key2)))
}
```
