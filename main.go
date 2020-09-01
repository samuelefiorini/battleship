package main

import "fmt"

func main(){
	// Initialize a new fleet
	fleet := newFleet()
	fmt.Println(fleet)

	// Initialize a new grid
	grid := newGrid(10, 10)
	
	printGrid(grid)
}