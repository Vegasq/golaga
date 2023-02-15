package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"time"
)

type GameStage struct {
	changeStage chan string

	player *Player
	aliens []*Alien

	bullets           []*Bullet
	lastBulletSpawned time.Time
}

func (g *GameStage) Update() error {
	if g.player == nil {
		g.player = NewPlayer()
		g.aliens = NewAliens()
	}

	if haveAliveAliens(g.aliens) == false {
		g.aliens = NewAliens()
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.pos.Translate(PlayerSpeed, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.pos.Translate(-1*PlayerSpeed, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if time.Since(g.lastBulletSpawned) > TimeBetweenBullets {
			g.bullets = append(g.bullets, NewBullet(*g.player.pos))
			g.lastBulletSpawned = time.Now()
		}
	}

	for _, a := range g.aliens {
		a.Update()
	}

	for j, b := range g.bullets {
		if b == nil {
			continue
		}

		xB := b.pos.Element(0, 2)
		yB := b.pos.Element(1, 2)

		if yB < -1000 {
			g.bullets[j] = nil
			continue
		}

		for i, a := range g.aliens {
			if a == nil {
				continue
			}

			xA := a.pos.Element(0, 2)
			yA := a.pos.Element(1, 2)

			withinXAxis := xB > xA && xB < xA+a.w
			withinYAxis := yB > yA && yB < yA+a.h
			if withinXAxis && withinYAxis {
				g.aliens[i] = nil
				g.bullets[j] = nil
			}
		}
	}

	return nil
}
func (g *GameStage) Draw(screen *ebiten.Image) {
	if alienTouchedTheGround(g.aliens) {
		g.changeStage <- "gameover"
	}

	screen.DrawImage(g.player.img, &ebiten.DrawImageOptions{
		GeoM: *g.player.pos,
	})

	for _, b := range g.bullets {
		if b == nil {
			continue
		}
		b.pos.Translate(0, -1*BulletSpeed)
		screen.DrawImage(b.img, &ebiten.DrawImageOptions{GeoM: *b.pos})
	}

	for _, a := range g.aliens {
		if a == nil {
			continue
		}
		a.pos.Translate(a.wiggle, AlienDissentSpeed)
		screen.DrawImage(a.img, &ebiten.DrawImageOptions{GeoM: *a.pos})
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))

}

func (g *GameStage) Reset() {
	g.player = nil
	g.aliens = nil
	g.bullets = nil
}

type Player struct {
	img *ebiten.Image
	pos *ebiten.GeoM

	w, h float64
}

type Bullet struct {
	img *ebiten.Image
	pos *ebiten.GeoM
}

type Alien struct {
	img *ebiten.Image
	pos *ebiten.GeoM

	w, h float64

	wiggleDirection bool
	wiggle          float64
	wiggleCountdown int
}

func (a *Alien) Update() {
	if a == nil {
		return
	}
	if a.wiggleCountdown > 50 {
		a.wiggleCountdown = 0
		a.wiggleDirection = !a.wiggleDirection
	}

	a.wiggleCountdown += 1
	if a.wiggleDirection {
		a.wiggle = 1
	} else {
		a.wiggle = -1
	}

}

func NewAlien(x, y float64) *Alien {
	img2Color := color.NRGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}
	alien := &Alien{
		img: ebiten.NewImage(30, 30),
		pos: &ebiten.GeoM{},

		w: 30, h: 30,
	}
	alien.img.Fill(img2Color)
	alien.pos.Translate(x, y)
	alien.wiggleCountdown = 0

	return alien
}

func NewAliens() []*Alien {
	var aliens []*Alien

	var i float64
	var j float64
	for i = 100; i < 600; i += 50 {
		for j = 0; j < 150; j += 50 {
			alien := NewAlien(i, j)
			aliens = append(aliens, alien)
		}
	}

	AlienDissentSpeed = AlienDissentSpeed * 1.1

	return aliens
}

func NewBullet(playerPos ebiten.GeoM) *Bullet {
	b := ebiten.NewImage(10, 10)
	img2Color := color.NRGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}
	b.Fill(img2Color)

	bpos := playerPos
	return &Bullet{b, &bpos}
}

func NewPlayer() *Player {
	player := &Player{}
	player.img = ebiten.NewImage(100, 30)
	player.w = 100
	player.h = 30
	img2Color := color.NRGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}
	player.img.Fill(img2Color)

	player.pos = &ebiten.GeoM{}
	player.pos.Translate(100, 440)
	return player
}

const TimeBetweenBullets = time.Duration(time.Millisecond * 300)
const BulletSpeed = 10
const PlayerSpeed = 4

var AlienDissentSpeed = 1.1

func haveAliveAliens(aliens []*Alien) bool {
	for _, a := range aliens {
		if a != nil {
			return true
		}
	}
	return false
}

func alienTouchedTheGround(aliens []*Alien) bool {
	for _, a := range aliens {
		if a == nil {
			continue
		}

		y := a.pos.Element(1, 2)
		if y > 480 {
			return true
		}
	}
	return false

}