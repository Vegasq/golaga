package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewBackground(sprite string) *Background {
	bg := &Background{}
	bg.bg = artCache[sprite]

	bg.pos1 = &ebiten.GeoM{}
	bg.pos2 = &ebiten.GeoM{}
	bg.pos2.Translate(0, -1920)

	return bg
}

type Background struct {
	bg   *ebiten.Image
	pos1 *ebiten.GeoM
	pos2 *ebiten.GeoM
}

func (b *Background) Update() error {
	return nil
}

func (b *Background) Draw(screen *ebiten.Image) {
	b.pos1.Translate(0, 1)
	screen.DrawImage(b.bg, &ebiten.DrawImageOptions{
		GeoM: *b.pos1,
	})

	b.pos2.Translate(0, 1)
	screen.DrawImage(b.bg, &ebiten.DrawImageOptions{
		GeoM: *b.pos2,
	})

	if b.pos1.Element(1, 2) >= 1920 {
		b.pos1.Translate(0, -1920-1920)
	}

	if b.pos2.Element(1, 2) >= 1920 {
		fmt.Println("moving pos2")
		b.pos2.Translate(0, -1920-1920)
	}
}
