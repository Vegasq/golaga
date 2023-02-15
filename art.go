package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"path"
)

var artCache = map[string]*ebiten.Image{}

type Art struct {
	path   string
	width  int
	height int
}

var artToLoad = []string{
	"art/Space-vertical-game-backgrounds/PNG/Space_BG_04/Space_BG_04.png",
	"art/Ufo-spaceship-game-sprites/PNG/Ships/Ship_06.png",
	"art/Spaceship-2d-game-sprites/PNG/Ship_Effects/Fire_Shot_4_2.png",
	"art/Spaceship-2d-game-sprites/PNG/Ship_01/Ship_LVL_1.png",
}

func loadArt() {
	for _, art := range artToLoad {
		img, _, err := ebitenutil.NewImageFromFile(art)
		if err != nil {
			log.Println(err)
		}
		name := path.Base(art)
		artCache[name] = img
	}
}
