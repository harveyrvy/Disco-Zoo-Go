package mountain

import (
	animal "github.com/discozoo/Animal"

	utils "github.com/discozoo/Utils"
)

// Coords are:
// {amt down, amt right}

var Goat = animal.NewAnimal("Goat", utils.TilesList{
	{0, 0},
	{1, 0},
	{1, 1},
	{1, 2},
})

var Cougar = animal.NewAnimal("Cougar", utils.TilesList{
	{0, 0},
	{1, 1},
	{2, 0},
	{2, 2},
})

var Elk = animal.NewAnimal("Elk", utils.TilesList{
	{0, 0},
	{1, 1},
	{0, 2},
	{1, 2},
})

var Eagle = animal.NewAnimal("Eagle", utils.TilesList{
	{0, 0},
	{1, 0},
	{2, 1},
})

var Coyote = animal.NewAnimal("Coyote", utils.TilesList{
	{0, 0},
	{0, 1},
	{1, 2},
})

var Aatxe = animal.NewAnimal("Aatxe", utils.TilesList{
	{0, 0},
	{-1, 2},
})

var Pika = animal.NewAnimal("Pika", utils.TilesList{
	{0, 0},
	{0, 2},
	{1, 2},
})
