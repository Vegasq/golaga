package golaga

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const WindowWidth = 1080
const WindowHeight = 1920

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
	return WindowWidth, WindowHeight
}

func PrepareGame() *Game {
	ebiten.SetWindowTitle("Golaga")
	ebiten.SetWindowSize(540, 960)

	loadArt()

	changeStage := make(chan string)
	var stages = map[string]Stage{
		"game":     &GameStage{changeStage: changeStage},
		"menu":     &MenuStage{changeStage: changeStage},
		"gameover": &GameOverStage{changeStage: changeStage},
		"prepare":  &PrepareStage{changeStage: changeStage},
		"theend":   &TheEndStage{changeStage: changeStage},
	}
	game := Game{stages: stages, currentStage: "menu", changeStageChan: changeStage}

	go func() {
		for {
			newStage := <-changeStage
			stages[newStage].Reset()
			game.currentStage = newStage
		}
	}()
	return &game
}

func main() {
	game := PrepareGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
