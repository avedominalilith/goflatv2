package MyFurniture

import "fmt"

type Furniture struct {
	name string
}

func FillFurniture() []Furniture {
	sofa := Furniture{
		name: "Диван",
	}
	shelf := Furniture{
		name: "Полка",
	}
	bed := Furniture{
		name: "Кровать",
	}
	desk := Furniture{
		name: "Стол",
	}
	armchair := Furniture{
		name: "Кресло",
	}
	return []Furniture{sofa, shelf, bed, armchair, desk}
}

func ShowFurniture(Furniture []Furniture) {
	fmt.Printf("Мебель в доме:\n")
	for _, furniture := range Furniture {
		fmt.Printf("%s\n", furniture.name)
	}
	fmt.Print("\n")
}
