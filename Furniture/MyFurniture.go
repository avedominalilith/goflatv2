package Furniture

import "fmt"

type StructFurniture struct {
	Name      string
	Room      string
	heightInM float64
	widthInM  float64
	lengthInM float64
}

func InitFurniture() StructFurniture {
	var f StructFurniture
	fmt.Println("Название предмета мебели: ")
	fmt.Scanln(&f.Name)
	fmt.Println("Комната: ")
	fmt.Scanln(&f.Room)
	fmt.Println("Высота: ")
	fmt.Scanln(&f.heightInM)
	fmt.Println("Ширина: ")
	fmt.Scanln(&f.widthInM)
	fmt.Println("Длина: ")
	fmt.Scanln(&f.lengthInM)
	return f
}

func InfoFurniture(f StructFurniture) {
	fmt.Printf("Название предмета: %s, комната: %s, занимаемая площадь: %.2f кв м, высота: %.2f м\n", f.Name, f.Room, f.lengthInM*f.widthInM, f.heightInM)
}

func CreateFurniture() {
	var numFurniture int
	fmt.Print("Количество предметов мебели: ")
	fmt.Scanln(&numFurniture)
	var furniture []StructFurniture
	for i := 0; i < numFurniture; i++ {
		objFurniture := InitFurniture()
		furniture = append(furniture, objFurniture)
	}
	fmt.Println("\nМебель:")
	for _, objFurniture := range furniture {
		InfoFurniture(objFurniture)
	}
}
