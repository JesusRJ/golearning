package main

type (
	Person struct {
		Name string
	}

	Team []Person
)

func main() {
	team := Team{
		Person{"John"},
		Person{"Jane"},
		Person{"Jack"},
		Person{"Jill"},
		Person{"Joe"},
	}

	print(team)
	addSurname(team)
	print(team)
}

func print(team Team) {
	for _, person := range team {
		println(person.Name)
	}
}

func addSurname(team Team) {
	for i, _ := range team {
		team[i].Name += " Doe"
	}
}
