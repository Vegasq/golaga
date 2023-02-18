package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vegasq/golaga"
	"log"
)

func main() {
	game := golaga.PrepareGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
