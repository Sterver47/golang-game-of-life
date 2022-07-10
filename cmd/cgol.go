package main

import (
	"cgol/internal/field"
	"fmt"
	"time"
)

func main() {
	game := field.LoadGrid("..\\..\\grid.txt")

	i := 1
	for i < 400 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Generation: %d\n", i)
		game.Evolve()
		i++
		game.Print()
	}
}
