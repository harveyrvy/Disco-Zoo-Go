package polar

import (
	animal "github.com/discozoo/Animal"

	utils "github.com/discozoo/Utils"
)

// Coords are:
// {amt down, amt right}

var Penguin = animal.NewAnimal("Penguin", utils.TilesList{
	{0, 0},
	{1, 0},
	{2, -1},
	{2, 1},
})

var Seal = animal.NewAnimal("Seal", utils.TilesList{
	{0, 0},
	{1, 1},
	{2, 2},
	{1, 3},
})

var Muskox = animal.NewAnimal("Muskox", utils.TilesList{
	{0, 0},
	{0, 1},
	{1, 0},
	{1, 2},
})

var PolarBear = animal.NewAnimal("PolBear", utils.TilesList{
	{0, 0},
	{0, 2},
	{1, 2},
})

var Walrus = animal.NewAnimal("Walrus", utils.TilesList{
	{0, 0},
	{1, 1},
	{1, 2},
})

var Yeti = animal.NewAnimal("Yeti", utils.TilesList{
	{0, 0},
	{2, 0},
})

var SnowyOwl = animal.NewAnimal("SnowyOwl", utils.TilesList{
	{0, 0},
	{0, 1},
	{1, 1},
})
