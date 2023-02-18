package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"time"
)

type Animation struct {
	frames []*ebiten.Image
	total  int

	expFrames []*ebiten.Image
	expTotal  int

	pos  int
	w, h int

	lastUpdate time.Time
	exploding  bool
}

const AnimationTimeFrame = time.Duration(75 * time.Millisecond)

func (a *Animation) Explode() {
	if a.exploding == false {
		a.exploding = true
		a.pos = 0
	}
}

func (a *Animation) GetFrame() *ebiten.Image {
	if a.exploding {
		log.Println("exploding frame")
		if time.Since(a.lastUpdate) > AnimationTimeFrame {
			a.pos++
			if a.pos >= a.expTotal {
				return nil
			}
			a.lastUpdate = time.Now()
			return a.expFrames[a.pos]
		}
		return a.expFrames[a.pos]
	} else {
		if time.Since(a.lastUpdate) > AnimationTimeFrame {
			a.pos++
			if a.pos >= a.total {
				a.pos = 0
			}
			a.lastUpdate = time.Now()
			return a.frames[a.pos]
		}
		return a.frames[a.pos]
	}
}

func buildPlayerAnimation() Animation {
	w, h := artCache["Player_0_0"].Size()
	return Animation{
		frames: []*ebiten.Image{
			artCache["Player_0_0"],
			artCache["Player_0_1"],
			artCache["Player_0_2"],
			artCache["Player_0_3"],
			artCache["Player_0_4"],
			artCache["Player_0_5"],
			artCache["Player_0_6"],
			artCache["Player_0_7"],
			artCache["Player_0_8"],
			artCache["Player_0_9"],
		},
		expFrames: []*ebiten.Image{
			artCache["Player_exp_0_0"],
			artCache["Player_exp_0_1"],
			artCache["Player_exp_0_2"],
			artCache["Player_exp_0_3"],
			artCache["Player_exp_0_4"],
			artCache["Player_exp_0_5"],
			artCache["Player_exp_0_6"],
			artCache["Player_exp_0_7"],
			artCache["Player_exp_0_8"],
		},
		pos:       0,
		total:     10,
		expTotal:  9,
		exploding: false,
		w:         w, h: h,
	}
}

func buildAlienAnimation() Animation {
	var w, h int

	w, h = artCache["Alien_0_0"].Size()
	a := Animation{
		frames: []*ebiten.Image{
			artCache["Alien_0_0"],
			artCache["Alien_0_1"],
			artCache["Alien_0_2"],
			artCache["Alien_0_3"],
			artCache["Alien_0_4"],
			artCache["Alien_0_5"],
			artCache["Alien_0_6"],
			artCache["Alien_0_7"],
			artCache["Alien_0_8"],
			artCache["Alien_0_9"],
		},
		expFrames: []*ebiten.Image{
			artCache["Alien_exp_0_0"],
			artCache["Alien_exp_0_1"],
			artCache["Alien_exp_0_2"],
			artCache["Alien_exp_0_3"],
			artCache["Alien_exp_0_4"],
			artCache["Alien_exp_0_5"],
			artCache["Alien_exp_0_6"],
			artCache["Alien_exp_0_7"],
			artCache["Alien_exp_0_8"],
		},
		pos:      0,
		total:    10,
		expTotal: 9,
		w:        w, h: h,
	}

	return a
}

func buildAlien1Animation() Animation {
	var w, h int

	w, h = artCache["Alien_0_0"].Size()
	a := Animation{
		frames: []*ebiten.Image{
			artCache["Alien_1_0"],
			artCache["Alien_1_1"],
			artCache["Alien_1_2"],
			artCache["Alien_1_3"],
			artCache["Alien_1_4"],
			artCache["Alien_1_5"],
			artCache["Alien_1_6"],
			artCache["Alien_1_7"],
			artCache["Alien_1_8"],
			artCache["Alien_1_9"],
		},
		expFrames: []*ebiten.Image{
			artCache["Alien_exp_1_0"],
			artCache["Alien_exp_1_1"],
			artCache["Alien_exp_1_2"],
			artCache["Alien_exp_1_3"],
			artCache["Alien_exp_1_4"],
			artCache["Alien_exp_1_5"],
			artCache["Alien_exp_1_6"],
			artCache["Alien_exp_1_7"],
			artCache["Alien_exp_1_8"],
		},
		pos:      0,
		total:    10,
		expTotal: 9,
		w:        w, h: h,
	}

	return a
}

var animations map[string]Animation
var PlayerAnimationName = "Player"
var AliensAnimationsNames = []string{"Alien_0", "Alien_1"}

func GetAnimations() map[string]Animation {
	if animations == nil {
		animations = map[string]Animation{
			PlayerAnimationName:      buildPlayerAnimation(),
			AliensAnimationsNames[0]: buildAlienAnimation(),
			AliensAnimationsNames[1]: buildAlien1Animation(),
		}
	}
	return animations
}
