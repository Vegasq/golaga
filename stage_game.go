package main

import (
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image"
	"log"
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

	missions []*Mission

	bullets           *Bullets
	alienBullets      *AliensBullets
	lastBulletSpawned time.Time

	background *Background
	init       bool

	score int
}

func (g *GameStage) Update() error {
	portal := make(chan *Step)

	if g.init == false {
		if g.missions == nil {
			g.missions = GetMissions()
		}

		g.player = NewPlayer()
		g.bullets = NewBullets(g.player)
		g.aliens = NewAliens()
		g.background = NewBackground(g.missions[MissionIndex].BackgroundSprite, 5)
		g.alienBullets = NewAliensBullets(g.aliens)

		g.init = true

		go AliensArrival(g, portal)
		go MissionPlayer(portal, g.missions[MissionIndex])
	}

	//if g.player == nil || g.aliens == nil || g.bullets == nil || g.background == nil || g.alienBullets == nil {
	//	return nil
	//}

	if err := g.player.Update(); err != nil {
		log.Printf("error in player update: %v", err)
	}
	if err := g.bullets.Update(); err != nil {
		log.Printf("error in bullets update: %v", err)
	}
	if err := g.background.Update(); err != nil {
		log.Printf("error in background update: %v", err)
	}
	if err := g.aliens.Update(); err != nil {
		log.Printf("error in aliens update: %v", err)
	}
	if err := g.alienBullets.Update(); err != nil {
		log.Printf("error in alien bullets update: %v", err)
	}

	aliensShoot(g.aliens, g.alienBullets)

	if haveAliveAliens(g.aliens) == false && g.missions[MissionIndex].Done && MissionIndex < len(g.missions)-1 {
		log.Println("move to prepare stage")
		g.init = false
		MissionIndex++
		g.changeStage <- "prepare"
	} else if haveAliveAliens(g.aliens) == false && g.missions[MissionIndex].Done && MissionIndex == len(g.missions)-1 {
		log.Println("move to prepare theend")
		g.init = false
		MissionIndex = 0
		g.missions = nil
		g.changeStage <- "theend"
	}

	bulletsAlienCollision(g)

	if bulletsPlayerCollision(g) {
		g.player.animation.Explode()
	}

	if g.player.alive == false {
		g.init = false
		log.Println("move to gameover stage")
		MissionIndex = 0
		g.missions = nil
		g.changeStage <- "gameover"
	}

	return nil
}

func (g *GameStage) Draw(screen *ebiten.Image) {
	if g.aliens == nil || g.player == nil || g.bullets == nil || g.background == nil || g.alienBullets == nil {
		return
	}

	if alienTouchedTheGround(screen, g.aliens) {
		g.init = false
		MissionIndex = 0
		g.missions = nil
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
	text.Draw(screen, fmt.Sprintf("Score: %d", g.score), faces[3], 450, 100, image.White)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
}

func (g *GameStage) Reset() {
	g.player = nil
	g.aliens = nil
	g.bullets = nil
	g.alienBullets = nil
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
	for j, b := range g.bullets.GetBullets() {
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
			if withinXAxis && withinYAxis && g.bullets.bullets[j] != nil && aliens[i].animation.exploding == false {
				aliens[i].animation.Explode()
				ShipExplode()
				g.bullets.bullets[j] = nil
				g.score += 1
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
			if g == nil || g.bullets == nil {
				continue
			}
			g.alienBullets.bullets[j] = nil

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
