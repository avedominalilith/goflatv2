package MyAppliances

import "fmt"

type Appliances struct {
	name string
}

func InitAppliances() []Appliances {
	playstation := Appliances{
		name: "плэйстэйшен",
	}
	computer := Appliances{
		name: "компьютер",
	}
	vacuum := Appliances{
		name: "пылесос",
	}
	fridge := Appliances{
		name: "холодильник",
	}
	electrokettle := Appliances{
		name: "электрочайник",
	}
	return []Appliances{playstation, computer, vacuum, fridge, electrokettle}
}

func ShowAppliances(Appliances []Appliances) {
	fmt.Printf("Техника в доме:\n")
	for _, appliances := range Appliances {
		fmt.Println(appliances.name, " ")
	}
}
