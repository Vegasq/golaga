package main

import (
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image"
	"math/rand"
	"time"
)

const TimeBetweenBullets = time.Duration(time.Millisecond * 300)
const TimeBetweenAlienBullets = time.Duration(time.Millisecond * 2000)
const BulletSpeed = 10
const PlayerSpeed = 8

var AlienDissentSpeed = float64(1.1)

type GameStage struct {
	changeStage chan string

	player *Player
	aliens *Aliens

	bullets           *Bullets
	alienBullets      *AliensBullets
	lastBulletSpawned time.Time

	background *Background
}

func (g *GameStage) Update() error {

	if g.player == nil {
		g.player = NewPlayer()
		g.bullets = NewBullets(g.player)
		g.aliens = NewAliens()
		backgrounds := []string{"Space_BG_01", "Space_BG_02", "Space_BG_03", "Space_BG_04"}
		i := rand.Intn(len(backgrounds))
		g.background = NewBackground(backgrounds[i], 5)
		g.alienBullets = NewAliensBullets(g.aliens)
	}

	g.player.Update()
	g.bullets.Update()
	g.background.Update()
	g.aliens.Update()
	g.alienBullets.Update()

	aliensShoot(g.aliens, g.alienBullets)

	if haveAliveAliens(g.aliens) == false {
		g.changeStage <- "prepare"
	}

	bulletsAlienCollision(g)

	if bulletsPlayerCollision(g) {
		g.player.animation.Explode()
	}

	if g.player.alive == false {
		g.changeStage <- "gameover"
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
	g.aliens.Draw(screen)
	g.alienBullets.Draw(screen)

	faces := getOpentypeFaces()
	text.Draw(screen, "Move: Arrows", faces[3], 50, 100, image.White)
	text.Draw(screen, "Shoot: Spacebar", faces[3], 50, 150, image.White)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
}

func (g *GameStage) Reset() {
	g.player = nil
	g.aliens = nil
	g.bullets = nil
	g.background = nil
}

func haveAliveAliens(aliens *Aliens) bool {
	for _, a := range aliens.GetAliens() {
		if a != nil {
			return true
		}
	}
	return false
}

func alienTouchedTheGround(screen *ebiten.Image, aliens *Aliens) bool {
	_, h := screen.Size()
	if aliens == nil {
		return false
	}
	for _, a := range aliens.GetAliens() {
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

		aliens := *g.aliens
		for i, a := range aliens {
			if a == nil {
				continue
			}

			xA := a.pos.Element(0, 2)
			yA := a.pos.Element(1, 2)

			withinXAxis := xB > xA && xB < xA+a.w
			withinYAxis := yB > yA && yB < yA+a.h
			if withinXAxis && withinYAxis {
				aliens[i].animation.Explode()
				g.bullets.bullets[j] = nil
			}
		}
	}
}

func bulletsPlayerCollision(g *GameStage) bool {
	for j, b := range g.alienBullets.bullets {
		if b == nil {
			continue
		}

		xB := b.pos.Element(0, 2)
		yB := b.pos.Element(1, 2)

		xA := g.player.pos.Element(0, 2)
		yA := g.player.pos.Element(1, 2)

		withinXAxis := xB > xA && xB < xA+g.player.w
		withinYAxis := yB > yA && yB < yA+g.player.h
		if withinXAxis && withinYAxis {
			g.bullets.bullets[j] = nil

			return true
		}
	}
	return false
}

func aliensShoot(aliens *Aliens, alienBullets *AliensBullets) {
	for _, a := range aliens.GetAliens() {
		if a == nil {
			continue
		}

		if time.Since(alienBullets.lastBulletSpawned) > TimeBetweenAlienBullets {
			alienBullets.Shoot()
			alienBullets.lastBulletSpawned = time.Now()
		}
	}
}
