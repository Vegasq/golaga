package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image"
)

type GameOverStage struct {
	changeStage chan string
	background  *Background
}

func (g *GameOverStage) Update() error {
	if g.background == nil {
		g.background = NewBackground("Space_BG_03")
	}

	g.background.Update()

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		fmt.Println("space")
		g.changeStage <- "game"
	}

	return nil
}
func (g *GameOverStage) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)

	faces := getOpentypeFaces()
	text.Draw(screen, "Game Over", faces[3], 100, 100, image.White)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
}

func (g *GameOverStage) Reset() {}
