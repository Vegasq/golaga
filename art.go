package main

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/liujiawm/graphics-go/graphics"
	"golang.org/x/image/draw"
	"image"
	"image/png"
	_ "image/png"
	"math"
)

// Player ship animation

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Exhaust/Exhaust_1_1_000.png
var Player_0_0 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Exhaust/Exhaust_1_1_001.png
var Player_0_1 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Exhaust/Exhaust_1_1_002.png
var Player_0_2 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Exhaust/Exhaust_1_1_003.png
var Player_0_3 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Exhaust/Exhaust_1_1_004.png
var Player_0_4 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Exhaust/Exhaust_1_1_005.png
var Player_0_5 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Exhaust/Exhaust_1_1_006.png
var Player_0_6 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Exhaust/Exhaust_1_1_007.png
var Player_0_7 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Exhaust/Exhaust_1_1_008.png
var Player_0_8 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Exhaust/Exhaust_1_1_009.png
var Player_0_9 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Explosion/Explosion_1_000.png
var Player_exp_0_0 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Explosion/Explosion_1_001.png
var Player_exp_0_1 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Explosion/Explosion_1_002.png
var Player_exp_0_2 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Explosion/Explosion_1_003.png
var Player_exp_0_3 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Explosion/Explosion_1_004.png
var Player_exp_0_4 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Explosion/Explosion_1_005.png
var Player_exp_0_5 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Explosion/Explosion_1_006.png
var Player_exp_0_6 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Explosion/Explosion_1_007.png
var Player_exp_0_7 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_01/Explosion/Explosion_1_008.png
var Player_exp_0_8 []byte

// END OF Player ship animation

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_05_Explosion_000.png
var Alien_exp_0_0 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_05_Explosion_001.png
var Alien_exp_0_1 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_05_Explosion_002.png
var Alien_exp_0_2 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_05_Explosion_003.png
var Alien_exp_0_3 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_05_Explosion_004.png
var Alien_exp_0_4 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_05_Explosion_005.png
var Alien_exp_0_5 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_05_Explosion_006.png
var Alien_exp_0_6 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_05_Explosion_007.png
var Alien_exp_0_7 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_05_Explosion_008.png
var Alien_exp_0_8 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_5_1_000.png
var Alien_0_0 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_5_1_001.png
var Alien_0_1 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_5_1_002.png
var Alien_0_2 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_5_1_003.png
var Alien_0_3 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_5_1_004.png
var Alien_0_4 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_5_1_005.png
var Alien_0_5 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_5_1_006.png
var Alien_0_6 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_5_1_007.png
var Alien_0_7 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_5_1_008.png
var Alien_0_8 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_5_1_009.png
var Alien_0_9 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_4_2_000.png
var Alien_1_0 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_4_2_001.png
var Alien_1_1 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_4_2_002.png
var Alien_1_2 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_4_2_003.png
var Alien_1_3 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_4_2_004.png
var Alien_1_4 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_4_2_005.png
var Alien_1_5 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_4_2_006.png
var Alien_1_6 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_4_2_007.png
var Alien_1_7 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_4_2_008.png
var Alien_1_8 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_4_2_009.png
var Alien_1_9 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_04_Explosion_000.png
var Alien_exp_1_0 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_04_Explosion_001.png
var Alien_exp_1_1 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_04_Explosion_002.png
var Alien_exp_1_2 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_04_Explosion_003.png
var Alien_exp_1_3 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_04_Explosion_004.png
var Alien_exp_1_4 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_04_Explosion_005.png
var Alien_exp_1_5 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_04_Explosion_006.png
var Alien_exp_1_6 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_04_Explosion_007.png
var Alien_exp_1_7 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Explosion/Ship_04_Explosion_008.png
var Alien_exp_1_8 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_6_1_000.png
var Alien_2_0 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_6_1_001.png
var Alien_2_1 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_6_1_002.png
var Alien_2_2 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_6_1_003.png
var Alien_2_3 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_6_1_004.png
var Alien_2_4 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_6_1_005.png
var Alien_2_5 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_6_1_006.png
var Alien_2_6 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_6_1_007.png
var Alien_2_7 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_6_1_008.png
var Alien_2_8 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships_Sprites/Exhaust/Exhaust_6_1_009.png
var Alien_2_9 []byte

// End of alien animation

//go:embed art/Space-vertical-game-backgrounds/PNG/Space_BG_01/Space_BG_01.png
var Space_BG_01 []byte

//go:embed art/Space-vertical-game-backgrounds/PNG/Space_BG_02/Space_BG_02.png
var Space_BG_02 []byte

//go:embed art/Space-vertical-game-backgrounds/PNG/Space_BG_03/Space_BG_03.png
var Space_BG_03 []byte

//go:embed art/Space-vertical-game-backgrounds/PNG/Space_BG_04/Space_BG_04.png
var Space_BG_04 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships/Ship_01.png
var Ship_01 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships/Ship_02.png
var Ship_02 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships/Ship_03.png
var Ship_03 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships/Ship_04.png
var Ship_04 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships/Ship_05.png
var Ship_05 []byte

//go:embed art/Ufo-spaceship-game-sprites/PNG/Ships/Ship_06.png
var Ship_06 []byte

//go:embed art/Spaceship-2d-game-sprites/PNG/Ship_Effects/Laser_1_4.png
var Laser_1_4 []byte

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

	Art{"Fire_Shot_4_2", Fire_Shot_4_2, 20, 20},
	Art{"Laser_1_4", Laser_1_4, 20, 20},
}
var playerAnimationFrames = []Art{
	//player
	Art{"Player_0_0", Player_0_0, 120, 120},
	Art{"Player_0_1", Player_0_1, 120, 120},
	Art{"Player_0_2", Player_0_2, 120, 120},
	Art{"Player_0_3", Player_0_3, 120, 120},
	Art{"Player_0_4", Player_0_4, 120, 120},
	Art{"Player_0_5", Player_0_5, 120, 120},
	Art{"Player_0_6", Player_0_6, 120, 120},
	Art{"Player_0_7", Player_0_7, 120, 120},
	Art{"Player_0_8", Player_0_8, 120, 120},
	Art{"Player_0_9", Player_0_9, 120, 120},
}
var playerExplosionFramesData = []Art{
	Art{"Player_exp_0_0", Player_exp_0_0, 120, 120},
	Art{"Player_exp_0_1", Player_exp_0_1, 120, 120},
	Art{"Player_exp_0_2", Player_exp_0_2, 120, 120},
	Art{"Player_exp_0_3", Player_exp_0_3, 120, 120},
	Art{"Player_exp_0_4", Player_exp_0_4, 120, 120},
	Art{"Player_exp_0_5", Player_exp_0_5, 120, 120},
	Art{"Player_exp_0_6", Player_exp_0_6, 120, 120},
	Art{"Player_exp_0_7", Player_exp_0_7, 120, 120},
	Art{"Player_exp_0_8", Player_exp_0_8, 120, 120},
}

var alien0ShipFramesData = []Art{
	//alien
	Art{"Alien_0_0", Alien_0_0, 120, 120},
	Art{"Alien_0_1", Alien_0_1, 120, 120},
	Art{"Alien_0_2", Alien_0_2, 120, 120},
	Art{"Alien_0_3", Alien_0_3, 120, 120},
	Art{"Alien_0_4", Alien_0_4, 120, 120},
	Art{"Alien_0_5", Alien_0_5, 120, 120},
	Art{"Alien_0_6", Alien_0_6, 120, 120},
	Art{"Alien_0_7", Alien_0_7, 120, 120},
	Art{"Alien_0_8", Alien_0_8, 120, 120},
	Art{"Alien_0_9", Alien_0_9, 120, 120},
}
var alien0ExpFramesData = []Art{
	//alien
	Art{"Alien_exp_0_0", Alien_exp_0_0, 120, 120},
	Art{"Alien_exp_0_1", Alien_exp_0_1, 120, 120},
	Art{"Alien_exp_0_2", Alien_exp_0_2, 120, 120},
	Art{"Alien_exp_0_3", Alien_exp_0_3, 120, 120},
	Art{"Alien_exp_0_4", Alien_exp_0_4, 120, 120},
	Art{"Alien_exp_0_5", Alien_exp_0_5, 120, 120},
	Art{"Alien_exp_0_6", Alien_exp_0_6, 120, 120},
	Art{"Alien_exp_0_7", Alien_exp_0_7, 120, 120},
	Art{"Alien_exp_0_8", Alien_exp_0_8, 120, 120},
}

var alien1ShipFramesData = []Art{
	//alien
	Art{"Alien_1_0", Alien_1_0, 120, 120},
	Art{"Alien_1_1", Alien_1_1, 120, 120},
	Art{"Alien_1_2", Alien_1_2, 120, 120},
	Art{"Alien_1_3", Alien_1_3, 120, 120},
	Art{"Alien_1_4", Alien_1_4, 120, 120},
	Art{"Alien_1_5", Alien_1_5, 120, 120},
	Art{"Alien_1_6", Alien_1_6, 120, 120},
	Art{"Alien_1_7", Alien_1_7, 120, 120},
	Art{"Alien_1_8", Alien_1_8, 120, 120},
	Art{"Alien_1_9", Alien_1_9, 120, 120},
}

var alien1ExpFramesData = []Art{
	//alien
	Art{"Alien_exp_1_0", Alien_exp_1_0, 120, 120},
	Art{"Alien_exp_1_1", Alien_exp_1_1, 120, 120},
	Art{"Alien_exp_1_2", Alien_exp_1_2, 120, 120},
	Art{"Alien_exp_1_3", Alien_exp_1_3, 120, 120},
	Art{"Alien_exp_1_4", Alien_exp_1_4, 120, 120},
	Art{"Alien_exp_1_5", Alien_exp_1_5, 120, 120},
	Art{"Alien_exp_1_6", Alien_exp_1_6, 120, 120},
	Art{"Alien_exp_1_7", Alien_exp_1_7, 120, 120},
	Art{"Alien_exp_1_8", Alien_exp_1_8, 120, 120},
}

var alien2ShipFramesData = []Art{
	//alien
	Art{"Alien_2_0", Alien_2_0, 120, 120},
	Art{"Alien_2_1", Alien_2_1, 120, 120},
	Art{"Alien_2_2", Alien_2_2, 120, 120},
	Art{"Alien_2_3", Alien_2_3, 120, 120},
	Art{"Alien_2_4", Alien_2_4, 120, 120},
	Art{"Alien_2_5", Alien_2_5, 120, 120},
	Art{"Alien_2_6", Alien_2_6, 120, 120},
	Art{"Alien_2_7", Alien_2_7, 120, 120},
	Art{"Alien_2_8", Alien_2_8, 120, 120},
	Art{"Alien_2_9", Alien_2_9, 120, 120},
}

func loadAlienShipAnimationArt() {
	packs := [][]Art{alien2ShipFramesData, alien1ShipFramesData, alien0ShipFramesData, alien0ExpFramesData, alien1ExpFramesData}

	for _, pack := range packs {
		for _, art2 := range pack {
			src, _ := png.Decode(bytes.NewReader(art2.data))
			dst := image.NewRGBA(image.Rect(0, 0, art2.width, art2.height))
			draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

			dst2 := image.NewRGBA(image.Rect(0, 0, art2.width, art2.height))

			graphics.Rotate(dst2, dst, &graphics.RotateOptions{math.Pi})

			artCache[art2.name] = ebiten.NewImageFromImage(dst2)
		}
	}
}

func loadPlayerAnimationArt() {
	packs := [][]Art{playerExplosionFramesData, playerAnimationFrames}

	for _, pack := range packs {
		for _, art2 := range pack {
			src, _ := png.Decode(bytes.NewReader(art2.data))
			dst := image.NewRGBA(image.Rect(0, 0, art2.width, art2.height))
			draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

			artCache[art2.name] = ebiten.NewImageFromImage(dst)
		}
	}
}

func loadArt() {
	for _, art2 := range artToLoad2 {
		src, _ := png.Decode(bytes.NewReader(art2.data))
		dst := image.NewRGBA(image.Rect(0, 0, art2.width, art2.height))
		draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

		artCache[art2.name] = ebiten.NewImageFromImage(dst)
	}
	loadAlienShipAnimationArt()
	loadPlayerAnimationArt()
}
