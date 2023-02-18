package golaga

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func NewBackground(sprite string, speed float64) *Background {
	bg := &Background{}
	bg.bg = artCache[sprite]
	bg.speed = speed

	bg.pos1 = &ebiten.GeoM{}
	bg.pos2 = &ebiten.GeoM{}
	bg.pos2.Translate(0, -WindowHeight)

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

	if b.pos1.Element(1, 2) >= WindowHeight {
		b.pos1.Translate(0, -WindowHeight-WindowHeight)
	}

	if b.pos2.Element(1, 2) >= WindowHeight {
		b.pos2.Translate(0, -WindowHeight-WindowHeight)
	}

	return nil
}

func (b *Background) Draw(screen *ebiten.Image) {

	if b == nil || b.bg == nil || b.pos1 == nil || b.pos2 == nil {
		return
	}

	screen.DrawImage(b.bg, &ebiten.DrawImageOptions{
		GeoM: *b.pos1,
	})

	screen.DrawImage(b.bg, &ebiten.DrawImageOptions{
		GeoM: *b.pos2,
	})
}
