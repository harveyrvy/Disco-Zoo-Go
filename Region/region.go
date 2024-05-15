package region

import (
	animal "github.com/discozoo/Animal"
	city "github.com/discozoo/Region/City"
	farm "github.com/discozoo/Region/Farm"
	iceage "github.com/discozoo/Region/IceAge"
	jungle "github.com/discozoo/Region/Jungle"
	jurassic "github.com/discozoo/Region/Jurassic"
	mars "github.com/discozoo/Region/Mars"
	moon "github.com/discozoo/Region/Moon"
	mountain "github.com/discozoo/Region/Mountain"
	northern "github.com/discozoo/Region/Northern"
	outback "github.com/discozoo/Region/Outback"
	polar "github.com/discozoo/Region/Polar"
	savannah "github.com/discozoo/Region/Savannah"
)

type Region struct {
	name    string
	animals []animal.Animal
}

func New(name string, animals []animal.Animal) Region {
	return Region{name, animals}
}

func GetAllRegions() []Region {
	return []Region{
		New("farm", []animal.Animal{farm.Rabbit, farm.Sheep, farm.Pig, farm.Cow, farm.Horse, farm.Unicorn, farm.Chicken}),
		New("outback", []animal.Animal{outback.Kangaroo, outback.Platypus, outback.Crocodile, outback.Koala, outback.Cockatoo, outback.Tiddalik, outback.Echidna}),
		New("savannah", []animal.Animal{savannah.Zebra, savannah.Hippo, savannah.Giraffe, savannah.Elephant, savannah.Lion, savannah.Gryphon, savannah.Rhino}),
		New("northern", []animal.Animal{northern.Bear, northern.Beaver, northern.Skunk, northern.Fox, northern.Moose, northern.Sasquatch, northern.Otter}),
		New("polar", []animal.Animal{polar.Penguin, polar.Seal, polar.Muskox, polar.PolarBear, polar.Walrus, polar.Yeti, polar.SnowyOwl}),
		New("jungle", []animal.Animal{jungle.Monkey, jungle.Gorilla, jungle.Toucan, jungle.Tiger, jungle.Panda, jungle.Phoenix, jungle.Lemur}),
		New("jurassic", []animal.Animal{jurassic.Diplodocus, jurassic.Raptor, jurassic.Triceratops, jurassic.Stegosaurus, jurassic.TRex, jurassic.Dragon, jurassic.Ankylosaurus}),
		New("iceage", []animal.Animal{iceage.DireWolf, iceage.GiantSloth, iceage.WoolyRhino, iceage.Mammoth, iceage.SabreTooth, iceage.Akhlut, iceage.YukonCamel}),
		New("city", []animal.Animal{city.Pigeon, city.Raccoon, city.Rat, city.Opossum, city.Squirrel, city.SewerTurtle, city.Opossum}),
		New("mountain", []animal.Animal{mountain.Elk, mountain.Goat, mountain.Coyote, mountain.Cougar, mountain.Eagle, mountain.Aatxe, mountain.Pika}),
		New("moon", []animal.Animal{moon.Moonkey, moon.LunarTick, moon.Tribble, moon.LunaMoth, moon.Moonicorn, moon.JadeRabbit, moon.Babmoon}),
		New("mars", []animal.Animal{mars.Marsmot, mars.Marsmoset, mars.Rock, mars.Rover, mars.Martian, mars.Marsmallow, mars.Marsten}),
	}
}

func GetRegionMap() map[string][]string {
	regions := GetAllRegions()
	m := make(map[string][]string)
	println(m)
	for _, r := range regions {
		m[r.name] = r.GetAnimalNames()
	}
	return m
}

func (r *Region) GetAnimals() []animal.Animal {
	return r.animals
}
func (r *Region) GetAnimalNames() []string {
	names := []string{}
	for _, a := range r.GetAnimals() {
		names = append(names, a.GetName())
	}
	return names
}

func (r *Region) GetRegionName() string {
	return r.name
}
