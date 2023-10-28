package ShowFlat

import (
	"goflatv2/MyAppliances"
	"goflatv2/MyFlat"
	"goflatv2/MyFurniture"
	"goflatv2/MyPeople"
)

func ShowFlat() {
	appliances := MyAppliances.InitAppliances()
	furniture := MyFurniture.FillFurniture()
	people := MyPeople.InitPeople()
	flat := MyFlat.InsideRooms()

	MyAppliances.ShowAppliances(appliances)
	MyFurniture.ShowFurniture(furniture)
	MyPeople.ShowFamily(people)
	MyFlat.ShowRoomsSizes(flat)
}
