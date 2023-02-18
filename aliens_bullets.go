package golaga

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"time"
)

func NewAliensBullets(aliens *Aliens) *AliensBullets {
	return &AliensBullets{
		bullets: []*AlienBullet{},
		aliens:  aliens,
	}
}

type AliensBullets struct {
	bullets           []*AlienBullet
	aliens            *Aliens
	lastBulletSpawned time.Time
}

func (b *AliensBullets) Shoot() {
	aliens := *b.aliens
	cleanAliens := []*Alien{}
	for _, alien := range aliens {
		if alien != nil && alien.alive {
			cleanAliens = append(cleanAliens, alien)
		}
	}

	total := len(cleanAliens)
	if total == 0 {
		return
	}

	p := rand.Intn(total)
	bullet := NewAlienBullet(*cleanAliens[p].pos)
	b.bullets = append(b.bullets, bullet)
}

func (b *AliensBullets) Update() error {
	for _, bullet := range b.bullets {
		bullet.Update()
	}
	return nil
}

func (b *AliensBullets) Draw(screen *ebiten.Image) {
	for _, bullet := range b.bullets {
		if bullet == nil {
			continue
		}
		screen.DrawImage(bullet.img, &ebiten.DrawImageOptions{
			GeoM: *bullet.pos,
		})
	}
}

func NewAlienBullet(alienPos ebiten.GeoM) *AlienBullet {
	bpos := ebiten.GeoM{}
	bpos.Translate(alienPos.Element(0, 2)+47, alienPos.Element(1, 2)+100)
	return &AlienBullet{artCache["Laser_1_4"], &bpos}
}

type AlienBullet struct {
	img *ebiten.Image
	pos *ebiten.GeoM
}

func (b *AlienBullet) Update() error {
	if b == nil {
		return nil
	}
	b.pos.Translate(0, 1*BulletSpeed)
	return nil
}

func (b *AlienBullet) Draw(screen *ebiten.Image) {
	screen.DrawImage(b.img, &ebiten.DrawImageOptions{
		GeoM: *b.pos,
	})
}
