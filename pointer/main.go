package main

import "fmt"

type User struct {
	Name string
	Pass string
}

func (u *User) UpdateName(name string) {
	u.Name = name
}

func main() {
	u := User{
		Name: "Jesus",
		Pass: "p@ss",
	}

	fmt.Printf("%+v\n", u)
	Notify(&u)
	fmt.Printf("%+v\n", u)

	u.UpdateName("Reginaldo")
	fmt.Printf("%+v\n", u)
}

func Notify(u *User) {
	u.Name = "Reginaldo"
}
