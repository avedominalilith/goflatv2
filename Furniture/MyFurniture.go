package Furniture

import "fmt"

type StructFurniture struct {
	Name string
	h    int
	w    int
	l    int
}

func InitFurniture() []StructFurniture {
	furniture := []StructFurniture{
		{"диван", 100, 120, 300},
		{"кресло", 100, 120, 80},
		{"стол", 90, 90, 90},
		{"шкаф", 200, 70, 100},
		{"полка", 50, 30, 100},
	}
	return furniture
}

func ShowFurniture(furniture []StructFurniture) {
	fmt.Println("\nМебель:")
	for _, ObjFurniture := range furniture {
		fmt.Printf("Предмет мебели: %s\nВысота: %d см\nШирина: %d см\nДлина: %d см\n", ObjFurniture.Name, ObjFurniture.h, ObjFurniture.w, ObjFurniture.l)
	}
}
