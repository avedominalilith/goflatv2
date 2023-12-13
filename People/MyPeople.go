package People

import "fmt"

type StructPeople struct {
	fName string
	lName string
	Age   int
}

func InitPeople() []StructPeople {
	people := []StructPeople{
		{"Томас", "Шелби", 55},
		{"Игорь", "Крутой", 69},
	}
	return people
}

func ShowPeople(people []StructPeople) {
	fmt.Println("\nЖители дома:")
	for _, person := range people {
		fmt.Printf("Имя: %s\nФамилия: %s\nВозраст: %d\n", person.fName, person.lName, person.Age)
	}
}
