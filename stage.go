package golaga

import "github.com/hajimehoshi/ebiten/v2"

type Stage interface {
	Draw(screen *ebiten.Image)
	Update() error

	Reset()
}
