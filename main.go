package main

import (
	"fmt"
	"log"
	"time"

	"github.com/eiannone/keyboard"
)

// In this Video I will screate a terminal snake game(I hope I can)

// print the square on agiven tile of the grid
const (
	GREEN       = "\033[32m"
	GRID_WIDTH  = 30
	GRID_HEIGHT = 20
	SQUARE_CHAR = GREEN + "■"
	EMPTY_CHAR  = " "
	FRAME_RATE  = 30 * time.Millisecond
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
	dx := 0
	// dy := 0
	// initialize the Keyboard
	err := keyboard.Open()
	if err != nil {
		// should we return or print? just use log
		log.Fatal(err)
	}
	defer keyboard.Close()

	grid := newGrid()
	// setting the player at the center of ther grid
	player := Player{x: GRID_WIDTH / 2, y: GRID_HEIGHT / 2}

	keyEvents, err := keyboard.GetKeys(10)
	if err != nil {
		fmt.Println("Error get keys", err)
		return
	}

	for {
		select {
		case event := <-keyEvents:
			if event.Err != nil {
				fmt.Println("Keyboard event error: ", event.Err)
				return
			}
			if event.Rune == 'w' || event.Rune == 'W' {
				player.y--

				// Rune is the key that is pressed
			}
			if event.Rune == 's' || event.Rune == 'S' {
				player.y++
			}
			if event.Rune == 'a' || event.Rune == 'A' {
				player.x--
			}
			if event.Rune == 'd' || event.Rune == 'D' {
				player.x++
			}
		default:
			// no key pressed
			dx += 1

		}

		// collision detectection of the bound edge
		if player.y < 0 {
			player.y = 0 // bound up
		} else if player.y >= GRID_HEIGHT {
			player.y = GRID_HEIGHT - 1 // bound bottom
		} else if player.x < 0 {
			player.x = 0
		} else if player.x >= GRID_WIDTH {
			player.x = GRID_WIDTH - 1
		}

		drawGrid(grid, player)
		fmt.Printf("%v", dx)
		time.Sleep(FRAME_RATE)
	}
}
