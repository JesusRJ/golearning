package main

import "fmt"

type Item struct {
	LockID string `dynamodbav:"LockID"`
	Digest string `dynamodbav:"Digest"`
}

func main() {
	// Testa se duas variáveis são iguais

	item1 := &Item{LockID: "lockid1", Digest: "teste1"}
	item2 := &Item{LockID: "lockid1", Digest: "teste"}

	if *item1 == *item2 {
		fmt.Println("São iguais")
	} else {
		fmt.Println("Não são iguais")
	}

}
