package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
)

const (
	GREEN       = "\033[32m"
	RED         = "\033[31m"
	GRID_WIDTH  = 30
	GRID_HEIGHT = 20
	SQUARE_CHAR = GREEN + "■"
	APPLE_CHAR  = RED + "■"
	EMPTY_CHAR  = " "
	FRAME_RATE  = 250 * time.Millisecond
)

type Grid [][]string

type Position struct {
	x, y int
}

type Player struct {
	body []Position // Store snake segments
}

type Direction struct {
	dx, dy int
}

type Apple struct {
	x, y int
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func newGrid() Grid {
	grid := make(Grid, GRID_HEIGHT)
	for i := range grid {
		grid[i] = make([]string, GRID_WIDTH)
		for j := range grid[i] {
			grid[i][j] = EMPTY_CHAR
		}
	}
	return grid
}

func drawGrid(grid Grid, player Player, apple Apple) {
	clearScreen()

	// Create a fresh grid for drawing
	tempGrid := newGrid()

	// Draw snake body
	for _, segment := range player.body {
		if segment.y >= 0 && segment.y < GRID_HEIGHT && segment.x >= 0 && segment.x < GRID_WIDTH {
			tempGrid[segment.y][segment.x] = SQUARE_CHAR
		}
	}

	// Draw apple
	if apple.y >= 0 && apple.y < GRID_HEIGHT && apple.x >= 0 && apple.x < GRID_WIDTH {
		tempGrid[apple.y][apple.x] = APPLE_CHAR
	}

	for _, row := range tempGrid {
		for _, char := range row {
			fmt.Printf("%s", char)
		}
		fmt.Println()
	}
}

func spawnApple(player Player) Apple {
	for {
		apple := Apple{x: rand.Intn(GRID_WIDTH), y: rand.Intn(GRID_HEIGHT)}
		// Ensure apple doesn't spawn on snake
		collision := false
		for _, segment := range player.body {
			if apple.x == segment.x && apple.y == segment.y {
				collision = true
				break
			}
		}
		if !collision {
			return apple
		}
	}
}

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	grid := newGrid()
	// Initialize player with one segment at the center
	player := Player{body: []Position{{x: GRID_WIDTH / 2, y: GRID_HEIGHT / 2}}}
	apple := spawnApple(player)
	dir := Direction{-1, 0}

	keyEvents, err := keyboard.GetKeys(10)
	if err != nil {
		fmt.Println("Error getting keys:", err)
		return
	}

	for {
		select {
		case event := <-keyEvents:
			if event.Err != nil {
				fmt.Println("Keyboard event error:", event.Err)
				return
			}
			switch event.Rune {
			case 'w', 'W':
				if dir.dy != 1 { // Prevent moving directly opposite
					dir = Direction{0, -1}
				}
			case 's', 'S':
				if dir.dy != -1 {
					dir = Direction{0, 1}
				}
			case 'a', 'A':
				if dir.dx != 1 {
					dir = Direction{-1, 0}
				}
			case 'd', 'D':
				if dir.dx != -1 {
					dir = Direction{1, 0}
				}
			case 'q', 'Q':
				return // Quit game
			}

		default:
			// Move snake
			head := player.body[0]
			newHead := Position{x: head.x + dir.dx, y: head.y + dir.dy}

			// Check boundaries
			if newHead.x < 0 {
				newHead.x = 0
			} else if newHead.x >= GRID_WIDTH {
				newHead.x = GRID_WIDTH - 1
			}
			if newHead.y < 0 {
				newHead.y = 0
			} else if newHead.y >= GRID_HEIGHT {
				newHead.y = GRID_HEIGHT - 1
			}

			// Check collision with self
			for _, segment := range player.body {
				if newHead.x == segment.x && newHead.y == segment.y {
					fmt.Println("Game Over: Collided with self!")
					return
				}
			}

			// Move snake by adding new head
			newBody := []Position{newHead}
			grow := false

			// Check for apple collision
			if newHead.x == apple.x && newHead.y == apple.y {
				grow = true
				apple = spawnApple(player)
			}

			// Add previous segments (excluding last if not growing)
			if grow {
				newBody = append(newBody, player.body...)
			} else {
				newBody = append(newBody, player.body[:len(player.body)-1]...)
			}
			player.body = newBody

			drawGrid(grid, player, apple)
			time.Sleep(FRAME_RATE)
		}
	}
}
