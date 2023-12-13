package House

import (
	"goflatv2/Appliances"
	"goflatv2/Furniture"
	"goflatv2/People"
	"goflatv2/Rooms"
)

func ShowHouse() {
	appliances := Appliances.InitAppliances()
	furniture := Furniture.InitFurniture()
	people := People.InitPeople()
	rooms := Rooms.InitRooms()

	Appliances.ShowAppliances(appliances)
	Furniture.ShowFurniture(furniture)
	People.ShowPeople(people)
	Rooms.ShowRooms(rooms)
}
