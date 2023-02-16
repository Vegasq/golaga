package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"time"
)

func NewBullets(player *Player) *Bullets {
	return &Bullets{
		player:  player,
		bullets: []*Bullet{},
	}
}

type Bullets struct {
	bullets           []*Bullet
	lastBulletSpawned time.Time

	player *Player
}

func (b *Bullets) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if time.Since(b.lastBulletSpawned) > TimeBetweenBullets {
			b.bullets = append(b.bullets, NewBullet(*b.player.pos))
			b.lastBulletSpawned = time.Now()
		}
	}

	for _, bullet := range b.bullets {
		if err := bullet.Update(); err != nil {
			log.Printf("Error updating bullet: %v", err)
		}
	}
	return nil
}

func (b *Bullets) Draw(screen *ebiten.Image) {
	if b == nil || b.bullets == nil {
		return
	}

	for _, bullet := range b.bullets {
		if bullet == nil {
			continue
		}
		screen.DrawImage(bullet.img, &ebiten.DrawImageOptions{
			GeoM: *bullet.pos,
		})
	}
}

func NewBullet(playerPos ebiten.GeoM) *Bullet {
	bpos := ebiten.GeoM{}
	bpos.Translate(playerPos.Element(0, 2)+47, playerPos.Element(1, 2)-20)
	return &Bullet{artCache["Fire_Shot_4_2"], &bpos}
}

type Bullet struct {
	img *ebiten.Image
	pos *ebiten.GeoM
}

func (b *Bullet) Update() error {
	if b == nil {
		return nil
	}
	b.pos.Translate(0, -1*BulletSpeed)
	return nil
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	screen.DrawImage(b.img, &ebiten.DrawImageOptions{
		GeoM: *b.pos,
	})
}
