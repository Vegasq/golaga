package main

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	img *ebiten.Image
	pos *ebiten.GeoM

	w, h float64
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

func (p *Player) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.pos.Translate(PlayerSpeed, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.pos.Translate(-1*PlayerSpeed, 0)
	}

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	screen.DrawImage(p.img, &ebiten.DrawImageOptions{
		GeoM: *p.pos,
	})
}
