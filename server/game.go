package main

import "math/rand"

const (
	CIVILIAN uint8 = iota
	MAFIA
	CHERIFF
	GHOST

	PLAYER
)

type Game struct {
	players *map[string]Player
	isNight bool
	vote    map[string]int
}

func (g *Game) MafiaCount() uint8 {
	var count uint8 = 0
	for _, player := range *g.players {
		if player.role == MAFIA {
			count++
		}
	}
	return count
}

func (g *Game) AssignRoles() {
	usernames := make([]string, len(*g.players))
	i := 0
	for name := range *g.players {
		usernames[i] = name
		i++
	}
	for i := range usernames {
		j := rand.Intn(i + 1)
		usernames[i], usernames[j] = usernames[j], usernames[i]
	}

	var mafiaCount int = len(*g.players) / 3
	i = 0
	for _, username := range usernames {
		if player, ok := (*g.players)[username]; ok {
			if i < mafiaCount {
				player.role = MAFIA
			} else if i == mafiaCount {
				player.role = CHERIFF
			} else {
				player.role = CIVILIAN
			}
		}
		i++
	}
}

func (g *Game) ChangeTime() {
	g.isNight = !g.isNight
}

func (g *Game) StartPhase() {
	if g.isNight {

	} else {

	}
}
