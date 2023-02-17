package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image"
	"log"
	"time"
)

type TheEndStage struct {
	changeStage chan string
	background  *Background

	init     bool
	initDate time.Time
}

const waitInTheEndStage = time.Duration(3 * time.Second)

func (g *TheEndStage) Update() error {
	if g.init == false {
		g.background = NewBackground("Space_BG_03", 1)
		g.initDate = time.Now()

		g.init = true
	}

	if err := g.background.Update(); err != nil {
		log.Printf("error in background update: %v", err)
	}

	if time.Since(g.initDate) > waitInTheEndStage && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.init = false
		log.Println("change stage to prepare")
		g.changeStage <- "prepare"
	}

	return nil
}
func (g *TheEndStage) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)

	faces := getOpentypeFaces()

	_, h := screen.Size()
	text.Draw(screen, fmt.Sprintf("Congratulations!"), faces[4], 100, h/2, image.White)
	text.Draw(screen, fmt.Sprintf("You have completed your mission \nand destroyed all enemy ships.\n\nYou have saved the galaxy from invasion \nand earned yourself a medal of honor."), faces[3], 100, h/2+100, image.White)
	text.Draw(screen, "Press space to do it again", faces[3], 100, h/2+300, image.White)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
}

func (g *TheEndStage) Reset() {}
