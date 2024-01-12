package People

import "fmt"

type StructPeople struct {
	fName string
	lName string
	Age   int
}

func InitPeople() StructPeople {
	var p StructPeople
	fmt.Println("Имя: ")
	fmt.Scanln(&p.fName)
	fmt.Println("Фамилия: ")
	fmt.Scanln(&p.lName)
	fmt.Println("Возраст: ")
	fmt.Scanln(&p.Age)
	return p
}

func InfoPeople(p StructPeople) {
	fmt.Printf("Проживающий: %s %s, возраст: %d\n", p.fName, p.lName, p.Age)
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
