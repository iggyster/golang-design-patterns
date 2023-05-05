package creational

import "fmt"

type IBuilding interface {
	GetInfo()
}

type Building struct {
	Height int
	Floors int
}

type EmpireState struct {
	Building
}

func (es *EmpireState) GetInfo() {
	fmt.Printf("Is a %v-story Art Deco skyscraper in Midtown Manhattan, New York City\n", es.Building.Height)
	fmt.Printf("It has %v floors\n", es.Building.Floors)
}

func newEmpireState() IBuilding {
	return &EmpireState{Building: Building{Height: 381, Floors: 102}}
}

type WhiteHouse struct {
	Building
}

func (wh *WhiteHouse) GetInfo() {
	fmt.Println("Is a the official residence of the president of the United States located is Washington, D.C.")
	fmt.Printf("It's %v meters heigh and it has %v floors\n", wh.Building.Height, wh.Building.Floors)
}

func newWhiteHouse() IBuilding {
	return &WhiteHouse{Building: Building{Height: 21, Floors: 4}}
}

func Factory(t string) IBuilding {
	if t == "white house" {
		return newWhiteHouse()
	}

	return newEmpireState()
}
