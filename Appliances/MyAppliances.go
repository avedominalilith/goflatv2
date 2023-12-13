package Appliances

import "fmt"

type StructAppliances struct {
	Name string
	Room string
}

func InitAppliances() []StructAppliances {
	appliances := []StructAppliances{
		{"телевизор", "Гостинная"},
		{"компьютер", "Кабинет"},
		{"стиральная машина", "Ванная"},
		{"холодильник", "Кухня"},
		{"чайник", "Кухня"},
	}
	return appliances
}

func ShowAppliances(appliances []StructAppliances) {
	fmt.Println("\nТехника:")
	for _, ObjAppliances := range appliances {
		fmt.Printf("%s (%s)\n", ObjAppliances.Name, ObjAppliances.Room)
	}
}
