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
const PlayerSpeed = 8

var AlienDissentSpeed = float64(2)

type GameStage struct {
	changeStage chan string

	player *Player
	aliens Aliens

	bullets           *Bullets
	lastBulletSpawned time.Time

	background *Background
}

func (g *GameStage) Update() error {

	if g.player == nil {
		g.player = NewPlayer()
		g.bullets = NewBullets(g.player)
		g.aliens = NewAliens()
		g.background = NewBackground("Space_BG_04")
	}

	g.player.Update()
	g.bullets.Update()
	g.background.Update()
	g.aliens.Update()

	if haveAliveAliens(g.aliens) == false {
		g.aliens = NewAliens()
	}

	bulletsAlienCollision(g)

	return nil
}

func (g *GameStage) Draw(screen *ebiten.Image) {
	if alienTouchedTheGround(screen, g.aliens) {
		g.changeStage <- "gameover"
	}

	g.background.Draw(screen)
	g.player.Draw(screen)
	g.bullets.Draw(screen)
	g.aliens.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
}

func (g *GameStage) Reset() {
	g.player = nil
	g.aliens = nil
	g.bullets = nil
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

func bulletsAlienCollision(g *GameStage) {
	for j, b := range g.bullets.bullets {
		if b == nil {
			continue
		}

		xB := b.pos.Element(0, 2)
		yB := b.pos.Element(1, 2)

		if yB < -1000 {
			g.bullets.bullets[j] = nil
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
				g.bullets.bullets[j] = nil
			}
		}
	}
}
