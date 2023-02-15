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
}

func (g *GameOverStage) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		fmt.Println("space")
		g.changeStage <- "game"
	}

	return nil
}
func (g *GameOverStage) Draw(screen *ebiten.Image) {
	faces := getOpentypeFaces()
	text.Draw(screen, "Game Over", faces[3], 100, 100, image.White)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
}

func (g *GameOverStage) Reset() {}
