package main

import (
	"log"
	"time"
)

type Step struct {
	Wait   time.Duration
	Aliens []*Alien
}

type Mission struct {
	BackgroundSprite string
	Steps            []Step
	Done             bool
}

func BuildMission1() *Mission {
	padding := float64(WindowWidth / 6)
	baseSpeed := float64(2.0)
	var mission1 = []Step{

		// Wave 1
		{Wait: time.Duration(1 * time.Second), Aliens: []*Alien{
			NewAlien(padding*3, -100, baseSpeed),
		}},

		{Wait: time.Duration(3 * time.Second), Aliens: []*Alien{
			NewAlien(padding*3, -100, baseSpeed),
		}},
		{Wait: time.Duration(4 * time.Second), Aliens: []*Alien{
			NewAlien(padding*2, -100, baseSpeed),
			NewAlien(padding*4, -100, baseSpeed),
		}},
		{Wait: time.Duration(4 * time.Second), Aliens: []*Alien{
			NewAlien(padding*2, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed),
			NewAlien(padding*4, -100, baseSpeed),
		}},

		// Wave 2
		{Wait: time.Duration(3 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding*2, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding*3, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding*4, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding*5, -100, baseSpeed),
		}},

		// Wave 1

		{Wait: time.Duration(4 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*2, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed),
			NewAlien(padding*4, -100, baseSpeed),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(4 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*2, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed),
			NewAlien(padding*4, -100, baseSpeed),
			NewAlien(padding*5, -100, baseSpeed),
		}},

		{Wait: time.Duration(5 * time.Second), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed),
			NewAlien(padding*2, -100, baseSpeed*2),
			NewAlien(padding*3, -100, baseSpeed*3),
			NewAlien(padding*4, -100, baseSpeed*2),
			NewAlien(padding*5, -100, baseSpeed),
		}},

		{Wait: time.Duration(4 * time.Second), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed*3),
			NewAlien(padding*2, -100, baseSpeed*2),
			NewAlien(padding*3, -100, baseSpeed),
			NewAlien(padding*4, -100, baseSpeed*2),
			NewAlien(padding*5, -100, baseSpeed*3),
		}},
		{Wait: time.Duration(3 * time.Second), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed),
			NewAlien(padding*2, -100, baseSpeed*2),
			NewAlien(padding*3, -100, baseSpeed*4),
			NewAlien(padding*4, -100, baseSpeed*2),
			NewAlien(padding*5, -100, baseSpeed),
		}},
	}
	return &Mission{
		BackgroundSprite: "Space_BG_01",
		Steps:            mission1,
	}
}
func BuildMission2() *Mission {
	baseSpeed := float64(3.0)
	padding := float64(WindowWidth / 6)
	var mission = []Step{
		// Wave 1
		{Wait: time.Duration(1 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed*2),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed*2),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed*2),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*5, -100, baseSpeed),
		}},

		// Wave 2
		{Wait: time.Duration(4 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*2, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed),
			NewAlien(padding*4, -100, baseSpeed),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(4 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*2, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed),
			NewAlien(padding*4, -100, baseSpeed),
			NewAlien(padding*5, -100, baseSpeed),
		}},

		{Wait: time.Duration(5 * time.Second), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed),
			NewAlien(padding*2, -100, baseSpeed*2),
			NewAlien(padding*3, -100, baseSpeed*3),
			NewAlien(padding*4, -100, baseSpeed*2),
			NewAlien(padding*5, -100, baseSpeed),
		}},

		{Wait: time.Duration(4 * time.Second), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed*3),
			NewAlien(padding*2, -100, baseSpeed*2),
			NewAlien(padding*3, -100, baseSpeed),
			NewAlien(padding*4, -100, baseSpeed*2),
			NewAlien(padding*5, -100, baseSpeed*3),
		}},
		{Wait: time.Duration(3 * time.Second), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed),
			NewAlien(padding*2, -100, baseSpeed*2),
			NewAlien(padding*3, -100, baseSpeed*4),
			NewAlien(padding*4, -100, baseSpeed*2),
			NewAlien(padding*5, -100, baseSpeed),
		}},

		// Wave 3
		{Wait: time.Duration(1 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed*2),
		}},
		{Wait: time.Duration(1 * time.Second), Aliens: []*Alien{
			NewAlien(padding*3, -100, baseSpeed*2),
		}},
		{Wait: time.Duration(1 * time.Second), Aliens: []*Alien{
			NewAlien(padding*5, -100, baseSpeed*2),
		}},
		{Wait: time.Duration(1 * time.Second), Aliens: []*Alien{
			NewAlien(padding*2, -100, baseSpeed*2),
		}},
		{Wait: time.Duration(1 * time.Second), Aliens: []*Alien{
			NewAlien(padding*3, -100, baseSpeed*2),
		}},
		{Wait: time.Duration(1 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed*2),
		}},
		{Wait: time.Duration(1 * time.Second), Aliens: []*Alien{
			NewAlien(padding*3, -100, baseSpeed*2),
		}},
		{Wait: time.Duration(1 * time.Second), Aliens: []*Alien{
			NewAlien(padding*5, -100, baseSpeed*2),
		}},
		{Wait: time.Duration(1 * time.Second), Aliens: []*Alien{
			NewAlien(padding*2, -100, baseSpeed*2),
		}},
		{Wait: time.Duration(1 * time.Second), Aliens: []*Alien{
			NewAlien(padding*3, -100, baseSpeed*2),
		}},
	}
	return &Mission{
		BackgroundSprite: "Space_BG_02",
		Steps:            mission,
	}
}
func BuildMission3() *Mission {
	baseSpeed := float64(3.2)
	padding := float64(WindowWidth / 6)
	var mission = []Step{
		// Wave 1
		{Wait: time.Duration(1 * time.Second), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed*1.1),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed*1.1),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed*1.1),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed*1.1),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed*1.1),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed*1.1),
			NewAlien(padding*5, -100, baseSpeed),
		}},

		// Wave 2
		{Wait: time.Duration(1 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed*2),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed*2),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed*2),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(2 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*5, -100, baseSpeed),
		}},

		{Wait: time.Duration(3 * time.Second), Aliens: []*Alien{
			NewAlien(padding*2, -100, baseSpeed),
			NewAlien(padding*2+10, -100, baseSpeed),
			NewAlien(padding*2+20, -100, baseSpeed),
			NewAlien(padding*2+30, -100, baseSpeed),
			NewAlien(padding*2+40, -100, baseSpeed),
			NewAlien(padding*2+50, -100, baseSpeed),
			NewAlien(padding*2+60, -100, baseSpeed),
		}},

		// Wave 3

		{Wait: time.Duration(4 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*2, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed),
			NewAlien(padding*4, -100, baseSpeed),
			NewAlien(padding*5, -100, baseSpeed),
		}},
		{Wait: time.Duration(4 * time.Second), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed),
			NewAlien(padding*2, -100, baseSpeed),
			NewAlien(padding*3, -100, baseSpeed),
			NewAlien(padding*4, -100, baseSpeed),
			NewAlien(padding*5, -100, baseSpeed),
		}},

		{Wait: time.Duration(4 * time.Second), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed),
			NewAlien(padding*2, -100, baseSpeed*2),
			NewAlien(padding*3, -100, baseSpeed*3),
			NewAlien(padding*4, -100, baseSpeed*2),
			NewAlien(padding*5, -100, baseSpeed),
		}},

		{Wait: time.Duration(4 * time.Second), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed*3),
			NewAlien(padding*2, -100, baseSpeed*2),
			NewAlien(padding*3, -100, baseSpeed),
			NewAlien(padding*4, -100, baseSpeed*2),
			NewAlien(padding*5, -100, baseSpeed*3),
		}},
		{Wait: time.Duration(3 * time.Second), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed),
			NewAlien(padding*2, -100, baseSpeed*2),
			NewAlien(padding*3, -100, baseSpeed*4),
			NewAlien(padding*4, -100, baseSpeed*2),
			NewAlien(padding*5, -100, baseSpeed),
		}},

		// Wave 4

		{Wait: time.Duration(1 * time.Second), Aliens: []*Alien{
			NewAlien(padding*2+200, -100, baseSpeed),
			NewAlien(padding*2+250, -100, baseSpeed),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*2+150, -100, baseSpeed),
			NewAlien(padding*2+200, -100, baseSpeed),
			NewAlien(padding*2+250, -100, baseSpeed),
			NewAlien(padding*2+300, -100, baseSpeed),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*2+50, -100, baseSpeed),
			NewAlien(padding*2+150, -100, baseSpeed),
			NewAlien(padding*2+200, -100, baseSpeed),
			NewAlien(padding*2+250, -100, baseSpeed),
			NewAlien(padding*2+300, -100, baseSpeed),
			NewAlien(padding*2+350, -100, baseSpeed),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*2+50, -100, baseSpeed),
			NewAlien(padding*2+150, -100, baseSpeed),
			NewAlien(padding*2+200, -100, baseSpeed),
			NewAlien(padding*2+250, -100, baseSpeed),
			NewAlien(padding*2+300, -100, baseSpeed),
			NewAlien(padding*2+350, -100, baseSpeed),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*2+50, -100, baseSpeed),
			NewAlien(padding*2+150, -100, baseSpeed),
			NewAlien(padding*2+200, -100, baseSpeed),
			NewAlien(padding*2+250, -100, baseSpeed),
			NewAlien(padding*2+300, -100, baseSpeed),
			NewAlien(padding*2+350, -100, baseSpeed),
		}},
	}
	return &Mission{
		BackgroundSprite: "Space_BG_03",
		Steps:            mission,
	}
}

var MissionIndex = 0

func GetMissions() []*Mission {
	var Missions = []*Mission{
		BuildMission1(),
		BuildMission2(),
		BuildMission3(),
	}
	return Missions
}

func MissionPlayer(portal chan *Step, mission *Mission) {
	for i := range mission.Steps {
		time.Sleep(mission.Steps[i].Wait)
		log.Println("Sending aliens")
		log.Println(mission.Steps[i])
		portal <- &mission.Steps[i]
	}
	close(portal)
	mission.Done = true
	log.Println("Closing mission player")
}

func AliensArrival(g *GameStage, portal chan *Step) {
	for {
		if g == nil || g.aliens == nil {
			continue
		}
		step, ok := <-portal

		if ok == false {
			return
		}

		for _, alien := range step.Aliens {
			log.Println("Alien arrived")
			g.aliens.AddAlien(alien)
		}
	}
}
