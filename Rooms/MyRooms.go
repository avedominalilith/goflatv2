package Rooms

import (
	"fmt"
)

type RoomsStruct struct {
	Name      string
	heightInM float64
	widthInM  float64
	lengthInM float64
}

func InitRooms() RoomsStruct {
	var r RoomsStruct
	fmt.Println("Название комнаты: ")
	fmt.Scanln(&r.Name)
	fmt.Println("Высота: ")
	fmt.Scanln(&r.heightInM)
	fmt.Println("Ширина: ")
	fmt.Scanln(&r.widthInM)
	fmt.Println("Длина: ")
	fmt.Scanln(&r.lengthInM)
	return r
}

func InfoRooms(r RoomsStruct) {
	fmt.Printf("Название комнаты: %s, площадь: %.2f кв и, высота потолков: %.2f м\n", r.Name, r.lengthInM*r.widthInM, r.heightInM)
}

func CreateRooms() {
	var numRooms int
	fmt.Print("Количество комнат: ")
	fmt.Scanln(&numRooms)
	var rooms []RoomsStruct
	for i := 0; i < numRooms; i++ {
		room := InitRooms()
		rooms = append(rooms, room)
	}
	fmt.Println("\nКомнаты:")
	for _, room := range rooms {
		InfoRooms(room)
	}
}
