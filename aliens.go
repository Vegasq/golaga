package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func NewAliens() *Aliens {
	var aliens Aliens

	var i float64
	var j float64
	for i = 100; i < 1000; i += 200 {
		for j = 0; j < 150; j += 120 {
			alien := NewAlien(i, j)
			aliens = append(aliens, alien)
		}
	}

	AlienDissentSpeed = float64(AlienDissentSpeed) * 1.1

	return &aliens
}

type Aliens []*Alien

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
			return
		}
		screen.DrawImage(nextFrame, &ebiten.DrawImageOptions{
			GeoM: *alien.pos,
		})
	}
}

func NewAlien(x, y float64) *Alien {
	animation := buildAlienAnimation()

	alien := &Alien{
		alive:     true,
		animation: animation,
		pos:       &ebiten.GeoM{},

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
	a.pos.Translate(a.wiggle, float64(AlienDissentSpeed))
}
