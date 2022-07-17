package main

import (
	"fmt"
	"github.com/Sterver47/golang-game-of-life/internal/field"
	"log"
	"time"
)

func main() {
	game, err := field.LoadFieldFile("grid.txt")
	if err != nil {
		log.Fatal(err)
	}
	//field.OVERFLOW = false

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
