package jurassic

import (
	animal "github.com/discozoo/Animal"

	utils "github.com/discozoo/Utils"
)

// Coords are:
// {amt down, amt right}

var Diplodocus = animal.NewAnimal("Diplo", utils.TilesList{
	{0, 0},
	{1, 1},
	{1, 2},
	{2, 1},
})

var Stegosaurus = animal.NewAnimal("Stego", utils.TilesList{
	{0, 0},
	{-1, 1},
	{-1, 2},
	{0, 3},
})

var Raptor = animal.NewAnimal("Raptor", utils.TilesList{
	{0, 0},
	{0, 1},
	{1, 1},
	{2, 2},
})

var TRex = animal.NewAnimal("T-Rex", utils.TilesList{
	{0, 0},
	{2, 0},
	{2, 1},
})

var Triceratops = animal.NewAnimal("Tricera", utils.TilesList{
	{0, 0},
	{1, 2},
	{2, 0},
})

var Dragon = animal.NewAnimal("Dragon", utils.TilesList{
	{0, 0},
	{1, 2},
})

var Ankylosaurus = animal.NewAnimal("Ankylo", utils.TilesList{
	{0, 0},
	{0, 2},
	{-1, 2},
})
