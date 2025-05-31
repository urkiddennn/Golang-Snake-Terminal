package main

import (
	"fmt"
	"time"
)

// In this Video I will screate a terminal snake game(I hope I can)

// print the square on agiven tile of the grid
const (
	GRID_WIDTH  = 30
	GRID_HEIGHT = 30
	SQUARE_CHAR = "■"
	EMPTY_CHAR  = " "
	FRAME_RATE  = 50 * time.Millisecond
)

// it's where we put our GRIDs
type Grid [][]string

// player x and Y axis
type Player struct {
	x, y int
}

// simple screen, it is short and readable but not scalable I think
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func newGrid() Grid { // return a GRid[][] refresh the location
	grid := make(Grid, GRID_HEIGHT)

	for i := range grid {
		//
		grid[i] = make([]string, GRID_WIDTH)
		for j := range grid[i] {
			// making that part of the grid to an empty space
			grid[i][j] = EMPTY_CHAR
		}
	}
	return grid
}

// draw the grid on the terminal
func drawGrid(grid Grid, player Player) {
	clearScreen()

	if player.y >= 0 && player.y < GRID_HEIGHT && player.x >= 0 && player.x < GRID_WIDTH {
		grid[player.y][player.x] = SQUARE_CHAR // set that position on that grid to be the square
	}

	for _, row := range grid {
		for _, char := range row {
			fmt.Printf("%s", char)
		}
		fmt.Println()
	}

	// complete opposite
	if player.y >= 0 && player.y < GRID_HEIGHT && player.x >= 0 && player.x < GRID_WIDTH {
		grid[player.y][player.x] = EMPTY_CHAR // set that position on that grid to be empty
	}
}

func main() {
	grid := newGrid()
	// setting the player at the center of ther grid
	player := Player{x: GRID_WIDTH / 2, y: GRID_HEIGHT / 2}
	for {
		clearScreen()

		drawGrid(grid, player)
		time.Sleep(FRAME_RATE)
	}
}
