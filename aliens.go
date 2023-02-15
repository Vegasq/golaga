package main

import "github.com/hajimehoshi/ebiten/v2"

type Aliens []*Alien

func (a Aliens) Update() error {
	for _, alien := range a {
		alien.Update()
	}
	return nil
}

func (a Aliens) Draw(screen *ebiten.Image) {
	for _, alien := range a {
		if alien == nil {
			continue
		}
		screen.DrawImage(alien.img, &ebiten.DrawImageOptions{
			GeoM: *alien.pos,
		})
	}
}

type Alien struct {
	img *ebiten.Image
	pos *ebiten.GeoM

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

func NewAlien(x, y float64) *Alien {
	w, h := artCache["Ship_06"].Size()
	alien := &Alien{
		img: artCache["Ship_06"],
		pos: &ebiten.GeoM{},

		w: float64(w), h: float64(h),
	}
	alien.pos.Translate(x, y)
	alien.wiggleCountdown = 0

	return alien
}

func NewAliens() Aliens {
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

	return aliens
}
