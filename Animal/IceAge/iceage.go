package iceage

import (
	animal "github.com/discozoo/Animal"

	utils "github.com/discozoo/Utils"
)

// Coords are:
// {amt down, amt right}

var WoolyRhino = animal.NewAnimal("WlyRhino", utils.TilesList{
	{0, 0},
	{1, 1},
	{-1, 2},
	{0, 3},
})

var GiantSloth = animal.NewAnimal("Sloth", utils.TilesList{
	{0, 0},
	{2, 0},
	{1, 2},
	{2, 2},
})

var DireWolf = animal.NewAnimal("DireWolf", utils.TilesList{
	{0, 0},
	{-1, 1},
	{1, 1},
	{0, 3},
})

var SabreTooth = animal.NewAnimal("SabTooth", utils.TilesList{
	{0, 0},
	{2, 1},
	{1, 2},
})

var Mammoth = animal.NewAnimal("Mammoth", utils.TilesList{
	{0, 0},
	{-1, 1},
	{1, 2},
})

var Akhlut = animal.NewAnimal("Akhlut", utils.TilesList{
	{0, 0},
	{-1, 2},
	{1, 2},
})

var YukonCamel = animal.NewAnimal("YukCamel", utils.TilesList{
	{0, 0},
	{-1, 2},
	{1, 3},
})
