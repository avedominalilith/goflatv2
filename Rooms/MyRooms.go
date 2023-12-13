package Rooms

import "fmt"

type RoomsStruct struct {
	Name string
	h    int
	w    int
	l    int
}

func InitRooms() []RoomsStruct {
	rooms := []RoomsStruct{
		{"Гостинная", 4, 2, 6},
		{"Кухня", 4, 2, 5},
		{"Спальня", 4, 3, 5},
		{"Кабинет", 4, 3, 4},
		{"Ванная", 3, 3, 4},
	}
	return rooms
}

func ShowRooms(rooms []RoomsStruct) {
	fmt.Println("\nМой дом:")
	for _, room := range rooms {
		fmt.Printf("Комната: %s \nПлощадь: %d кв м\n", room.Name, room.w*room.l)
	}
}
