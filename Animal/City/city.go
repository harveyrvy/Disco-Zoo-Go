package city

import (
	animal "github.com/discozoo/Animal"

	utils "github.com/discozoo/Utils"
)

// Coords are:
// {amt down, amt right}

var Raccoon = animal.NewAnimal("Raccoon", utils.TilesList{
	{0, 0},
	{1, 0},
	{0, 2},
	{1, 3},
})

var Pigeon = animal.NewAnimal("Pigeon", utils.TilesList{
	{0, 0},
	{1, 1},
	{2, 1},
	{2, 2},
})

var Rat = animal.NewAnimal("Rat", utils.TilesList{
	{0, 0},
	{0, 1},
	{1, 1},
	{1, 3},
})

var Squirrel = animal.NewAnimal("Squirrel", utils.TilesList{
	{0, 0},
	{1, 1},
	{-1, 2},
})

var Opossum = animal.NewAnimal("Opossum", utils.TilesList{
	{0, 0},
	{1, 0},
	{1, 2},
})

var SewerTurtle = animal.NewAnimal("Turtle", utils.TilesList{
	{0, 0},
	{0, 1},
})

var Chipmunk = animal.NewAnimal("Chipmunk", utils.TilesList{
	{0, 0},
	{-1, 1},
	{0, 3},
})
