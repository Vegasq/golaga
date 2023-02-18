package golaga

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/opentype"
	"image"
	"log"

	"golang.org/x/image/font"
)

var arcadeFonts map[int]font.Face

const arcadeFontBaseSize = 8

func getOpentypeFaces() map[int]font.Face {
	if arcadeFonts == nil {

		tt, err := opentype.Parse(fonts.PressStart2P_ttf)
		if err != nil {
			log.Fatal(err)
		}

		arcadeFonts = map[int]font.Face{}
		for i := 1; i <= 4; i++ {
			const dpi = 72
			arcadeFonts[i], err = opentype.NewFace(tt, &opentype.FaceOptions{
				Size:    float64(arcadeFontBaseSize * i),
				DPI:     dpi,
				Hinting: font.HintingFull,
			})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return arcadeFonts
}

type MenuStage struct {
	changeStage chan string

	background *Background
}

func (g *MenuStage) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.changeStage <- "prepare"
	}

	if g.background == nil {
		g.background = NewBackground("Space_BG_01", 1)
	}

	g.background.Update()

	return nil
}
func (g *MenuStage) Draw(screen *ebiten.Image) {
	faces := getOpentypeFaces()
	g.background.Draw(screen)

	_, h := screen.Size()
	text.Draw(screen, "Galaxy Under Attack", faces[4], 100, h/2, image.White)
	text.Draw(screen, "Press SPACE to start", faces[3], 100, (h/2)+100, image.White)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Version 0.1, FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
}
func (g *MenuStage) Reset() {}
