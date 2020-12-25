package rcombat

import (
	"fmt"
	"github.com/mbordner/aoc2020/common/array/ints"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	rePlayerHeader = regexp.MustCompile(`^Player\s(\d+):$`)
)

type Game struct {
	states  map[string]int
	round   int
	winner  int
	players []*Player
}

func (g *Game) updateStates(s map[string]int) {
	g.states = s
}

func (g *Game) State() string {
	ps := make([]string, len(g.players), len(g.players))
	for i := range g.players {
		ps[i] = g.players[i].String()
	}
	return fmt.Sprintf("--\n%s", strings.Join(ps, "\n"))
}

func (g *Game) playRound() bool {
	cards := make([]int, 0, len(g.players))

	winningPlayer := -1
	winningCard := -1

	loopDetected := false

	state := g.State()
	if _, ok := g.states[state]; ok {
		loopDetected = true
	}
	g.states[state] = g.round

	if loopDetected {
		g.winner = 0 // loop detected forces player 1 to win
		return true
	} else {
		for i := range g.players {
			card := g.players[i].RemoveTop()
			if card > winningCard {
				winningCard = card
				winningPlayer = i
			}
			cards = append(cards, card)
		}
		canRecurse := true
		for i := range cards {
			if g.players[i].NumCards() < cards[i] {
				canRecurse = false
				break
			}
		}
		if canRecurse {

			rg := NewRecursiveGame(g, cards)
			rg.PlayGame()
			winningPlayer = rg.winner
			winningCard = cards[winningPlayer]

			cards = append(ints.Remove(cards, winningCard), winningCard)

		} else {
			sort.Ints(cards)
		}
	}

	g.players[winningPlayer].AddCards(cards)
	g.round++
	return g.isGameOver()
}

func (g *Game) isGameOver() bool {
	winner := -1
	playersHavingCards := 0
	for i := range g.players {
		if g.players[i].HasCards() {
			playersHavingCards++
			winner = i
		}
	}
	if playersHavingCards == 1 {
		g.winner = winner
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

func (g *Game) init() {
	g.players = make([]*Player, 0, 5)
	g.states = make(map[string]int)
}

func NewRecursiveGame(g *Game, cardsCount []int) *Game {
	rg := &Game{}
	rg.init()
	//rg.updateStates(g.states)

	for i, gp := range g.players {
		rp := NewPlayer(gp.GetID())
		cards := ints.Reverse(gp.GetTopCards(cardsCount[i]))
		for c := range cards {
			rp.AddCard(cards[c])
		}
		rg.players = append(rg.players, rp)
	}

	return rg
}

func NewGame(lines []string) *Game {
	g := &Game{}
	g.init()

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
