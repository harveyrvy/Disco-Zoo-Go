package jungle

import (
	animal "github.com/discozoo/Animal"

	utils "github.com/discozoo/Utils"
)

// Coords are:
// {amt down, amt right}

var Monkey = animal.NewAnimal("Monkey", utils.TilesList{
	{0, 0},
	{1, 1},
	{0, 2},
	{1, 3},
})

var Toucan = animal.NewAnimal("Toucan", utils.TilesList{
	{0, 0},
	{1, -1},
	{2, 0},
	{3, 0},
})

var Gorilla = animal.NewAnimal("Gorilla", utils.TilesList{
	{0, 0},
	{1, 0},
	{0, 2},
	{1, 2},
})

var Panda = animal.NewAnimal("Panda", utils.TilesList{
	{0, 0},
	{-1, 2},
	{1, 2},
})

var Tiger = animal.NewAnimal("Tiger", utils.TilesList{
	{0, 0},
	{0, 2},
	{0, 3},
})

var Phoenix = animal.NewAnimal("Phoenix", utils.TilesList{
	{0, 0},
	{2, 2},
})

var Lemur = animal.NewAnimal("Lemur", utils.TilesList{
	{0, 0},
	{1, 1},
	{2, 0},
})
