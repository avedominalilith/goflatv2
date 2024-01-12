package People

import "fmt"

type StructPeople struct {
	firstName string
	lastName  string
	Age       int
}

func InitPeople() StructPeople {
	var p StructPeople
	fmt.Println("Имя: ")
	fmt.Scanln(&p.firstName)
	fmt.Println("Фамилия: ")
	fmt.Scanln(&p.lastName)
	fmt.Println("Возраст: ")
	fmt.Scanln(&p.Age)
	return p
}

func InfoPeople(p StructPeople) {
	fmt.Printf("Проживающий: %s %s, возраст: %d\n", p.firstName, p.lastName, p.Age)
}

func CreatePeople() {
	var numPeople int
	fmt.Print("Количество обитателей дома: ")
	fmt.Scanln(&numPeople)
	var people []StructPeople
	for i := 0; i < numPeople; i++ {
		person := InitPeople()
		people = append(people, person)
	}
	fmt.Println("\nОбитатели:")
	for _, person := range people {
		InfoPeople(person)
	}
}
