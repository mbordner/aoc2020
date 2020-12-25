package combat

import (
	"regexp"
	"sort"
	"strconv"
)

var (
	rePlayerHeader = regexp.MustCompile(`^Player\s(\d+):$`)
)

type Game struct {
	round   int
	players []*Player
}

func (g *Game) playRound() bool {
	cards := make([]int, 0, len(g.players))
	winningPlayer := -1
	winningCard := -1
	for i := range g.players {
		card := g.players[i].RemoveTop()
		if card > winningCard {
			winningCard = card
			winningPlayer = i
		}
		cards = append(cards, card)
	}
	sort.Ints(cards)
	g.players[winningPlayer].AddCards(cards)
	g.round++
	return g.isGameOver()
}

func (g *Game) isGameOver() bool {
	playersHavingCards := 0
	for i := range g.players {
		if g.players[i].HasCards() {
			playersHavingCards++
		}
	}
	if playersHavingCards == 1 {
		return true
	}
	return false
}

func (g *Game) PlayGame() int {
	for !g.playRound() {
	}

	var winner *Player
	for i := range g.players {
		if g.players[i].HasCards() {
			winner = g.players[i]
			break
		}
	}

	score := 0

	if winner != nil {
		cards := winner.GetCards()
		for i, c := range cards {
			score += c * (i + 1)
		}
	}

	return score
}

func NewGame(lines []string) *Game {
	g := &Game{}
	g.players = make([]*Player, 0, 5)

	i := 0
	var player *Player
	for i < len(lines) {
		if lines[i] == "" {
			player = nil
		} else {
			if player == nil {
				matches := rePlayerHeader.FindStringSubmatch(lines[i])
				id, _ := strconv.ParseInt(matches[1], 10, 32)
				player = NewPlayer(int(id))
				g.players = append(g.players, player)
			} else {
				card, _ := strconv.ParseInt(lines[i], 10, 32)
				player.AddCard(int(card))
			}
		}

		i++
	}

	return g
}
