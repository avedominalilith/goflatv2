package ShowFlat

import (
	"goflatv2/Appliances"
	"goflatv2/Flat"
	"goflatv2/Furniture"
	"goflatv2/People"
)

func ShowFlat() {
	appliances := Appliances.InitAppliances()
	furniture := Furniture.FillFurniture()
	people := People.InitPeople()
	flat := Flat.InsideRooms()

	Appliances.ShowAppliances(appliances)
	Furniture.ShowFurniture(furniture)
	People.ShowFamily(people)
	Flat.ShowRoomsSizes(flat)
}
