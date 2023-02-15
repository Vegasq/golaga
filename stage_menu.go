package main

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

	bg *ebiten.Image
}

func (g *MenuStage) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		fmt.Println("space")
		g.changeStage <- "game"
	}

	if g.bg == nil {
		img, _, err := ebitenutil.NewImageFromFile("art/Space-vertical-game-backgrounds/PNG/Space_BG_04/Space_BG_04.png")
		if err != nil {
			fmt.Println("Failed to load image: ", err)
		}
		g.bg = img
	}

	return nil
}
func (g *MenuStage) Draw(screen *ebiten.Image) {
	faces := getOpentypeFaces()
	geoM := ebiten.GeoM{}
	//geoM.Scale(0.3, 0.3)
	screen.DrawImage(g.bg, &ebiten.DrawImageOptions{
		GeoM: geoM,
	})
	text.Draw(screen, "Press SPACE to start", faces[3], 100, 100, image.White)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
}
func (g *MenuStage) Reset() {}
