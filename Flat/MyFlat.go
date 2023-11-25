package Flat

import (
	"fmt"
)

type Rooms struct {
	name  string
	value int
}

func InsideRooms() []Rooms {
	floor := Rooms{
		name:  "Этаж: ",
		value: 1,
	}
	square := Rooms{
		name:  "Площадь, кв.м: ",
		value: 40,
	}
	numRooms := Rooms{
		name:  "Количество комнат: ",
		value: 2,
	}
	return []Rooms{floor, square, numRooms}
}

func ShowRoomsSizes(Rooms []Rooms) {
	fmt.Printf("Моя квартира:\n")
	for _, Rooms1 := range Rooms {
		fmt.Printf("%s\n %d\n", Rooms1.name, Rooms1.value)
	}
	fmt.Print("\n")
}
