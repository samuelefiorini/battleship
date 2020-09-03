package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Coordinate of the grid
type coordinate struct {
	row string
	col int
}

// Cell is the content of a grid
// an opponent cell can be either hit (H), ocean (O) or undefined (?)
// the player cell can be either ship OK (#), ship KO (X) or ocean (O)
type Cell struct {
	state string
	coord coordinate
}

// Grid definition
type Grid map[string][]Cell

// Capital characters
var asciiOffset rune = 65

// Grids are organized row-wise
func newGrid(nRows int, nCols int) Grid {
	grid := make(Grid)
	for i := 0; i < nRows; i++ {
		rowKey := string(asciiOffset + rune(i))
		row := make([]Cell, nCols)
		for j := 0; j < nCols; j++ {
			row[j] = Cell{"?", coordinate{rowKey, j}}
		}
		grid[rowKey] = row
	}
	return grid
}

// Randomly place input ships on an input grid
func (g Grid) populate(f Fleet) Grid {
	// Get number of rows and columns of the grid
	nRows := len(g)
	nCols := len(g["A"])
	fmt.Println(nRows, nCols)

	// Init random seed
	rand.Seed(time.Now().UnixNano())

	// Orient each ship either vertical (1) or horizontal (0)
	orientation := make([]int, len(f))
	for i := 0; i < len(f); i++ {
		orientation[i] = rand.Intn(len(f)) % 2
	}
	fmt.Println(orientation)

	// Fill the grid with ships
	for i, s := range f {
		fmt.Println(i, s)

		// Pick the random position of the first element (top-left) of the grid
		// then fill the rest of the grid
		var topLeft coordinate
		if false { //orientation[i] == 1 {
			// Vertical
			topLeft = coordinate{row: string(asciiOffset + rune(rand.Intn(nRows-s.size)+s.size)), col: rand.Intn(nCols)}
		} else {
			// Horizontal
			// rand.Intn(max - min) + min
			// we also have to exclude rows that do not have s.size + 2 free available spaces
			topLeft = coordinate{row: string(asciiOffset + rune(rand.Intn(nRows))), col: rand.Intn(nCols - s.size)}
			fmt.Println("TL: ", topLeft)
			// Fill grid with the rest of the ship
			for j := 0; j <= s.size; j++ {
				c := coordinate{row: topLeft.row, col: topLeft.col + j}
				fmt.Println(c)
				g[c.row][c.col] = Cell{string(s.kind[0]), c}
			}
		}
	}
	fmt.Println(g)
	return g
}

// Display the content of a grid
func printGrid(g Grid, verbose bool) {
	// Create a dummy fixed-lengh frame
	frameLength := 30
	_frame := make([]string, frameLength)
	for i := 0; i < frameLength; i++ {
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
	if verbose {
		fmt.Printf("Input grid size: (%v x %v)\n", nRows, nCols)
		fmt.Println(frame)
	}
	// Print columns index
	fmt.Printf("  ")
	for i := 0; i < nCols; i++ {
		fmt.Printf("%v ", i+1)
	}
	fmt.Printf("\n")
	for _, key := range keys {
		fmt.Printf("%v ", key)
		for _, c := range g[key] {
			fmt.Printf("%v ", c.state)
		}
		fmt.Printf("\n")
	}
	fmt.Println(frame)
}
