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
	aliens Aliens

	bullets           *Bullets
	lastBulletSpawned time.Time

	background *Background
}

func (g *GameStage) Update() error {
	if g.background == nil {
		g.background = NewBackground("Space_BG_04")
	}

	if g.bg == nil {
		g.bg = artCache["Space_BG_04"]
	}
	if g.player == nil {
		g.player = NewPlayer()
		g.aliens = NewAliens()
	}

	g.player.Update()
	if g.bullets == nil {
		g.bullets = NewBullets(g.player)
	}
	g.bullets.Update()
	g.background.Update()

	if haveAliveAliens(g.aliens) == false {
		g.aliens = NewAliens()
	}

	//if ebiten.IsKeyPressed(ebiten.KeySpace) {
	//	if time.Since(g.lastBulletSpawned) > TimeBetweenBullets {
	//		g.bullets = append(g.bullets, NewBullet(*g.player.pos))
	//		g.lastBulletSpawned = time.Now()
	//	}
	//}

	g.aliens.Update()

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

	return nil
}

func (g *GameStage) Draw(screen *ebiten.Image) {
	if alienTouchedTheGround(screen, g.aliens) {
		g.changeStage <- "gameover"
	}

	g.background.Draw(screen)
	g.player.Draw(screen)
	g.bullets.Draw(screen)

	//for _, b := range g.bullets.bullets {
	//	if b == nil {
	//		continue
	//	}
	//	b.pos.Translate(0, -1*BulletSpeed)
	//	screen.DrawImage(b.img, &ebiten.DrawImageOptions{GeoM: *b.pos})
	//}

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
