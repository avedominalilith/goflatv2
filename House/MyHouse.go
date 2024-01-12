package House

import (
	"goflatv2/Appliances"
	"goflatv2/Furniture"
	"goflatv2/People"
	"goflatv2/Rooms"
)

func CreateHouse() {
	Appliances.CreateAppliances()
	Furniture.CreateFurniture()
	People.CreatePeople()
	Rooms.CreateRooms()

}
