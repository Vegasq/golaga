package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func NewBackground(sprite string, speed float64) *Background {
	bg := &Background{}
	bg.bg = artCache[sprite]
	bg.speed = speed

	bg.pos1 = &ebiten.GeoM{}
	bg.pos2 = &ebiten.GeoM{}
	bg.pos2.Translate(0, -1920)

	return bg
}

type Background struct {
	bg    *ebiten.Image
	pos1  *ebiten.GeoM
	pos2  *ebiten.GeoM
	speed float64
}

func (b *Background) Update() error {
	b.pos1.Translate(0, b.speed)
	b.pos2.Translate(0, b.speed)

	if b.pos1.Element(1, 2) >= 1920 {
		b.pos1.Translate(0, -1920-1920)
	}

	if b.pos2.Element(1, 2) >= 1920 {
		b.pos2.Translate(0, -1920-1920)
	}

	return nil
}

func (b *Background) Draw(screen *ebiten.Image) {
	screen.DrawImage(b.bg, &ebiten.DrawImageOptions{
		GeoM: *b.pos1,
	})

	screen.DrawImage(b.bg, &ebiten.DrawImageOptions{
		GeoM: *b.pos2,
	})
}
