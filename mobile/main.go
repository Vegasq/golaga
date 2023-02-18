package main

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"
	"github.com/vegasq/golaga"
)

func main() {
	game := golaga.PrepareGame()
	mobile.SetGame(game)
}
func Dummy() {}
