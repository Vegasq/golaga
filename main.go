package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct {
	stages          map[string]Stage
	currentStage    string
	changeStageChan chan string
}

func (g *Game) Update() error {
	return g.stages[g.currentStage].Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.stages[g.currentStage].Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1080, 1920
}

func main() {
	ebiten.SetWindowTitle("Golaga")
	ebiten.SetWindowSize(540, 960)

	loadArt()

	changeStage := make(chan string)
	var stages = map[string]Stage{
		"game":     &GameStage{changeStage: changeStage},
		"menu":     &MenuStage{changeStage, nil},
		"gameover": &GameOverStage{changeStage},
	}
	game := Game{stages: stages, currentStage: "menu", changeStageChan: changeStage}

	go func() {
		for {
			newStage := <-changeStage
			stages[newStage].Reset()
			game.currentStage = newStage
		}
	}()

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
