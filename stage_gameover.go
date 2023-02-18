package golaga

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image"
	"time"
)

type GameOverStage struct {
	changeStage chan string
	background  *Background

	init     bool
	initDate time.Time
}

const waitInGameoverStage = time.Duration(1 * time.Second)

func (g *GameOverStage) Update() error {
	if g.init == false {
		g.init = true
		g.initDate = time.Now()
		g.background = NewBackground("Space_BG_03", 1)
	}

	g.background.Update()

	if time.Since(g.initDate) > waitInGameoverStage && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.init = false
		g.changeStage <- "game"
	}

	return nil
}
func (g *GameOverStage) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)

	faces := getOpentypeFaces()

	_, h := screen.Size()
	text.Draw(screen, "Game Over", faces[4], 100, h/2, image.White)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
}

func (g *GameOverStage) Reset() {}
