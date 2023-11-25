package People

import "fmt"

type StructPeople struct {
	firstname string
	lastname  string
	age       int
}

func InitPeople() []StructPeople {
	person_1 := StructPeople{
		firstname: "Анастасия",
		lastname:  "Беброва",
		age:       30,
	}
	person_2 := StructPeople{
		firstname: "Александр",
		lastname:  "Брадков",
		age:       30,
	}
	return []StructPeople{person_1, person_2}
}

func ShowFamily(people []StructPeople) {
	fmt.Printf("Моя семья:\n")
	for _, person := range people {
		fmt.Println(person.firstname, person.lastname, person.age)
	}
}
