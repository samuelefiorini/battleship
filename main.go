package main

import "fmt"

func main() {
	// Initialize a new fleet
	fleet := newFleet()
	fmt.Println(fleet)

	// Initialize a new grid
	grid := newGrid(10, 10)
	fmt.Println("New empty grid:")
	printGrid(grid, false)

	// Randomly populate the grid with ships
	grid.populate(fleet[:2])
	fmt.Println("Populated grid:")
	printGrid(grid, false)
}
