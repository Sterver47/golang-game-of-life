package main

import (
	"fmt"
	"github.com/Sterver47/golang-game-of-life/internal/game"
	"log"
	"time"
)

func main() {
	g, err := game.LoadFieldFile("grid.txt")
	if err != nil {
		log.Fatal(err)
	}
	g.SeparatorCelChar = "-"

	i := 1
	for i <= 400 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Generation: %d\n", i)
		g.Evolve()
		i++
		g.Print()
	}
}
