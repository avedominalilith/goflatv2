package Furniture

import "fmt"

type StructFurniture struct {
	Name string
	Room string
	h    float64
	w    float64
	l    float64
}

func InitFurniture() StructFurniture {
	var f StructFurniture
	fmt.Println("Название предмета мебели: ")
	fmt.Scanln(&f.Name)
	fmt.Println("Комната: ")
	fmt.Scanln(&f.Room)
	fmt.Println("Высота: ")
	fmt.Scanln(&f.h)
	fmt.Println("Ширина: ")
	fmt.Scanln(&f.w)
	fmt.Println("Длина: ")
	fmt.Scanln(&f.l)
	return f
}

func InfoFurniture(f StructFurniture) {
	fmt.Printf("Название предмета: %s, комната: %s, занимаемая площадь: %.2f кв м, высота: %.2f м\n", f.Name, f.Room, f.l*f.w, f.h)
}

func ShowFurniture() {
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
