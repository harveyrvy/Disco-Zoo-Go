package northern

import (
	animal "github.com/discozoo/Animal"

	utils "github.com/discozoo/Utils"
)

// Coords are:
// {amt down, amt right}

var Bear = animal.NewAnimal("Bear", utils.TilesList{
	{0, 0},
	{0, 1},
	{1, 1},
	{2, 1},
})

var Skunk = animal.NewAnimal("Skunk", utils.TilesList{
	{0, 0},
	{0, 1},
	{-1, 1},
	{-1, 2},
})

var Beaver = animal.NewAnimal("Beaver", utils.TilesList{
	{0, 0},
	{0, 1},
	{-1, 2},
	{1, 2},
})

var Moose = animal.NewAnimal("Moose", utils.TilesList{
	{0, 0},
	{1, 1},
	{0, 2},
})

var Fox = animal.NewAnimal("Fox", utils.TilesList{
	{0, 0},
	{0, 1},
	{1, 2},
})

var Sasquatch = animal.NewAnimal("Sasqtch", utils.TilesList{
	{0, 0},
	{1, 0},
})

var Otter = animal.NewAnimal("Otter", utils.TilesList{
	{0, 0},
	{1, 0},
	{1, 1},
})
