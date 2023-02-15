package main

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/draw"
	"image"
	"image/png"
	_ "image/png"
)

//go:embed art/Space-vertical-game-backgrounds/PNG/Space_BG_01/Space_BG_01.png
var Space_BG_01 []byte

//go:embed art/Space-vertical-game-backgrounds/PNG/Space_BG_02/Space_BG_02.png
var Space_BG_02 []byte

//go:embed art/Space-vertical-game-backgrounds/PNG/Space_BG_03/Space_BG_03.png
var Space_BG_03 []byte

//go:embed art/Space-vertical-game-backgrounds/PNG/Space_BG_04/Space_BG_04.png
var Space_BG_04 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships/Ship_06.png
var Ship_06 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_Effects/Fire_Shot_4_2.png
var Fire_Shot_4_2 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Ship_LVL_1.png
var Ship_LVL_1 []byte

var artCache = map[string]*ebiten.Image{}

type Art struct {
	name   string
	data   []byte
	width  int
	height int
}

var artToLoad2 = []Art{
	Art{"Space_BG_01", Space_BG_01, 1080, 1920},
	Art{"Space_BG_02", Space_BG_02, 1080, 1920},
	Art{"Space_BG_03", Space_BG_03, 1080, 1920},
	Art{"Space_BG_04", Space_BG_04, 1080, 1920},
	Art{"Ship_06", Ship_06, 100, 100},
	Art{"Fire_Shot_4_2", Fire_Shot_4_2, 20, 20},
	Art{"Ship_LVL_1", Ship_LVL_1, 120, 120},
}

func loadArt() {
	for _, art2 := range artToLoad2 {

		src, _ := png.Decode(bytes.NewReader(art2.data))

		//dst := image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X/2, src.Bounds().Max.Y/2))
		dst := image.NewRGBA(image.Rect(0, 0, art2.width, art2.height))
		draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

		artCache[art2.name] = ebiten.NewImageFromImage(dst)
	}
	//
	//for _, art := range artToLoad {
	//	img, _, err := ebitenutil.NewImageFromFile(art)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	name := path.Base(art)
	//	artCache[name] = img
	//}
}
