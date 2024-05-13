package farm

import (
	animal "github.com/discozoo/Animal"

	utils "github.com/discozoo/Utils"
)

// Coords are:
// {amt down, amt right}

func GetAnimals() []animal.Animal {
	return []animal.Animal{Rabbit, Sheep, Pig, Cow, Horse, Unicorn, Chicken}
}
func GetAnimalNames() []string {
	names := []string{}
	for _, a := range []animal.Animal{Rabbit, Sheep, Pig, Cow, Horse, Unicorn, Chicken} {
		names = append(names, a.GetName())
	}
	return names
}

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
