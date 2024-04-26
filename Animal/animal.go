package Animal

import (
	utils "github.com/discozoo/Utils"
)

type Animal struct {
	name  string
	tiles utils.TilesList
}

func NewDefault() Animal {
	return Animal{"X", utils.NewTilesList()}
}

func NewAnimal(name string, tiles utils.TilesList) Animal {
	return Animal{name, tiles}
}

func (a *Animal) SetName(name string) {
	a.name = name
}

func (a *Animal) GetName() string {
	return a.name
}

func (a *Animal) SetTiles(tiles utils.TilesList) {
	a.tiles = tiles
}

func (a *Animal) GetTiles() utils.TilesList {
	return a.tiles
}

func (a *Animal) String() string {
	return a.name
}
