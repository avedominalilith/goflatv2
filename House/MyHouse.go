package House

import (
	"goflatv2/Appliances"
	"goflatv2/Furniture"
	"goflatv2/People"
	"goflatv2/Rooms"
)

func ShowHouse() {
	Appliances.ShowAppliances()
	Furniture.ShowFurniture()
	People.ShowPeople()
	Rooms.ShowRooms()

}
