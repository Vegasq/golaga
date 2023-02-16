package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image"
	"time"
)

type PrepareStage struct {
	changeStage chan string
	background  *Background

	stage    int
	init     bool
	initDate time.Time
}

const waitInPrepareStage = time.Duration(1 * time.Second)

func (g *PrepareStage) Update() error {
	if g.init == false {
		g.background = NewBackground("Space_BG_03", 1)
		g.initDate = time.Now()
		g.stage++

		g.init = true
	}

	g.background.Update()

	if time.Since(g.initDate) > waitInPrepareStage && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.init = false
		g.changeStage <- "game"
	}

	return nil
}
func (g *PrepareStage) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)

	faces := getOpentypeFaces()

	_, h := screen.Size()
	text.Draw(screen, fmt.Sprintf("Stage %d", g.stage), faces[4], 100, h/2, image.White)
	text.Draw(screen, "Press space to start", faces[3], 100, h/2+100, image.White)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
}

func (g *PrepareStage) Reset() {}
