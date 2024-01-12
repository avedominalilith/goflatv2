package Rooms

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type RoomsStruct struct {
	Name string
	h    float64
	w    float64
	l    float64
}

func InitRooms() RoomsStruct {
	var r RoomsStruct
	fmt.Println("Название комнаты: ")
	fmt.Scanln(&r.Name)
	fmt.Println("Высота: ")
	fmt.Scanln(&r.h)
	fmt.Println("Ширина: ")
	fmt.Scanln(&r.w)
	fmt.Println("Длина: ")
	fmt.Scanln(&r.l)
	return r
}

func InfoRooms(r RoomsStruct) {
	fmt.Printf("Название комнаты: %s, площадь: %.2f кв и, высота потолков: %.2f м\n", r.Name, r.l*r.w, r.h)
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
