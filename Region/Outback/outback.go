package outback

import (
	animal "github.com/discozoo/Animal"

	utils "github.com/discozoo/Utils"
)

// Coords are:
// {amt down, amt right}

var Kangaroo = animal.NewAnimal("Kangaroo", utils.TilesList{
	{0, 0},
	{1, 1},
	{2, 2},
	{3, 3},
})

var Platypus = animal.NewAnimal("Platypus", utils.TilesList{
	{0, 0},
	{0, 1},
	{1, 1},
	{1, 2},
})

var Crocodile = animal.NewAnimal("Croc", utils.TilesList{
	{0, 0},
	{0, 1},
	{0, 2},
	{0, 3},
})

var Koala = animal.NewAnimal("Koala", utils.TilesList{
	{0, 0},
	{0, 1},
	{1, 1},
})

var Cockatoo = animal.NewAnimal("Cockatoo", utils.TilesList{
	{0, 0},
	{1, 1},
	{2, 1},
})

var Tiddalik = animal.NewAnimal("Tiddalik", utils.TilesList{
	{0, 0},
	{-1, 1},
	{0, 2},
})

var Echidna = animal.NewAnimal("Echidna", utils.TilesList{
	{0, 0},
	{0, 1},
	{-1, 2},
})
