package savannah

import (
	animal "github.com/discozoo/Animal"

	utils "github.com/discozoo/Utils"
)

// Coords are:
// {amt down, amt right}

var Zebra = animal.NewAnimal("Zebra", utils.TilesList{
	{0, 0},
	{-1, 1},
	{0, 2},
	{1, 1},
})

var Hippo = animal.NewAnimal("Hippo", utils.TilesList{
	{0, 0},
	{0, 2},
	{2, 0},
	{2, 2},
})

var Giraffe = animal.NewAnimal("Giraffe", utils.TilesList{
	{0, 0},
	{1, 0},
	{2, 0},
	{3, 0},
})

var Lion = animal.NewAnimal("Lion", utils.TilesList{
	{0, 0},
	{0, 1},
	{0, 2},
})

var Elephant = animal.NewAnimal("Elephant", utils.TilesList{
	{0, 0},
	{0, 1},
	{1, 0},
})

var Gryphon = animal.NewAnimal("Gryphon", utils.TilesList{
	{0, 0},
	{0, 2},
	{1, 1},
})

var Rhino = animal.NewAnimal("Rhino", utils.TilesList{
	{0, 0},
	{-1, 1},
	{1, 1},
})
