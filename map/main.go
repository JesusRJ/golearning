package main

import "fmt"

type User struct {
	Address map[int]string
}

func (u *User) Addresses() map[int]string {
	// Para evitar alterar o ponteiro do map, fazemos um snapshot
	// Serve fazer uma cópia tanto na saída quanto na entrada para evitar alterações externas
	snapshot := make(map[int]string, len(u.Address))
	for k, v := range u.Address {
		snapshot[k] = v
	}
	return snapshot
}

func (u *User) GetAddress(key int) string {
	a, ok := u.Address[key]
	if !ok {
		return ""
	}
	return a
}

// Olhar no fw NoOperations (no-op)
func (u *User) DelAddress(key int) {
	// Delete é safe pois sempre executa sem falhas
	delete(u.Address, key)
}

func main() {
	// Mais indicado para qnd se vai inicializar o map com valores
	m := map[int]string{}
	fmt.Printf("%+v\n", m)

	// Indicado para qnd não se conhece os valores iniciais. Se souber o tamanho é ideal informar o len
	m2 := make(map[int]string)
	fmt.Printf("%+v\n", m2)

	u := User{
		Address: map[int]string{
			1: "Teste1",
		},
	}

	fmt.Printf("%+v\n", u)

	// Maps por natureza são ponteiros
	addr := u.Addresses()
	addr[2] = "Teste 2"

	fmt.Printf("%+v\n", u)
}
