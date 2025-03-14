package farm

import (
	animal "github.com/discozoo/Animal"

	utils "github.com/discozoo/Utils"
)

// Coords are:
// {amt down, amt right}

var Rabbit = animal.NewAnimal("Rabbit", utils.TilesList{
	{0, 0},
	{1, 0},
	{2, 0},
	{3, 0},
})

var Sheep = animal.NewAnimal("Sheep", utils.TilesList{
	{0, 0},
	{0, 1},
	{0, 2},
	{0, 3},
})

var Pig = animal.NewAnimal("Pig", utils.TilesList{
	{0, 0},
	{0, 1},
	{1, 0},
	{1, 1},
})

var Horse = animal.NewAnimal("Horse", utils.TilesList{
	{0, 0},
	{1, 0},
	{2, 0},
})

var Cow = animal.NewAnimal("Cow", utils.TilesList{
	{0, 0},
	{0, 1},
	{0, 2},
})

var Unicorn = animal.NewAnimal("Unicorn", utils.TilesList{
	{0, 0},
	{1, 1},
	{1, 2},
})

// origin : bottom left corner
var Chicken = animal.NewAnimal("Chicken", utils.TilesList{
	{0, 0},
	{-1, 1},
	{-2, 2},
})
