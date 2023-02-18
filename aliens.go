package golaga

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func NewAliens() *Aliens {
	var aliens Aliens
	return &aliens
}

type Aliens []*Alien

func (a *Aliens) AddAlien(alien *Alien) {
	*a = append(*a, alien)
}

func (a Aliens) GetAliens() []*Alien {
	aliens := []*Alien{}
	for _, alien := range a {
		if alien.alive {
			aliens = append(aliens, alien)
		}
	}
	return aliens
}

func (a Aliens) Update() error {
	for _, alien := range a {
		if alien.alive {
			alien.Update()
		}
	}
	return nil
}

func (a Aliens) Draw(screen *ebiten.Image) {
	for _, alien := range a {
		if alien.alive == false {
			continue
		}
		if alien == nil {
			continue
		}
		nextFrame := alien.animation.GetFrame()
		if nextFrame == nil {
			alien.alive = false
			continue
		}
		screen.DrawImage(nextFrame, &ebiten.DrawImageOptions{
			GeoM: *(alien).pos,
		})
	}
}

func NewAlien(x, y, speed float64, animationName string) *Alien {
	animations := GetAnimations()
	animation := animations[animationName]
	pos := ebiten.GeoM{}
	alien := &Alien{
		alive:     true,
		animation: animation,
		pos:       &pos,

		speed: speed,

		w: float64(animation.w), h: float64(animation.h),
	}

	alien.pos.Translate(x, y)
	alien.wiggleCountdown = 0

	return alien
}

type Alien struct {
	//img       *ebiten.Image
	alive     bool
	animation Animation
	pos       *ebiten.GeoM

	speed float64

	w, h float64

	wiggleDirection bool
	wiggle          float64
	wiggleCountdown int
}

func (a *Alien) Update() {
	if a == nil {
		return
	}
	if a.wiggleCountdown > 50 {
		a.wiggleCountdown = 0
		a.wiggleDirection = !a.wiggleDirection
	}

	a.wiggleCountdown += 1
	if a.wiggleDirection {
		a.wiggle = 1
	} else {
		a.wiggle = -1
	}
	a.pos.Translate(a.wiggle, a.speed)
}
