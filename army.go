package main

// Ship has name and size
type Ship struct {
	kind string
	size int
}

type Fleet []Ship

func newFleet() Fleet {
	// Number ships in the fleet
	fleetComposition := map[string][]int{
		"Carrier":    {5, 1},
		"Battleship": {4, 2},
		"Cruiser":    {3, 3},
		"Submarine":  {3, 4},
		"Destroyer":  {2, 5},
	}

	var fleet Fleet
	for n, s := range fleetComposition {
		for i := 0; i < s[1]; i++ {
			fleet = append(fleet, Ship{n, s[0]})
		}
	}

	return fleet
}
