package mars

import (
	animal "github.com/discozoo/Animal"

	utils "github.com/discozoo/Utils"
)

// Coords are:
// {amt down, amt right}

var Marsmot = animal.NewAnimal("Marsmot", utils.TilesList{
	{0, 0},
	{1, 0},
	{2, 0},
	{2, -1},
})

var Marsmoset = animal.NewAnimal("Mrsmoset", utils.TilesList{
	{0, 0},
	{0, 2},
	{1, 2},
	{2, 1},
})

var Rock = animal.NewAnimal("Rock", utils.TilesList{
	{0, 0},
	{1, 0},
	{0, 1},
	{1, 1},
})

var Rover = animal.NewAnimal("Rover", utils.TilesList{
	{0, 0},
	{-1, 1},
	{0, 2},
})

var Martian = animal.NewAnimal("Martian", utils.TilesList{
	{0, 0},
	{1, 1},
	{0, 2},
})

var Marsmallow = animal.NewAnimal("Mallow", utils.TilesList{
	{0, 0},
	{2, 0},
})

var Marsten = animal.NewAnimal("Marsten", utils.TilesList{
	{0, 0},
	{0, 2},
	{1, 3},
})
