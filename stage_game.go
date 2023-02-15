package main

import (
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"time"
)

const TimeBetweenBullets = time.Duration(time.Millisecond * 300)
const BulletSpeed = 10
const PlayerSpeed = 4

var AlienDissentSpeed = float64(1)

type GameStage struct {
	changeStage chan string

	bg *ebiten.Image

	player *Player
	aliens []*Alien

	bullets           []*Bullet
	lastBulletSpawned time.Time

	background *Background
}

func (g *GameStage) Update() error {
	if g.background == nil {
		g.background = NewBackground("Space_BG_04")
	}

	g.background.Update()

	if g.player == nil {
		g.player = NewPlayer()
		g.aliens = NewAliens()
	}

	if g.bg == nil {
		g.bg = artCache["Space_BG_04"]
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
	if alienTouchedTheGround(screen, g.aliens) {
		g.changeStage <- "gameover"
	}

	g.background.Draw(screen)

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
		a.pos.Translate(a.wiggle, float64(AlienDissentSpeed))
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
	w, h := artCache["Ship_06"].Size()
	alien := &Alien{
		img: artCache["Ship_06"],
		pos: &ebiten.GeoM{},

		w: float64(w), h: float64(h),
	}
	alien.pos.Translate(x, y)
	//alien.pos.Scale(0.125, 0.125)
	alien.wiggleCountdown = 0

	return alien
}

func NewAliens() []*Alien {
	var aliens []*Alien

	var i float64
	var j float64
	for i = 100; i < 1000; i += 200 {
		for j = 0; j < 150; j += 120 {
			alien := NewAlien(i, j)
			aliens = append(aliens, alien)
		}
	}

	AlienDissentSpeed = float64(AlienDissentSpeed) * 1.1

	return aliens
}

func NewBullet(playerPos ebiten.GeoM) *Bullet {
	bpos := ebiten.GeoM{}
	bpos.Translate(playerPos.Element(0, 2)+47, playerPos.Element(1, 2)-20)
	return &Bullet{artCache["Fire_Shot_4_2"], &bpos}
}

func NewPlayer() *Player {
	player := &Player{}

	player.img = artCache["Ship_LVL_1"]

	w, h := player.img.Size()

	player.w = float64(w)
	player.h = float64(h)

	player.pos = &ebiten.GeoM{}
	player.pos.Translate(100, 1720)
	return player
}

func haveAliveAliens(aliens []*Alien) bool {
	for _, a := range aliens {
		if a != nil {
			return true
		}
	}
	return false
}

func alienTouchedTheGround(screen *ebiten.Image, aliens []*Alien) bool {
	_, h := screen.Size()
	for _, a := range aliens {
		if a == nil {
			continue
		}

		y := int(a.pos.Element(1, 2))
		if y > h {
			return true
		}
	}
	return false
}
