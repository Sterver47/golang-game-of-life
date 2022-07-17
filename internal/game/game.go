package game

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type grid [][]bool

type Game struct {
	Field            *grid
	Overflow         bool
	LiveCelChar      string
	DeadCelChar      string
	SeparatorCelChar string
}

func LoadFieldFile(filename string) (*Game, error) {
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
			return nil, fmt.Errorf("lines length mismatch: Line %d contains %d characters, expected %d",
				y+1, len(line), width)
		}
		field[y] = make([]bool, width)
		for x, c := range line {
			if c == '-' {
				field[y][x] = false
			} else if c == 'x' {
				field[y][x] = true
			} else {
				return nil, fmt.Errorf("invalid character: %q at position %d:%d", c, y+1, x+1)
			}
		}
	}

	return &Game{
		Field:            &field,
		Overflow:         true,
		LiveCelChar:      "▓",
		DeadCelChar:      "░",
		SeparatorCelChar: " ",
	}, nil
}

func (g Game) Print() {
	for y := range *g.Field {
		for _, live := range (*g.Field)[y] {
			if live {
				fmt.Print(g.LiveCelChar + g.SeparatorCelChar)
			} else {
				fmt.Print(g.DeadCelChar + g.SeparatorCelChar)
			}
		}
		fmt.Println()
	}
}

func (g *Game) Evolve() {
	height, width := len(*g.Field), len((*g.Field)[1])
	newGrid := make(grid, height)
	for y := range newGrid {
		newGrid[y] = make([]bool, width)
	}

	for y := range *g.Field {
		for x, live := range (*g.Field)[y] {
			neighboursYX := generateNeighbors(y, x, height, width, g.Overflow)
			liveNeighbors := 0
			for _, nyx := range *neighboursYX {
				if (*g.Field)[nyx[0]][nyx[1]] {
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
	g.Field = &newGrid
}

func generateNeighbors(y, x, yMax, xMax int, of bool) *[][2]int {
	var neighbours [][2]int
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
	return &neighbours
}
