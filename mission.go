package main

import (
	"log"
	"math/rand"
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

func BuildSleepStep(sleep time.Duration) []Step {
	return []Step{
		{Wait: sleep, Aliens: []*Alien{}},
	}
}

func BuildPattern1(baseSpeed float64, animationName string) []Step {
	// A     A
	//  A   A
	//    A
	padding := float64(WindowWidth / 6)

	return []Step{
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*3, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*2, -100, baseSpeed, animationName),
			NewAlien(padding*4, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed, animationName),
			NewAlien(padding*5, -100, baseSpeed, animationName),
		}},
	}
}

func BuildPattern2(baseSpeed float64, animationName string) []Step {
	// A
	//   A
	//     A
	//       A
	//         A

	padding := float64(WindowWidth / 6)

	return []Step{
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*2, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*3, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*4, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*5, -100, baseSpeed, animationName),
		}},
	}
}

func BuildPattern3(baseSpeed float64, animationName string) []Step {
	//         A
	//       A
	//     A
	//   A
	// A

	padding := float64(WindowWidth / 6)

	return []Step{
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*5, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*4, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*3, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*2, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed, animationName),
		}},
	}
}

func BuildPattern4(baseSpeed float64, animationName string) []Step {
	// A A A A A
	// A A A A A
	padding := float64(WindowWidth / 6)

	return []Step{
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed, animationName),
			NewAlien(padding*2, -100, baseSpeed*2, animationName),
			NewAlien(padding*3, -100, baseSpeed*3, animationName),
			NewAlien(padding*4, -100, baseSpeed*2, animationName),
			NewAlien(padding*5, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(1 * time.Second), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed*3, animationName),
			NewAlien(padding*2, -100, baseSpeed*2, animationName),
			NewAlien(padding*3, -100, baseSpeed, animationName),
			NewAlien(padding*4, -100, baseSpeed*2, animationName),
			NewAlien(padding*5, -100, baseSpeed*3, animationName),
		}},
	}
}

func BuildPattern5(baseSpeed float64, animationName string) []Step {
	// DaBoss
	padding := float64(WindowWidth / 6)

	return []Step{
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*2+200, -100, baseSpeed, animationName),
			NewAlien(padding*2+250, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*2+150, -100, baseSpeed, animationName),
			NewAlien(padding*2+200, -100, baseSpeed, animationName),
			NewAlien(padding*2+250, -100, baseSpeed, animationName),
			NewAlien(padding*2+300, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*2+50, -100, baseSpeed, animationName),
			NewAlien(padding*2+150, -100, baseSpeed, animationName),
			NewAlien(padding*2+200, -100, baseSpeed, animationName),
			NewAlien(padding*2+250, -100, baseSpeed, animationName),
			NewAlien(padding*2+300, -100, baseSpeed, animationName),
			NewAlien(padding*2+350, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*2+50, -100, baseSpeed, animationName),
			NewAlien(padding*2+150, -100, baseSpeed, animationName),
			NewAlien(padding*2+200, -100, baseSpeed, animationName),
			NewAlien(padding*2+250, -100, baseSpeed, animationName),
			NewAlien(padding*2+300, -100, baseSpeed, animationName),
			NewAlien(padding*2+350, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*2+50, -100, baseSpeed, animationName),
			NewAlien(padding*2+150, -100, baseSpeed, animationName),
			NewAlien(padding*2+200, -100, baseSpeed, animationName),
			NewAlien(padding*2+250, -100, baseSpeed, animationName),
			NewAlien(padding*2+300, -100, baseSpeed, animationName),
			NewAlien(padding*2+350, -100, baseSpeed, animationName),
		}},
	}
}

func BuildPattern6(baseSpeed float64, animationName string) []Step {
	// DaBoss
	padding := float64(WindowWidth / 6)

	return []Step{
		{Wait: time.Duration(500 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*3, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(300 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*2, -100, baseSpeed, animationName),
			NewAlien(padding*4, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(300 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed, animationName),
			NewAlien(padding*5, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(1000 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed, animationName),
			NewAlien(padding*5, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(1000 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed, animationName),
			NewAlien(padding*5, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(1000 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*1, -100, baseSpeed, animationName),
			NewAlien(padding*5, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(300 * time.Millisecond), Aliens: []*Alien{
			NewAlien(padding*2, -100, baseSpeed, animationName),
			NewAlien(padding*4, -100, baseSpeed, animationName),
		}},
		{Wait: time.Duration(0 * time.Second), Aliens: []*Alien{
			NewAlien(padding*3, -100, baseSpeed, animationName),
		}},
	}
}

func BuildRandomMission(background string, s, m, b int, speed float64, padding time.Duration) *Mission {
	simple := []func(float64, string) []Step{
		BuildPattern2,
		BuildPattern3,
	}
	medium := []func(float64, string) []Step{
		BuildPattern1,
		BuildPattern4,
		BuildPattern6,
	}
	bosses := []func(float64, string) []Step{
		BuildPattern5,
	}

	patternsBuckets := map[string][]func(float64, string) []Step{
		"simple": simple,
		"medium": medium,
		"bosses": bosses,
	}

	mission := []Step{}

	randAnimation := AliensAnimationsNames[rand.Intn(len(AliensAnimationsNames))]
	for i := 0; i < s; i++ {
		mission = append(mission, patternsBuckets["simple"][rand.Intn(len(patternsBuckets["simple"]))](speed, randAnimation)...)
	}
	mission = append(mission, BuildSleepStep(padding)...)

	randAnimation = AliensAnimationsNames[rand.Intn(len(AliensAnimationsNames))]
	for i := 0; i < m; i++ {
		mission = append(mission, patternsBuckets["medium"][rand.Intn(len(patternsBuckets["medium"]))](speed, randAnimation)...)
	}
	mission = append(mission, BuildSleepStep(padding)...)

	randAnimation = AliensAnimationsNames[rand.Intn(len(AliensAnimationsNames))]
	for i := 0; i < b; i++ {
		mission = append(mission, patternsBuckets["bosses"][rand.Intn(len(patternsBuckets["bosses"]))](speed, randAnimation)...)
	}

	return &Mission{
		BackgroundSprite: background,
		Steps:            mission,
	}
}

var MissionIndex = 0

func GetMissions() []*Mission {
	var Missions = []*Mission{
		BuildRandomMission("Space_BG_01", 5, 1, 1, 2.0, 3000*time.Millisecond),
		BuildRandomMission("Space_BG_02", 4, 2, 1, 2.1, 3000*time.Millisecond),
		BuildRandomMission("Space_BG_03", 2, 4, 1, 2.2, 3000*time.Millisecond),

		BuildRandomMission("Space_BG_01", 3, 4, 1, 2.2, 3000*time.Millisecond),
		BuildRandomMission("Space_BG_02", 4, 4, 1, 2.3, 3000*time.Millisecond),
		BuildRandomMission("Space_BG_03", 4, 5, 1, 2.4, 3000*time.Millisecond),

		BuildRandomMission("Space_BG_01", 5, 6, 1, 2.4, 3000*time.Millisecond),
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
