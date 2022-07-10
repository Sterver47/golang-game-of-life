package field

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type grid [][]bool

var OVERFLOW = true

func LoadGrid(filename string) grid {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	s := strings.ReplaceAll(string(f), "\r", "")
	lines := strings.Split(s, "\n")
	height := len(lines)
	width := len(lines[0])

	field := make(grid, height)
	for y, line := range lines {
		if len(line) != width {
			log.Fatalf("[Error] Lines length mismatch: Line %d contains %d characters, expected %d.",
				y+1, len(line), width)
		}
		field[y] = make([]bool, width)
		for x, c := range line {
			if c == '-' {
				field[y][x] = false
			} else if c == 'x' {
				field[y][x] = true
			} else {
				log.Fatalf("[Error] Invalid character: %q at position %d:%d.", c, y+1, x+1)
			}
		}
	}

	return field
}

func (g grid) Print() {
	for y := range g {
		for _, live := range g[y] {
			if live {
				fmt.Printf("▓ ")
			} else {
				fmt.Printf("░ ")
			}
		}
		fmt.Printf("\n")
	}
}

func (g *grid) Evolve() {
	height, width := len(*g), len((*g)[1])
	newGrid := make(grid, height)
	for y := range newGrid {
		newGrid[y] = make([]bool, width)
	}

	for y := range *g {
		for x, live := range (*g)[y] {
			neighboursYX := generateNeighbors(y, x, height, width, OVERFLOW)
			liveNeighbors := 0
			for _, nyx := range neighboursYX {
				if (*g)[nyx[0]][nyx[1]] {
					liveNeighbors++
				}
			}

			var nextLive bool
			if (live && !(liveNeighbors < 2 || liveNeighbors > 3)) || (!live && liveNeighbors == 3) {
				nextLive = true
			} else {
				nextLive = false
			}

			newGrid[y][x] = nextLive
		}
	}
	*g = newGrid
}

func generateNeighbors(y, x, yMax, xMax int, of bool) [][2]int {
	neighbours := make([][2]int, 0, 8)
	for ny := -1; ny <= 1; ny++ {
		for nx := -1; nx <= 1; nx++ {
			if ny == 0 && nx == 0 {
				continue
			}
			yy := y + ny
			xx := x + nx

			if (!of) && (yy < 0 || yy > yMax-1 || xx < 0 || xx > xMax-1) {
				continue
			} else {
				if yy < 0 {
					yy = yMax + yy
				} else if yy > yMax-1 {
					yy = yy - yMax
				}
				if xx < 0 {
					xx = xMax + xx
				} else if xx > xMax-1 {
					xx = xx - xMax
				}
			}
			neighbours = append(neighbours, [2]int{yy, xx})
		}
	}
	return neighbours
}
