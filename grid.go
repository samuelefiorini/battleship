package main	

import (
	"fmt"
	"strings"
	"sort"
)

// Coordinate of the grid
type coordinate struct{
	row string
	col int
}

// Cell is the content of a grid
// a cell can be either hit (H), ocean (O)
// or undefined (?)
type Cell struct{
	state string
	coord coordinate
}

// Grids are organized row-wise
func newGrid(nRows int, nCols int) map[string]map[int]Cell{
	asciiOffset := 65 // Capital characters
	grid := make(map[string]map[int]Cell)
	
	for i := 0; i<nRows; i++ {
		rowKey := string(asciiOffset + i)
		row := make(map[int]Cell)
		for j :=0; j<nCols; j++ {
			 row[j+1] = Cell{"?", coordinate{rowKey, j+1}}
		}
		grid[rowKey] = row
	}

	return grid
}

func printGrid(g map[string]map[int]Cell) {
	// Create a dummy fixed-lengh frame
	frameLength := 30
	_frame := make([]string, frameLength)
	for i := 0; i<frameLength; i++ {
		_frame[i] = "-"
	}
	frame := strings.Join(_frame, "")
	// Get number of rows and columns of the grid
	nRows := len(g)
	nCols := len(g["A"])
	// Get sorted map keys
	keys := make([]string, 0, nRows)
	for k := range g {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// Start printing
	fmt.Println(frame)
	fmt.Printf("Input grid size: (%v x %v)\n", nRows, nCols)
	fmt.Println(frame)
	// Print columns index
	fmt.Printf("  ")
	for i := 0; i<nCols; i++ {
		fmt.Printf("%v ", i+1)
	}
	fmt.Printf("\n")
	for _, key := range keys{
		fmt.Printf("%v ", key)
		for _, c := range g[key] {
			fmt.Printf("%v ", c.state)
		}
		fmt.Printf("\n")
	}
	fmt.Println(frame)
}