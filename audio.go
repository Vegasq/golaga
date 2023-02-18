package main

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"time"
)

const (
	sampleRate = 32000
)

//go:embed sound/lavik89/168984__lavik89__digital-hit.wav
var GunSound []byte

//go:embed sound/MATRIXXX_/403298__matrixxx__retro_explosion_02.wav
var ExplosionSound []byte

var Context *audio.Context

func Bang() {
	if Context == nil {
		Context = audio.NewContext(sampleRate)
	}
	player := Context.NewPlayerFromBytes(GunSound)
	player.Seek(time.Duration(100 * time.Millisecond))
	player.Play()
}

func ShipExplode() {
	if Context == nil {
		Context = audio.NewContext(sampleRate)
	}
	player := Context.NewPlayerFromBytes(ExplosionSound)
	player.SetVolume(0.02)
	player.Seek(time.Duration(100 * time.Millisecond))
	player.Play()
}
