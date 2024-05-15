package moon

import (
	animal "github.com/discozoo/Animal"

	utils "github.com/discozoo/Utils"
)

// Coords are:
// {amt down, amt right}

var Moonkey = animal.NewAnimal("Moonkey", utils.TilesList{
	{0, 0},
	{1, 0},
	{1, 2},
	{2, 2},
})

var LunarTick = animal.NewAnimal("LunarTic", utils.TilesList{
	{0, 0},
	{2, 0},
	{3, -1},
	{3, 1},
})

var Tribble = animal.NewAnimal("Tribble", utils.TilesList{
	{0, 0},
	{1, 0},
	{1, -1},
	{1, 1},
})

var Moonicorn = animal.NewAnimal("Mooncorn", utils.TilesList{
	{0, 0},
	{1, 0},
	{1, 1},
})

var LunaMoth = animal.NewAnimal("LunaMoth", utils.TilesList{
	{0, 0},
	{0, 2},
	{2, 1},
})

var JadeRabbit = animal.NewAnimal("JadeRab", utils.TilesList{
	{0, 0},
	{2, 1},
})

var Babmoon = animal.NewAnimal("Babmoon", utils.TilesList{
	{0, 0},
	{1, 1},
	{2, -1},
})
