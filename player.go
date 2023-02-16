package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func NewPlayer() *Player {
	player := &Player{}

	player.animation = buildPlayerAnimation()

	player.w = float64(player.animation.w)
	player.h = float64(player.animation.h)

	player.pos = &ebiten.GeoM{}
	player.pos.Translate(100, 1720)
	player.alive = true
	return player
}

type Player struct {
	//img       *ebiten.Image
	animation *Animation
	pos       *ebiten.GeoM
	alive     bool

	w, h float64
}

func (p *Player) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.pos.Translate(PlayerSpeed, 0)

		xPos := p.pos.Element(0, 2)
		if xPos > WindowWidth-p.w {
			diff := WindowWidth - p.w - xPos
			p.pos.Translate(diff, 0)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.pos.Translate(-1*PlayerSpeed, 0)

		xPos := p.pos.Element(0, 2)
		if xPos < 0 {
			p.pos.Translate(-1*xPos, 0)
		}
	}

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	if p == nil || p.animation == nil {
		return
	}

	img := p.animation.GetFrame()
	if img == nil {
		p.alive = false
		return
	}
	screen.DrawImage(img, &ebiten.DrawImageOptions{
		GeoM: *p.pos,
	})
}
