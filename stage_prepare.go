package golaga

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image"
	"log"
	"time"
)

type PrepareStage struct {
	changeStage chan string
	background  *Background

	init     bool
	initDate time.Time
}

const waitInPrepareStage = time.Duration(1 * time.Second)

func (g *PrepareStage) Update() error {
	if g.init == false {
		g.background = NewBackground("Space_BG_03", 1)
		g.initDate = time.Now()

		g.init = true
	}

	if err := g.background.Update(); err != nil {
		log.Printf("error in background update: %v", err)
	}

	if time.Since(g.initDate) > waitInPrepareStage && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.init = false
		log.Println("change stage to game")
		g.changeStage <- "game"
	}

	return nil
}
func (g *PrepareStage) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)

	faces := getOpentypeFaces()

	_, h := screen.Size()
	text.Draw(screen, fmt.Sprintf("Stage %d", MissionIndex+1), faces[4], 100, h/2, image.White)
	text.Draw(screen, "Press space to start", faces[3], 100, h/2+100, image.White)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
}

func (g *PrepareStage) Reset() {}
