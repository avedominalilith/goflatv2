package Appliances

import "fmt"

type StructAppliances struct {
	Name string
	Room string
}

func InitAppliances() StructAppliances {
	var a StructAppliances
	fmt.Println("Название устройства: ")
	fmt.Scanln(&a.Name)
	fmt.Println("Комната: ")
	fmt.Scanln(&a.Room)
	return a
}

func InfoAppliances(a StructAppliances) {
	fmt.Printf("Название устройства: %s, комната: %s\n", a.Name, a.Room)
}

func ShowAppliances() {
	var numAppliances int
	fmt.Print("Количество девайсов: ")
	fmt.Scanln(&numAppliances)
	var appliances []StructAppliances
	for i := 0; i < numAppliances; i++ {
		appliance := InitAppliances()
		appliances = append(appliances, appliance)
	}
	fmt.Println("\nТехника:")
	for _, appliance := range appliances {
		InfoAppliances(appliance)
	}
}
